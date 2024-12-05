from fastapi import APIRouter, Body, Path, File, HTTPException, UploadFile
from starlette import status

from app import schemas
from app.api.deps import SessionDep
from app.crud import article

router = APIRouter()


@router.get("/{article_id}", status_code=status.HTTP_201_CREATED, response_model=schemas.ArticleTag)
async def get_article_by_article_id(
        session: SessionDep,
        article_id: int = Path(...)
):
    db_article_tag = article.get_article_tag(session, article_id)
    if not db_article_tag:
        raise HTTPException(status_code=401, detail="Article not found")

    return schemas.ArticleTag.from_orm(db_article_tag)
