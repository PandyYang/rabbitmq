package main

import (
"fmt"
RabbitMQ "rabbitmq/ra2"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"imoocSimple")
	rabbitmq.PublishSimple("Hello imooc!")
	fmt.Println("发送成功！")
}
