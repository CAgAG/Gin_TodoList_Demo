package dao

import (
	"TodoList_demo/app/task/repository/db/model"
	"TodoList_demo/grpc_proto/pb"
	"context"
	"gorm.io/gorm"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewDBClient(ctx)}
}

func (dao *TaskDao) CreateTask(data *model.Task) error {
	return dao.Model(&model.Task{}).Create(&data).Error
}

// uid 就是 user_id
func (dao *TaskDao) GetTasksByUId(user_id uint64, start, limit int) (tasks []*model.Task, count int64, err error) {
	err = dao.Model(&model.Task{}).Offset(start).Limit(limit).Where("uid=?", user_id).Find(&tasks).Error
	if err != nil {
		return
	}
	err = dao.Model(&model.Task{}).Where("uid=?", user_id).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (dao *TaskDao) GetTaskByUidAndId(uid, id uint64) (task *model.Task, err error) {
	err = dao.Model(&model.Task{}).Where("id=? AND uid=?", id, uid).Find(&task).Error
	return
}

func (dao *TaskDao) UpdateTaskByRequest(req *pb.TaskRequest) (err error) {
	var task *model.Task
	err = dao.Model(&model.Task{}).Where("id=? AND uid=?", req.Id, req.Uid).Find(&task).Error
	if err != nil {
		return
	}
	task.Title = req.Title
	task.Status = int(req.Status)
	task.Content = req.Content
	err = dao.Save(&task).Error
	return
}

func (dao *TaskDao) DelTaskByUidAndId(uid, id uint64) (err error) {
	err = dao.Model(&model.Task{}).Where("id=? AND uid=?", id, uid).Delete(&model.Task{}).Error
	return
}
