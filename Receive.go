package main

import (
RabbitMQ "rabbitmq/ra2"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"imoocSimple")
	rabbitmq.ConsumeSimple()
}
