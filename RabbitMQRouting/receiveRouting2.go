package main

import "rabbitmq/RabbitMQ"

func main() {
	one := RabbitMQ.NewRabbitMQRouting("222", "222_2")
	one.ReceiveRouting()
}
