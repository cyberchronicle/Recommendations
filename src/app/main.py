from contextlib import asynccontextmanager

from fastapi import FastAPI
from starlette.middleware.cors import CORSMiddleware

from app.api.main import api_router
from app.common import BaseEntity
from app.core import db
from app.core.config import settings

from strawberry.fastapi import GraphQLRouter
from app.my_graphql import schema


@asynccontextmanager
async def lifespan(_: FastAPI):
    BaseEntity.metadata.create_all(bind=db.engine)
    yield


app = FastAPI(
    title=settings.PROJECT_NAME,
    openapi_url=f"{settings.API_V1_STR}/openapi.json",
    lifespan=lifespan,
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.include_router(api_router, prefix=settings.API_V1_STR)

graphql_app = GraphQLRouter(schema)

app.include_router(graphql_app, prefix="/graphql")
