from confluent_kafka import Consumer, KafkaError
import requests
from my_database import SessionLocal
from models import Article

# Configuration for Kafka #TODO: get from config
KAFKA_BROKERS = "kafka:29092"
KAFKA_TOPIC = "scrapping"
SCRAPPER_API_URL = "http://localhost:9003/api/v1/scrapper/article/"

def fetch_article_from_scrapper(article_id):
    """Fetch article data from the scrapper API."""
    response = requests.get(f"{SCRAPPER_API_URL}{article_id}")
    if response.status_code == 200:
        return response.json()
    else:
        print(f"Failed to fetch article {article_id}: {response.status_code}")
        return None

def save_article_to_db(article_data):
    """Save the fetched article data to the database."""
    db = SessionLocal()
    article = Article(
        id=article_data['id'],
        name=article_data['name'],
        text=article_data['text'],
        complexity=article_data['complexity'],
        reading_time=article_data['readingTime'],
        tags=",".join(article_data['tags'])
    )
    db.add(article)
    db.commit()
    db.close()

def consume_kafka_messages():
    """Consume messages from Kafka and process them."""
    consumer = Consumer({
        'bootstrap.servers': KAFKA_BROKERS,
        'group.id': 'article-consumer-group',
        'auto.offset.reset': 'earliest'
    })

    consumer.subscribe([KAFKA_TOPIC])

    try:
        while True:
            msg = consumer.poll(1.0)  # Timeout in seconds

            if msg is None:
                continue
            if msg.error():
                if msg.error().code() == KafkaError._PARTITION_EOF:
                    continue
                else:
                    print(msg.error())
                    break

            article_id = msg.value().decode('utf-8')
            print(f"Received article ID: {article_id}")

            article_data = fetch_article_from_scrapper(article_id)
            if article_data:
                save_article_to_db(article_data)

    finally:
        consumer.close()

if __name__ == "__main__":
    consume_kafka_messages()