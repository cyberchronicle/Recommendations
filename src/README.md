### Run

```bash
docker compose up
```

### For testing /add_article

```bash
curl -X POST http://0.0.0.0:8000/add_article/ \
     -H "Content-Type: application/json" \
     -d '{"id": 1, "key_words": ["fastapi", "python", "api"]}'
```

### For testing /search_by_keyword 

```bash
curl -X GET "http://0.0.0.0:8000/search_by_keyword/?keyword=fastapi"
```

Script to fill db with data (lays in data/articles.csv):

```bash
cd src/
python -m app.scripts.initialize_article_data
```

Do not forget to fill .env with credentials.
