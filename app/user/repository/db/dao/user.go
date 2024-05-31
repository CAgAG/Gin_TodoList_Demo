package dao

import (
	"TodoList_demo/app/user/repository/db/model"
	"context"
	"gorm.io/gorm"
)

// 对应的数据库操作
type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) FindUserByUserName(user_name string) (ret *model.User, err error) {
	err = dao.Model(&model.User{}).Where("user_name=?", user_name).Find(&ret).Error
	return
}

func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = dao.Model(&model.User{}).Create(&user).Error
	return
}
