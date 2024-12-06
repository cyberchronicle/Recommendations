import yake
from typing import List, Union, Dict
import yaml
from ml_models import ArticleRequest

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