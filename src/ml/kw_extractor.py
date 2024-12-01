import yake


class KeyWordsExtractor:
    def __init__(self, config: dict = None):
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
        if config:
            self.kw_extractor = yake.KeywordExtractor(lan=config['language'],
                                                      n=config['max_ngram_size'],
                                                      windowsSize=config['windowSize'],
                                                      top=config['numOfKeywords'],
                                                      dedupLim=0.9,
                                                      dedupFunc='seqm',
                                                      )
        else:
            self.kw_extractor = yake.KeywordExtractor()

    def extract(self, text: str):
        keywords_score = self.kw_extractor.extract_keywords(text)
        keywords = []
        for keyword in keywords_score:
            keywords.append(keyword[0])
        return keywords
