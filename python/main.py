import time
import pika
import json
import requests
import os
from datetime import datetime

QUEUE_NAME = os.getenv("QUEUE_NAME", "weather_queue")
CITY = os.getenv("CITY", "Brasilia")
RABBITMQ_HOST = os.getenv("RABBITMQ_HOST", "rabbitmq")
RABBITMQ_PORT = int(os.getenv("RABBITMQ_PORT", 5672))
RABBITMQ_USER = os.getenv("RABBITMQ_USER", "guest")
RABBITMQ_PASS = os.getenv("RABBITMQ_PASS", "guest")

def get_weather():
    # Exemplo de coleta de dados (substitua pela API real)
    response = requests.get(f"https://api.open-meteo.com/v1/forecast?latitude=-15.78&longitude=-47.93&current_weather=true")
    data = response.json()["current_weather"]
    return {
        "city": CITY,
        "temperature": data["temperature"],
        "wind_speed": data["windspeed"],
        "condition": data["weathercode"],
        "timestamp": datetime.utcnow().isoformat()
    }

def send_to_queue(weather_data):
    credentials = pika.PlainCredentials(RABBITMQ_USER, RABBITMQ_PASS)
    connection = pika.BlockingConnection(pika.ConnectionParameters(
        host=RABBITMQ_HOST,
        port=RABBITMQ_PORT,
        credentials=credentials
    ))
    channel = connection.channel()
    channel.queue_declare(queue=QUEUE_NAME, durable=True)
    channel.basic_publish(
        exchange="",
        routing_key=QUEUE_NAME,
        body=json.dumps(weather_data),
        properties=pika.BasicProperties(delivery_mode=2)
    )
    print("Enviando dados:", weather_data)
    connection.close()

if __name__ == "__main__":
    while True:
        try:
            weather_data = get_weather()
            send_to_queue(weather_data)
        except Exception as e:
            print("Erro ao coletar/enviar:", e)
        time.sleep(3600)  # coleta a cada 1 hora
