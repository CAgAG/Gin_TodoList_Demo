package mq

import (
	"TodoList_demo/conf"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	logging "github.com/sirupsen/logrus"
)

var RabbitMq *amqp091.Connection

func NewRabbitMQ() error {
	conn_str := fmt.Sprintf("%s://%s:%s@%s:%s/", conf.RabbitMQ, conf.RabbitMQUser, conf.RabbitMQPassWord, conf.RabbitMQHost, conf.RabbitMQPort)
	logging.Info("MQ ======================")
	logging.Info(conn_str)
	logging.Info("MQ ======================")
	conn, err := amqp091.Dial(conn_str)
	if err != nil {
		logging.Info("RabbitMQ 连接失败")
		logging.Info(err)
		return err
	}
	RabbitMq = conn
	return nil
}
