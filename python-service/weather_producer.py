import pika
import json

connection = pika.BlockingConnection(pika.ConnectionParameters(host="rabbitmq"))
channel = connection.channel()
channel.queue_declare(queue="weather")

message = {"temperature": 25, "humidity": 60}
channel.basic_publish(exchange="", routing_key="weather", body=json.dumps(message))

print("Mensagem enviada:", message)
connection.close()
