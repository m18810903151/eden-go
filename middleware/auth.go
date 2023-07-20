package middleware

import (
	"eden/serializer"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CheckUserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
		accessToken := c.Request.Header.Get("access_token")
		if accessToken == "" {
			c.JSON(http.StatusOK, serializer.Fail("Token不存在"))
			c.Abort()
			return
		}
		//此处可以配合jwt解析token
		userId, _ := strconv.Atoi(accessToken)
		c.Set("userId", userId)
		c.Next() // 后续的处理函数可以用过c.Get("userId")来获取当前请求的用户信息

	}
}
