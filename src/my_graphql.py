import strawberry
from typing import List
from sqlalchemy.orm import Session
from my_database import SessionLocal
from models import Article

@strawberry.type
class ArticleInfo:
    id: int
    name: str
    text: str
    complexity: str
    reading_time: str
    tags: List[str]
    likes: int = 0
    liked_by_user: bool = False

@strawberry.type
class PaginationInfo:
    page: int
    page_size: int
    has_next_page: bool
    has_previous_page: bool

@strawberry.type
class ArticlesPagination:
    items: List[ArticleInfo]
    page_info: PaginationInfo

def get_articles(page: int = 1, page_size: int = 20) -> ArticlesPagination:
    db: Session = SessionLocal()
    query = db.query(Article).offset((page - 1) * page_size).limit(page_size).all()
    db.close()
    items = [
        ArticleInfo(
            id=article.id,
            name=article.name,
            text=article.text,
            complexity=article.complexity,
            reading_time=article.reading_time,
            tags=article.tags.split(','),
            likes=0,
            liked_by_user=False
        )
        for article in query
    ]
    page_info = PaginationInfo(
        page=page,
        page_size=page_size,
        has_next_page=len(items) == page_size,
        has_previous_page=page > 1
    )
    return ArticlesPagination(items=items, page_info=page_info)

def get_article(id: int) -> ArticleInfo:
    db: Session = SessionLocal()
    article = db.query(Article).filter(Article.id == id).first()
    db.close()
    if article:
        return ArticleInfo(
            id=article.id,
            name=article.name,
            text=article.text,
            complexity=article.complexity,
            reading_time=article.reading_time,
            tags=article.tags.split(','),
            likes=0,
            liked_by_user=False
        )
    return None

@strawberry.type
class Query:
    articles: ArticlesPagination = strawberry.field(resolver=get_articles)
    article: ArticleInfo = strawberry.field(resolver=get_article)

schema = strawberry.Schema(query=Query)