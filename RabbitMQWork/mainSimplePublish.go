package main

import (
	"fmt"
	"rabbitmq/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("111")
	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("send Message" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	//rabbitmq.PublishSimple("Hello RabbitMQ!")
	fmt.Println("发送成功!")
}
