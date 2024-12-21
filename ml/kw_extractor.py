import yake
from typing import List, Union, Dict
import yaml
from ml_models import ArticleRequest
import pandas as pd
from collections import Counter

class KeyWordsExtractor:
    def __init__(self, config_path: Union[str, dict] = None):
        """
        config example:
                config = {
                "language": "ru",
                "numOfKeywords": 10,
                "max_ngram_size": 1,
                "windowSize": 5,
            }
        :param config:
        """ 
        if config_path is None:
            self.kw_extractor = yake.KeywordExtractor()
            return
        elif isinstance(config_path, str):
            with open(config_path, 'r') as f:
                config = yaml.load(f, Loader=yaml.SafeLoader)
        elif isinstance(config_path, dict):
            config = config_path

        self.kw_extractor = yake.KeywordExtractor(
            lan=config.get('language', 'ru'),
            n=config.get('max_ngram_size', 1),
            windowsSize=config.get('windowSize', 5),
            top=config.get('numOfKeywords', 10),
            dedupLim=0.9,
            dedupFunc='seqm'
        )
        with open('test.txt', 'r') as f:
            text = " ".join(f.readlines())
        warmup_res = self.extract(text)
        print(f"INFO: ML-service. Test text first 5 tags: {warmup_res[:5]}. {len(warmup_res)} tags at all")
        
    def extract(self, text: str) -> List[str]:
        keywords_score = self.kw_extractor.extract_keywords(text)
        keywords = []
        for keyword in keywords_score:
            keywords.append(keyword[0])
        return keywords

    def score_articles(self, user_tags: List[str], articles: List[ArticleRequest]) -> List[Union[str,int]]:
        scored_articles = []
        for article in articles:
            score = sum(tag in user_tags for tag in article.tags) / (max(len(user_tags), len(article.tags)) + 1e-8)
            scored_articles.append((article.id, score))
            
        scored_articles.sort(key=lambda x: x[1], reverse=True)
        res_list = [article_id for article_id, _ in scored_articles]
    
        return res_list
    
    def all_tags(self, filepath: str) -> bool:
        df = pd.read_csv(filepath)
        articles_text = df['text'].tolist()
        all_keywords = []

        for article in articles_text:
            keywords = self.kw_extractor.extract(article)
            all_keywords.extend(keywords)
            print(keywords)

        keyword_counts = Counter(all_keywords)
        keyword_df = pd.DataFrame(keyword_counts.items(), columns=['tag', 'rate'])
        keyword_df = keyword_df.sort_values(by="rate", ascending=False)

        out_filepath = 'keywords_rate_' + filepath[3:].replace('/', '-')
        keyword_df.to_csv(out_filepath, index=False)
        return True

        
kw_extractor = KeyWordsExtractor(config_path="../configs/key_words_extractor.yaml") 
# TODO: get path from settings
