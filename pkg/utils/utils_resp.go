package utils

import (
	"TodoList_demo/pkg/status"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status int         `json:"status,omitempty"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

func RespError(ctx *gin.Context, err error, data string, code ...int) *Response {
	status_code := status.Error
	if code != nil {
		status_code = code[0]
	}
	return &Response{
		Status: status_code,
		Data:   data,
		Msg:    status.TransStatus(status_code),
		Error:  err.Error(),
	}
}

func RespSuccess(ctx *gin.Context, data interface{}, code ...int) *Response {
	status_code := status.Success
	if code != nil {
		status_code = code[0]
	}
	if data == nil {
		data = "操作成功"
	}

	return &Response{
		Status: status_code,
		Data:   data,
		Msg:    status.TransStatus(status_code),
		Error:  "",
	}
}
