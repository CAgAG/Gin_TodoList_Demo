package utils

import (
	"context"
	"errors"
)

var userInfoKey string = "user_key"

type UserInfo struct {
	Id uint `json:"id"`
}

// 在上下文中传递值
func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userInfoKey, u)
}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userInfoKey).(*UserInfo)
	return u, ok
}

func GetContextUserInfo(ctx context.Context) (*UserInfo, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("获取用户信息错误")

	}
	return user, nil
}
