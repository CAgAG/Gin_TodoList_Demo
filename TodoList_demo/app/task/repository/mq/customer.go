package mq

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	logging "github.com/sirupsen/logrus"
)

func ReceiveMessageByMq(ctx context.Context, queueName string) (msg <-chan amqp091.Delivery, err error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		logging.Info("RabbitMq 管道创建失败")
		return
	}

	q, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		logging.Info("RabbitMq 消息队列创建失败")
		return
	}
	err = ch.Qos(1, 0, false)
	if err != nil {
		logging.Info("RabbitMq 消息接收失败")
		return
	}
	return ch.Consume(q.Name, "", false, false, false, false, nil)
}
