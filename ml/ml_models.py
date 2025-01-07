from pydantic import BaseModel
from typing import List, Dict, Union


class ArticleRequest(BaseModel):
    id: str
    tags: List[str]
    embedding: List[float] = []

class Articles(BaseModel):
    articles: List[ArticleRequest]

class SuggestRequest(BaseModel):
    user_tags: List[str]

class TextProcessRequest(BaseModel):
    text: str

class TextEmbeddingRequest(BaseModel):
    text: str
    
class SuggestResponse(BaseModel):
    ids: List[str]

class TextEmbeddingResponse(BaseModel):
    embedding: List[float] 

class TextProcessResponse(BaseModel):
    tags: List[str]