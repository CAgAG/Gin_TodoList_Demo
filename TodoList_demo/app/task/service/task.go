package service

import (
	"TodoList_demo/app/task/repository/db/dao"
	"TodoList_demo/app/task/repository/db/model"
	"TodoList_demo/app/task/repository/mq"
	"TodoList_demo/grpc_proto/pb"
	"TodoList_demo/pkg/status"
	"context"
	"encoding/json"
	"sync"
)

type TaskService struct{}

var TaskServiceInstance *TaskService
var TaskServiceOnce sync.Once

// 懒汉式的单例模式; lazy loading
// 只在一个协程里创建对象
func GetTaskService() *TaskService {
	TaskServiceOnce.Do(func() {
		TaskServiceInstance = &TaskService{}
	})
	return TaskServiceInstance
}

// ==> 对立的是 饿汉式 ==> 会引发并发问题 ==> 多个协程创建对象
func GetTaskServiceHungry() *TaskService {
	if TaskServiceInstance == nil {
		return new(TaskService)
	}
	return TaskServiceInstance
}

// 发送任务到mq, 再由mq去存入数据库
func (ts *TaskService) CreateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = status.Success

	// 异步发送消息
	body, err := json.Marshal(req)
	err = mq.SendMessage2MQ(body)
	if err != nil {
		resp.Code = status.Error
		return
	}
	return nil
}

// Mq 存入数据库
func TaskMq2DB(ctx context.Context, req *pb.TaskRequest) (err error) {
	task := &model.Task{
		Uid:       uint(req.Uid),
		Title:     req.Title,
		Status:    int(req.Status),
		Content:   req.Content,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	return dao.NewTaskDao(ctx).CreateTask(task)
}

func (ts *TaskService) GetTasksList(ctx context.Context, request *pb.TaskRequest, response *pb.TaskListResponse) (err error) {
	response.Code = status.Success

	if request.Limit == 0 {
		request.Limit = 10
	}
	tasks, count, err := dao.NewTaskDao(ctx).GetTasksByUId(request.Uid, int(request.Start), int(request.Limit))
	if err != nil {
		response.Code = status.Error
		return
	}

	var taskRes []*pb.TaskModel
	for _, item := range tasks {
		taskRes = append(taskRes, Tasks2PbTaskModel(item))
	}
	response.TaskList = taskRes
	response.Count = uint32(count)

	return
}

func Tasks2PbTaskModel(task *model.Task) *pb.TaskModel {
	return &pb.TaskModel{
		Id:          uint64(task.ID),
		Uid:         uint64(task.Uid),
		Title:       task.Title,
		Content:     task.Content,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,
		Status:      int64(task.Status),
		CreatedTime: task.CreatedAt.Unix(),
		UpdatedTime: task.UpdatedAt.Unix(),
	}
}

func (ts *TaskService) GetTask(ctx context.Context, request *pb.TaskRequest, response *pb.TaskDetailResponse) (err error) {
	response.Code = status.Success

	task, err := dao.NewTaskDao(ctx).GetTaskByUidAndId(request.Uid, request.Id)
	if err != nil {
		response.Code = status.Error
		return
	}
	if task == nil || task.ID == 0 {
		return
	}
	response.TaskDetail = Tasks2PbTaskModel(task)

	return
}

func (ts *TaskService) UpdateTask(ctx context.Context, request *pb.TaskRequest, response *pb.TaskDetailResponse) (err error) {
	response.Code = status.Success

	err = dao.NewTaskDao(ctx).UpdateTaskByRequest(request)
	if err != nil {
		response.Code = status.Error
		return
	}
	return
}

func (ts *TaskService) DeleteTask(ctx context.Context, request *pb.TaskRequest, response *pb.TaskDetailResponse) (err error) {
	response.Code = status.Success

	err = dao.NewTaskDao(ctx).DelTaskByUidAndId(request.Uid, request.Id)
	if err != nil {
		response.Code = status.Error
		return
	}
	return
}
