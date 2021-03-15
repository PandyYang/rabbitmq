package main

import (
	"rabbitmq/RabbitMQ"
)

func main() {
	rabbitmq := rabbitmq.NewRabbitMQPubSub("pubsub")
	rabbitmq.ReceiveSub()
}
