from fastapi import FastAPI, HTTPException
from typing import List, Dict, Union

from kw_extractor import KeyWordsExtractor
from vector_extractor import VectorExtractor
from ml_models import SuggestRequest, TextProcessRequest, TextEmbeddingRequest, ArticleRequest
from ml_models import SuggestResponse, TextProcessResponse, TextEmbeddingResponse

import json

app = FastAPI()
kw_extractor = KeyWordsExtractor(config_path="/ml/configs/key_words_extractor.yaml") 
v_extractor = VectorExtractor(config_path="/ml/configs/vector_extractor.yaml") 

@app.post("/suggest/{user_id}")
def suggest(user_id: int, request: SuggestRequest) -> SuggestResponse:
    try:
        recommended_ids = kw_extractor.score_articles(request.user_tags, request.articles)
        response = SuggestResponse(ids=recommended_ids)
        return response
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))

@app.post("/text/process")
def text_process(request: TextProcessRequest) -> TextProcessResponse:
    try:
        tags = kw_extractor.extract(request.text)
        response = TextProcessResponse(tags=tags)
        return response
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))


@app.post("/text/embedding")
def text_embedding(request: TextEmbeddingRequest):
    try:
        embedding = v_extractor.extract(request.text)
        response = TextEmbeddingResponse(embedding=embedding)
        return response
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))
    