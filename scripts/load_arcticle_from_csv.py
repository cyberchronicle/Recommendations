import csv
import sys
from sqlalchemy.orm import Session
from my_database import SessionLocal, engine
from src.models import Article

from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

    
def import_articles_from_csv(csv_file_path):
    db = SessionLocal()

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
            db.add(article)
        
        db.commit()
    
    db.close()

if __name__ == "__main__":
    csv_file_path = 'data/articles.csv'
    import_articles_from_csv(csv_file_path)