package main

import (
	"rabbitmq/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("pubsub")
	rabbitmq.ReceiveSub()
}
