package listener

import (
	"TodoList_demo/app/task/repository/mq"
	"TodoList_demo/app/task/service"
	"TodoList_demo/conf"
	"TodoList_demo/grpc_proto/pb"
	"context"
	"encoding/json"
	logging "github.com/sirupsen/logrus"
)

type SyncTask struct{}

func (st *SyncTask) ListenMqMessageByTaskService(ctx context.Context) (err error) {
	rabbitQueue := conf.RabbitMqTaskQueue
	msg_chan, err := mq.ReceiveMessageByMq(ctx, rabbitQueue) // 通过 chan 来接收消息
	if err != nil {
		return
	}

	var keep_msg_chan chan struct{}
	go func() {
		for d := range msg_chan {
			// 数据库操作
			req := new(pb.TaskRequest)
			err = json.Unmarshal(d.Body, req)
			if err != nil {
				logging.Info("json 反序列化失败")
				return
			}
			err = service.TaskMq2DB(ctx, req)
			if err != nil {
				logging.Info("RabbitMq消息 存入数据库失败")
				return
			}
			// 返回应答, 防止消息积压
			d.Ack(false)
		}
	}()
	select {
	case <-keep_msg_chan:
		return
	}
}
