package main

import (
	"fmt"
	"rabbitmq/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("111")
	rabbitmq.PublishSimple("Hello RabbitMQ!")
	fmt.Println("发送成功!")
}
