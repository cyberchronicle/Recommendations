from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import List
from sqlalchemy import create_engine, Column, Integer, String
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

app = FastAPI()

DATABASE_URL = "sqlite:///./test.db"  # TODO: унести адресс в конфиги/окружение 
engine = create_engine(DATABASE_URL, connect_args={"check_same_thread": False})
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

Base = declarative_base()


class Article(Base):
    __tablename__ = "articles"

    id = Column(Integer, primary_key=True, index=True)
    key_words = Column(String, index=True)


Base.metadata.create_all(bind=engine)


class ArticleCreate(BaseModel):
    id: int
    key_words: List[str]


@app.post("/add_article/")
async def add_article(article: ArticleCreate):
    db = SessionLocal()
    db_article = Article(id=article.id, key_words=",".join(article.key_words))
    db.add(db_article)
    db.commit()
    db.refresh(db_article)
    db.close()
    return {"message": "Article added successfully"}


@app.get("/search_by_keyword/")
async def search_by_keyword(keyword: str):
    db = SessionLocal()
    results = db.query(Article).filter(Article.key_words.like(f"%{keyword}%")).all()
    db.close()
    if results:
        return {"ids": [result.id for result in results]}
    else:
        raise HTTPException(status_code=404, detail="No articles found for this keyword")
