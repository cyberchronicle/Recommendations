from sqlalchemy import Column, Integer, String
from my_database import Base

class Article(Base):
    __tablename__ = "articles"

    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, index=True)
    text = Column(String)
    complexity = Column(String)
    reading_time = Column(String)
    tags = Column(String)  # Store tags as a comma-separated string