package model

import (
	. "eden/common"
	"time"
)

type User struct {
	Id         uint      `json:"id"`
	Name       string    `json:"name"`
	Nickname   string    `json:"nickname"`
	Phone      string    `json:"phone"`
	Password   string    `json:"password"`
	Salt       string    `json:"salt"`
	Age        string    `json:"age"`
	Gender     byte      `json:"gender"`
	Birthday   time.Time `json:"birthday"`
	Avatar     string    `json:"avatar"`
	UserStatus byte      `json:"userStatus"`
}

func LoginByPwd() {

}

func QueryUserByPhone(phone string) (User, error) {
	var usr User
	first := DB.Where("phone=?", phone).First(&usr)
	if first.Error != nil {
		return usr, first.Error
	}
	return usr, nil
}

func QueryUserById(id int64) (User, error) {
	var usr User
	first := DB.First(&usr, id)
	if first.Error != nil {
		return usr, first.Error
	}
	return usr, nil
}
