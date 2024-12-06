from kafka import KafkaConsumer
import multiprocessing
import yaml

stop_event = multiprocessing.Event()

with open("/var/configs/scrapper_kafka.yaml", 'r') as f:
    config = yaml.load(f, Loader=yaml.SafeLoader)
    config = config.get('kafka', {})
    
    
server_url = config.get("brokers", ["localhost:9092"])[0]
topic = config.get("topic", "scrapping")

def main():
    consumer = KafkaConsumer(bootstrap_servers=server_url)
    consumer.subscribe([topic])
    
    while not stop_event.is_set():    
        for message in consumer:
            print(" Topic: " + str(message[0])
            + "\n Message: " + str(message[6], 'utf-8')
            + "\n Record: " + str(message))
            if stop_event.is_set():
                break
    consumer.close()

if __name__ == '__main__':
    main()