package controllers

import (
	"eden/model"
	. "eden/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Id struct {
	Id int64 `json:"id" form:"id"  binding:"required"`
}

func GetUserById(c *gin.Context, id *Id) Response {

	user, err := model.QueryUserById(id.Id)
	if err != nil {
		return Fail("用户不存在")
	}
	return Ok(user)
}

// LoginParam 登录参数
type LoginParam struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func Login(c *gin.Context, p *LoginParam) Response {
	fmt.Println(p)
	return Ok("Ok")

}
