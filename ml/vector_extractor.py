from typing import List, Union
import pandas as pd
import yaml

from sentence_transformers import SentenceTransformer
import numpy as np
import random

class VectorExtractor:
    def __init__(self, config_path: Union[str, dict] = None):
        """
        config example:
                config = {
                //
            }
        :param config:
        """ 
        if config_path is None:
            self.v_extractor = SentenceTransformer("deepvk/USER-bge-m3")
            return
        elif isinstance(config_path, str):
            with open(config_path, 'r') as f:
                config = yaml.load(f, Loader=yaml.SafeLoader)
        elif isinstance(config_path, dict):
            config = config_path


        self.v_extractor = SentenceTransformer(
            config.get('model_name', "deepvk/USER-bge-m3"))
        
        #main_requst_compare
        filepath = '../data/articles.csv'
        text = self.read_article(filepath)
        request = 'Как сделать бота на битоне код для чайников'

        similarity_score = self.compare(text, request)
        print(f"Text: {text[:30]}...")
        print(f"Request: {request}")
        print(f"Similarity score: {similarity_score:.4f}")

    def read_article(self, filepath: str) -> str:
        df = pd.read_csv(filepath)
        articles_text = df['text'].tolist()
        ind = random.randrange(0, len(articles_text))
        
        text = articles_text[ind]
        return text
        
    def extract(self, text: str) -> np.ndarray:
        embeddings = self.v_extractor.encode(text)
        return embeddings
    
    def cosine_similarity(self, 
                          text1: np.ndarray, 
                          text2: np.ndarray) -> float:
        norm1 = np.linalg.norm(text1)
        norm2 = np.linalg.norm(text2)
        cos_v = np.dot(text1, text2) / (norm1 * norm2) if norm1 and norm2 else 0.0
        return cos_v

    def compare(self, text: str, request: str) -> float:
        text_embeddings = self.extract(text)
        request_embeddings = self.extract(request)

        similarity_score = self.cosine_similarity(text_embeddings, request_embeddings)
        return similarity_score

        
v_extractor = VectorExtractor(config_path="../configs/vector_extractor.yaml")
