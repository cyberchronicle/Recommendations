import pytest
from httpx import AsyncClient
from fastapi import FastAPI
from src.app.main import app  

@pytest.mark.asyncio
async def test_get_articles():
    query = """
    query {
        articles(page: 1, pageSize: 2) {
            items {
                id
                name
                text
                complexity
                reading_time
                tags
                likes
                liked_by_user
            }
            page_info {
                page
                page_size
                has_next_page
                has_previous_page
            }
        }
    }
    """
    async with AsyncClient(app=app, base_url="http://test") as ac:
        response = await ac.post("/relevant", json={"query": query})
    assert response.status_code == 200
    data = response.json()
    assert "data" in data
    assert "articles" in data["data"]
    assert "items" in data["data"]["articles"]
    assert len(data["data"]["articles"]["items"]) <= 2

@pytest.mark.asyncio
async def test_get_article():
    query = """
    query {
        article(id: 1) {
            id
            name
            text
            complexity
            reading_time
            tags
            likes
            liked_by_user
        }
    }
    """
    async with AsyncClient(app=app, base_url="http://test") as ac:
        response = await ac.post("/relevant", json={"query": query})
    assert response.status_code == 200
    data = response.json()
    assert "data" in data
    assert "article" in data["data"]
    assert data["data"]["article"]["id"] == 1
