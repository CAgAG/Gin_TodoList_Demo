package status

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400
	NotFoundCode  = 40000
)

var TransStatusCode = map[int]string{
	Success:       "成功",
	Error:         "失败",
	InvalidParams: "非法请求参数",
	NotFoundCode:  "未知状态码",
}

func TransStatus(code int) string {
	if ret, ok := TransStatusCode[code]; ok {
		return ret
	}
	return TransStatusCode[NotFoundCode]
}
