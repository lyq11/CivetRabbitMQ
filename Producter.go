package CivetRabbitMQ

import (
	"CivetRabbitMQ/BaseMQInstance"
	"github.com/streadway/amqp"
)

type Producer struct {
	MQInstance   *BaseMQInstance.RabbitMQ
	ProducerType int
	QueueName    string
	URL          string
	Exchange     string
	key          string
}

func NewProducer(url, QueueName, exchange, key string, producerType int) *Producer {
	newB := BaseMQInstance.NewRabbitMQ(QueueName, exchange, key, url)
	NewP := &Producer{
		MQInstance:   newB,
		ProducerType: producerType,
		URL:          url,
		Exchange:     exchange,
		key:          key,
		QueueName:    QueueName,
	}
	return NewP
}
func (p *Producer) PublishPub(message string) {
	//1.尝试创建交换机
	err := p.MQInstance.Channel.ExchangeDeclare(
		p.Exchange,
		//订阅模式下为广播类型
		"fanout",
		true,
		false,
		//true表示这个exchange不可以被client用来推送消息,仅用来进行exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	p.MQInstance.FailOnErr(err, "Failed to declare an exchange!")

	//2.发送消息
	err = p.MQInstance.Channel.Publish(
		p.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}
