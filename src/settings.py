import os
from pydantic_settings import BaseSettings

class Settings(BaseSettings):
    database_url: str = "sqlite:///./articles.db"

    class Config:
        env_file = ".env"

settings = Settings()