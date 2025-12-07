package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/streadway/amqp"
)

type Weather struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	WindSpeed   float64 `json:"wind_speed"`
	Condition   int     `json:"condition"`
	Timestamp   string  `json:"timestamp"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	rabbitmqUser := os.Getenv("RABBITMQ_USER")
	rabbitmqPass := os.Getenv("RABBITMQ_PASS")
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	rabbitmqPort := os.Getenv("RABBITMQ_PORT")
	queueName := os.Getenv("QUEUE_NAME")

	if rabbitmqUser == "" {
		rabbitmqUser = "guest"
	}
	if rabbitmqPass == "" {
		rabbitmqPass = "guest"
	}
	if rabbitmqHost == "" {
		rabbitmqHost = "rabbitmq"
	}
	if rabbitmqPort == "" {
		rabbitmqPort = "5672"
	}
	if queueName == "" {
		queueName = "weather_queue"
	}

	// Conexão com RabbitMQ
	conn, err := amqp.Dial("amqp://" + rabbitmqUser + ":" + rabbitmqPass + "@" + rabbitmqHost + ":" + rabbitmqPort + "/")
	failOnError(err, "Erro ao conectar ao RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Erro ao abrir canal")
	defer ch.Close()

	// Garante que a fila exista
	_, err = ch.QueueDeclare(
		queueName, // nome da fila
		true,      // durable
		false,     // auto-delete
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Erro ao declarar fila")

	msgs, err := ch.Consume(
		queueName, // fila
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Erro ao consumir fila")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var weather Weather
			err := json.Unmarshal(d.Body, &weather)
			if err != nil {
				log.Printf("Erro ao decodificar mensagem: %s", err)
				continue
			}

			log.Printf("Mensagem recebida: %+v", weather)

			// Envia para o backend
			backendURL := "http://backend:3000/weather"
			jsonData, _ := json.Marshal(weather)
			resp, err := http.Post(backendURL, "application/json",
				bytes.NewBuffer(jsonData))
			if err != nil {
				log.Printf("Erro ao enviar para backend: %s", err)
			} else {
				log.Printf("Enviado para backend, status: %s", resp.Status)
				resp.Body.Close()
			}
		}
	}()

	log.Printf(" [*] Aguardando mensagens da fila %s. Para sair pressione CTRL+C", queueName)
	<-forever
}
