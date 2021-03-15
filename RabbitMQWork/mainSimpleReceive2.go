package main

import (
	"rabbitmq/RabbitMQ"
)

func main() {
	rabbitmq := rabbitmq.NewRabbitMQSimple("111")
	rabbitmq.ConsumeSimple()
}
