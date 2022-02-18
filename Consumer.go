package CivetRabbitMQ

import (
	"CivetRabbitMQ/BaseMQInstance"
	"github.com/streadway/amqp"
)

type Consumer struct {
	MQInstance   *BaseMQInstance.RabbitMQ
	ConsumerType int
	QueueName    string
	URL          string
	Exchange     string
	key          string
	ConsumerName string
}

func NewConsumer(url, QueueName, exchange, key string, ConsumerType int, ConsumerName string) *Consumer {
	newB := BaseMQInstance.NewRabbitMQ(QueueName, exchange, key, url)
	NewP := &Consumer{
		MQInstance:   newB,
		ConsumerType: ConsumerType,
		URL:          url,
		Exchange:     exchange,
		key:          key,
		QueueName:    QueueName,
		ConsumerName: ConsumerName,
	}
	return NewP
}
func (r *Consumer) StartReceive() {
	//1.试探性创建交换机
	err := r.MQInstance.Channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"fanout",
		true,
		false,
		//true表示这个exchange不可以被client用来推送消息,仅用来进行exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	r.MQInstance.FailOnErr(err, "Failed to declare an exchange!")

	//2.试探性创建队列,注意队列名称不要写
	q, err := r.MQInstance.Channel.QueueDeclare(
		r.QueueName, //随机生成队列名称
		false,
		false,
		false,
		false,
		nil,
	)
	r.MQInstance.FailOnErr(err, "Failed to declare an exchange!")

	//3.绑定队列到exchange中
	err = r.MQInstance.Channel.QueueBind(
		q.Name,
		//在pub/sub模式下,这里的key要为空
		"",
		r.Exchange,
		false,
		nil,
	)

	////4.消费消息
	//
	//go func() {
	//	for d := range messages {
	//		log.Printf("%s:Received a message :%s", r.ConsumerName, d.Body)
	//	}
	//}()
	//
	//fmt.Println("[*] Waiting for messages,To exit press CTRL+C")

}
func (r *Consumer) Received() <-chan amqp.Delivery {
	messages, err := r.MQInstance.Channel.Consume(
		r.QueueName,
		r.ConsumerName,
		true,
		false,
		false,
		false,
		nil,
	)
	r.MQInstance.FailOnErr(err, "received fail")
	return messages
}
