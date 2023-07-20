package routers

import (
	. "eden/controller/login"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	// 初始化用户api
	InitLoginRouters(v1)

}
