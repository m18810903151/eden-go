package serializer

import (
	"eden/common"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

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
			resp := Fail(err.Error())
			resp.TraceId = c.GetString(common.TraceKey)
			c.JSON(http.StatusOK, resp)
			return
		}
		resp := call(c, instance)
		resp.TraceId = c.GetString(common.TraceKey)
		c.JSON(http.StatusOK, resp)
	}
}

// BindUri 调用路由上面带参数绑定
func BindUri[T any](call func(*gin.Context, *T) Response) func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		var instance = new(T)
		err = c.ShouldBindUri(instance)
		if err != nil {
			resp := Fail(err.Error())
			resp.TraceId = c.GetString(common.TraceKey)
			c.JSON(http.StatusOK, resp)
			return
		}
		resp := call(c, instance)
		resp.TraceId = c.GetString(common.TraceKey)
		c.JSON(http.StatusOK, resp)
	}
}
