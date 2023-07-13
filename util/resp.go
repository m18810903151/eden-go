package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response code 自己定义的状态码
// msg 返回的信息
// data 返回的内容
func Response(context *gin.Context, httpStatus int, code int, msg string, data interface{}) {
	context.JSON(httpStatus, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func ResponseOK(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"code": "0",
		"msg":  "Success",
		"data": data,
	})
}

func ResponseFail(context *gin.Context, msg string) {
	context.JSON(http.StatusOK, gin.H{
		"code": "0",
		"msg":  msg,
		"data": nil,
	})
}
