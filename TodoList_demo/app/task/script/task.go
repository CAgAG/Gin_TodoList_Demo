package script

import (
	"TodoList_demo/app/task/repository/mq/listener"
	"context"
	logging "github.com/sirupsen/logrus"
)

func TaskCreateSync(ctx context.Context) {
	st := new(listener.SyncTask)
	err := st.ListenMqMessageByTaskService(ctx)
	if err != nil {
		logging.Info("启动task监听失败")
		return
	}
}
