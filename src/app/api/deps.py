from collections.abc import Generator
from typing import Annotated

from fastapi import Depends
from fastapi.security import HTTPBearer
from sqlalchemy.orm import Session

from app.core.db import engine

oauth2_scheme = HTTPBearer()

storage = {}


def get_db() -> Generator[Session, None, None]:
    with Session(engine) as session:
        yield session


def get_storage() -> Generator[Session, None, None]:
    yield storage


SessionDep = Annotated[Session, Depends(get_db)]
TokenDep = Depends(HTTPBearer())
StorageDep = Annotated[dict, Depends(get_storage)]

