import json

from pydantic import BaseModel
from typing import List


class ArticleCreate(BaseModel):
    id: int
    name: str
    text: str
    complexity: str
    reading_time: str
    tags: List[str]  # Use a list for tags

    class Config:
        orm_mode = True
        from_attributes = True


class ArticleTag(BaseModel):
    id: int
    tags: List[str]

    class Config:
        orm_mode = True
        from_attributes = True

    @classmethod
    def from_orm(cls, obj):
        tags_as_list = json.loads(obj.tags) if obj.tags else []
        return cls(id=obj.id, tags=tags_as_list)
