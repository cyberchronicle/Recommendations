import csv
import sys

from sqlalchemy.orm import Session

from app.models import ArticleTag
from app.core import db


def initialize_article_data(file_path: str):
    session = db.SessionLocal()
    csv.field_size_limit(sys.maxsize)

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
    csv_file_path = '../data/articles.csv'
    # import_articles_from_csv(csv_file_path)
    initialize_article_data(csv_file_path)
