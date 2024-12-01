from sqlalchemy import Column, Integer, String
from app.common import BaseEntity


class Article(BaseEntity):
    __tablename__ = "articles"

    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, index=True)
    text = Column(String)
    complexity = Column(String)
    reading_time = Column(String)
    tags = Column(String)  # Store tags as a comma-separated string


class ArticleTag(BaseEntity):
    __tablename__ = "articles_tags"

    id = Column(Integer, primary_key=True, index=True)
    tags = Column(String)
