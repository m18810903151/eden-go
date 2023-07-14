package serializer

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

// Response 自定义相应结构体
type Response struct {
	Code int
	Msg  string
	Data any
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

// Bind 调用路由方法并绑定实体
func Bind[T any](call func(*gin.Context, *T) Response, bind ...binding.Binding) func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error

		var instance = new(T)
		if len(bind) > 0 {
			err = c.ShouldBindWith(instance, bind[0])
		} else {
			err = c.ShouldBind(instance)
		}

		if err != nil {
			c.JSON(http.StatusOK, Fail(err.Error()))
			return
		}

		c.JSON(http.StatusOK, call(c, instance))
	}
}
