package main

import (
	"rabbitmq/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("111")
	rabbitmq.ConsumeSimple()
}
