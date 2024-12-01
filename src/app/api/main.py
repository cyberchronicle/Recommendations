from fastapi import APIRouter

from app.api.routes import article


api_router = APIRouter()


api_router.include_router(article.router, prefix="/article", tags=["article"])
