import json

from sqlalchemy.orm import Session

from app.models import ArticleTag


def get_article_tag(session: Session, article_id: int) -> ArticleTag:
    db_article_tag = session.query(ArticleTag).filter(ArticleTag.id == article_id).first()
    return db_article_tag
