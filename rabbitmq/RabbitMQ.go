package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const MQURL = "amqp://admin:admin@192.168.241.241:5672/rabbitmq"

type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//Key
	Key string
	//连接信息
	Mqurl string
}

//创建RabbitMQ结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName:queueName,Exchange:exchange,Key:key,Mqurl:MQURL}
}

// 断开Channel和Connection
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

//错误处理逻辑
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatal("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

//1.创建简单模式下的rabbitmq实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	//创建RabbitMQ实例
	rabbitmq := NewRabbitMQ(queueName, "", "")
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabb"+
		"itmq!")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}

//2.简单那模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	//申请队列 队列不存在 创建队列 否则跳过创建
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, //是否自动持久化
		false, //是否为自动删除
		false, //是否具有排他性 其他用户不能访问
		false, //是否阻塞
		nil, //额外属性
		)
	if err != nil {
		fmt.Println(err)
	}
	//发送消息到队列
	r.channel.Publish(
			r.Exchange,
			r.QueueName,
			false,//如果为true 根据exchange类型和routkey规则 如果无法找到符合条件的队列 那么会把发送的消息返回给发送者
			false, //如果为true 当exchange发送消息后,发现队列上没有绑定消费者 则会把消息还给发送者
			amqp.Publishing{
				ContentType: "text/plain",
				Body: []byte(message),
			})
}

//3.简单模式下消费代码
func (r *RabbitMQ) ConsumeSimple() {
	//申请队列 队列不存在 创建队列 否则跳过创建
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		false, //是否自动持久化
		false, //是否为自动删除
		false, //是否具有排他性 其他用户不能访问
		false, //是否阻塞
		nil, //额外属性
	)
	if err != nil {
		fmt.Println(err)
	}
	//接收消息
	msgs, err := r.channel.Consume(
			q.Name,
			"", //区分多个消费者
			//是否自动应答
			true,
			//是否具有排他性
			false,
			//如果设置为true 表示不能将同一个connection中发送的消息传递给这个connection中的消费者
			false,
			false, //消费是否阻塞
			nil,
			)

	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msgs {
			//实现我们要处理的逻辑函数
			log.Printf("Received a message: %s\n", d.Body)
		}
	}()

	log.Printf("[*] waiting for messages to exit press exit + c")
	<- forever
}
