package mq

import (
	"TodoList_demo/conf"
	"github.com/rabbitmq/amqp091-go"
	logging "github.com/sirupsen/logrus"
)

func SendMessage2MQ(body []byte) (err error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		logging.Info("RabbitMq 管道创建失败")
		return
	}

	q, err := ch.QueueDeclare(conf.RabbitMqTaskQueue, true, false, false, false, nil)
	if err != nil {
		logging.Info("RabbitMq 消息队列创建失败")
		return
	}

	err = ch.Publish("", q.Name, false, false, amqp091.Publishing{
		DeliveryMode: amqp091.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})

	if err != nil {
		logging.Info("RabbitMq 消息队列发送失败")
		return
	}
	logging.Info("RabbitMq 消息队列发送成功")
	return
}
