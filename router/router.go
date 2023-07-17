package router

import (
	. "eden/router/controllers"
	. "eden/serializer"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	user := v1.Group("user")
	{
		user.POST("/login", Bind(Login, binding.JSON))
		user.GET("/:id", Bind(GetUserById, binding.Uri))

	}

	return r

}
