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