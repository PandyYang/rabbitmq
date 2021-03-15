package main

import (
	"fmt"
	"rabbitmq/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	one := RabbitMQ.NewRabbitMQRouting(
		"111", "111_1")
	two := RabbitMQ.NewRabbitMQRouting(
		"222", "222_2")
	for i := 0; i <= 10; i++ {
		one.PublishRouting("1111111111111111" + strconv.Itoa(i))
		two.PublishRouting("2222222222222222" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
