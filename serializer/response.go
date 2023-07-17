package serializer

import (
	"github.com/gin-gonic/gin"
)

// Response 自定义相应结构体
type Response struct {
	Code      int
	Msg       string
	Data      any
	RequestId string
}

const (
	success = 0
	fail    = -1
)

// Ok 自定义成功返回
func Ok(data any) Response {
	return Response{
		Code: success,
		Msg:  "成功",
		Data: data,
	}
}

// Fail 自定义失败返回
func Fail(msg string) Response {
	return Response{
		Code: fail,
		Msg:  msg,
		Data: nil,
	}
}

// With 自定义路由处理方法
func With(call func(*gin.Context) Response) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, call(c))
	}
}
