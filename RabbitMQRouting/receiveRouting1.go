package main

import "rabbitmq/RabbitMQ"

func main() {
	two := RabbitMQ.NewRabbitMQRouting("111", "111_1")
	two.ReceiveRouting()
}
