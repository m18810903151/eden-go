package user

import (
	"eden/common"
	"eden/common/hashid"
	"eden/model"
	. "eden/serializer"
	"eden/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"strings"
	"time"
)

// LoginParam 登录参数
type LoginParam struct {
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// LoginResp 登录响应
type LoginResp struct {
	AccessToken string `json:"accessToken"`
}

// EdenClaims claim
type EdenClaims struct {
	Phone  string `json:"phone"`
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func InitLoginRouters(group *gin.RouterGroup) {
	user := group.Group("user")
	{
		user.POST("login", Bind(Login, binding.JSON))
	}
}

func Login(c *gin.Context, p *LoginParam) Response {
	user, err := model.QueryUserByPhone(p.Phone)
	if err != nil {
		return Fail(common.ACCOUNT_PWD_ERROR)
	}
	newUUID, _ := uuid.NewUUID()
	// 生成盐
	salt := utils.MD5(common.PWD_SALT_KEY + strings.ReplaceAll(newUUID.String(), "-", ""))
	inputPwd := p.Password
	pwd := utils.MD5Salt(inputPwd, salt)
	if pwd != user.Password {
		return Fail(common.ACCOUNT_PWD_ERROR)
	}
	// 校验通过则执行登录逻辑
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, createClaims(user.Phone, hashid.HashID(user.Id, hashid.UserID)))
	ss, err := token.SignedString([]byte(common.LOGIN_SIGN))
	if err != nil {
		return Fail(common.SYSTEM_ERROR)
	}
	return Ok(LoginResp{ss})

}

func createClaims(phone string, userId string) EdenClaims {
	newUUID, _ := uuid.NewUUID()
	id := strings.ReplaceAll(newUUID.String(), "-", "")
	claims := EdenClaims{
		phone,
		userId,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "www.eden.com",
			Subject:   "www.eden.com",
			ID:        id,
			Audience:  []string{"www.eden.com"},
		},
	}
	return claims
}
