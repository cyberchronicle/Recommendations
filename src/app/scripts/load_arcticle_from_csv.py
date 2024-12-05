import csv
import sys
import os

# Automatically determine the parent directory two levels up
parent_dir = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

from src.app.core import db
from src.app.models import Article, ArticleTag

    
def import_articles_from_csv(csv_file_path):
    session = db.SessionLocal()

    csv.field_size_limit(sys.maxsize)

    with open(csv_file_path, newline='', encoding='utf-8') as csvfile:
        reader = csv.DictReader(csvfile)
        
        for row in reader:
            article = Article(
                id=int(row['id']),
                name=row['name'],
                text=row['text'],
                complexity=row['complexity'],
                reading_time=row['reading_time'],
                tags=row['tags']  # Assuming tags are comma-separated in the CSV
            )
            session.add(article)
        
        session.commit()
    
    session.close()


def initialize_article_data(file_path: str):
    session = db.SessionLocal()

    with open(file_path, newline="", encoding="utf-8") as csvfile:
        reader = csv.DictReader(csvfile)
        rows_inserted = 0
        for row in reader:
            try:
                if len(row['tags']):
                    tag_entry = ArticleTag(
                        id=int(row['id']),
                        tags=row['tags']
                    )
                    session.add(tag_entry)
                    rows_inserted += 1
            except Exception as e:
                print(f"Skipping row due to error: {row}")
                print(f"Error: {str(e)}")

        session.commit()
        print(f"Inserted {rows_inserted} rows into the database.")

    session.close()


if __name__ == "__main__":
    csv_file_path = 'data/articles.csv'
    # import_articles_from_csv(csv_file_path)
    initialize_article_data(csv_file_path)
