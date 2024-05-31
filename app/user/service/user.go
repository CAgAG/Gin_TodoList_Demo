package service

import (
	"TodoList_demo/app/user/repository/db/dao"
	"TodoList_demo/app/user/repository/db/model"
	"TodoList_demo/grpc_proto/pb"
	"TodoList_demo/pkg/status"
	"context"
	"errors"
	"sync"
)

type UserService struct{}

var userServiceInstance *UserService
var userServiceOnce sync.Once

// 懒汉式的单例模式; lazy loading
// 只在一个协程里创建对象
func GetUserService() *UserService {
	userServiceOnce.Do(func() {
		userServiceInstance = &UserService{}
	})
	return userServiceInstance
}

// ==> 对立的是 饿汉式 ==> 会引发并发问题 ==> 多个协程创建对象
func GetUserServiceHungry() *UserService {
	if userServiceInstance == nil {
		return new(UserService)
	}
	return userServiceInstance
}

func (us *UserService) UserLogin(ctx context.Context, req *pb.UserRequest, resp *pb.UserResponse) (err error) {
	resp.Code = status.Success
	if req.UserName == "" {
		err = errors.New("用户不存在")
		resp.Code = status.Error
		return
	}
	// 查看用户是否存在
	find_user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		return
	}
	if find_user.ID == 0 {
		err = errors.New("用户不存在")
		resp.Code = status.Error
		return
	}
	if !find_user.CheckPassword(req.Password) {
		err = errors.New("用户密码错误")
		resp.Code = status.Error
		return
	}
	resp.UserDetail = ToPBUser(find_user)
	return nil
}

func (us *UserService) UserRegister(ctx context.Context, req *pb.UserRequest, resp *pb.UserResponse) (err error) {
	resp.Code = status.Success
	if req.Password != req.PasswordConfirm {
		resp.Code = status.Error
		err = errors.New("两次密码不一致")
		return
	}

	if req.UserName == "" {
		resp.Code = status.Error
		err = errors.New("用户名为空")
		return
	}

	// 查看用户是否存在
	find_user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if find_user.ID > 0 {
		err = errors.New("用户已存在")
		resp.Code = status.Error
		return
	}

	new_user := new(model.User)
	new_user.UserName = req.UserName
	if err = new_user.SetPassword(req.Password); err != nil {
		resp.Code = status.Error
		return
	}
	if err = dao.NewUserDao(ctx).CreateUser(new_user); err != nil {
		resp.Code = status.Error
		return
	}
	return nil
}

func ToPBUser(user *model.User) *pb.UserModel {
	return &pb.UserModel{
		Id:        uint32((user.ID)),
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}
}
