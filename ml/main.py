from fastapi import FastAPI, HTTPException
from typing import List, Dict, Union
import pandas as pd
import numpy as np
import ast
from kw_extractor import KeyWordsExtractor
from vector_extractor import VectorExtractor
from ml_models import SuggestRequest, TextProcessRequest, TextEmbeddingRequest, ArticleRequest, Articles
from ml_models import SuggestResponse, TextProcessResponse, TextEmbeddingResponse

import json

import logging

logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(levelname)s - %(message)s',
    datefmt='%Y-%m-%d %H:%M:%S'
)
def load_articles():
    path = "/ml/data/articles.csv"
    path_embs = "/ml/data/embs.npy"
    
    articles = pd.read_csv(path)
    embs = np.load(path_embs, allow_pickle=True)
    
    raw_list = [ArticleRequest(id=str(id_), tags=ast.literal_eval(tags)) for id_, tags in zip(articles['id'], articles['kw_tags'])]
    for i, ar in enumerate(raw_list):
        ar.embedding = embs[i]
    
    form_articles = Articles(articles=raw_list)
    
    return raw_list
    

app = FastAPI()
kw_extractor = KeyWordsExtractor(config_path="/ml/configs/key_words_extractor.yaml") 
v_extractor = VectorExtractor(config_path="/ml/configs/vector_extractor.yaml")

articles = load_articles()

@app.post("/suggest")
def suggest(request: SuggestRequest) -> SuggestResponse:
    try:
        recommended_ids = kw_extractor.score_articles(request.user_tags, articles)
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
    