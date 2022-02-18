package BaseMQInstance

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	Channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	key string
	//连接信息
	Mqurl string
}

func NewRabbitMQ(queueName, exchange, key string, url string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName: queueName, Exchange: exchange, key: key, Mqurl: url}
	var err error
	//创建RabbitMQ连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.FailOnErr(err, "创建连接错误!")
	rabbitmq.Channel, err = rabbitmq.conn.Channel()
	rabbitmq.FailOnErr(err, "获取channel失败!")
	return rabbitmq
}

// Destroy 断开channel和connection
func (r *RabbitMQ) Destroy() {
	r.Channel.Close()
	r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) FailOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s,%s", message, err))
	}
}
