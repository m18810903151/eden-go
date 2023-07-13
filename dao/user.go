package dao

import (
	db "eden/conf"
	"time"
)

type User struct {
	Id       float64
	Name     string
	Nickname string
	Age      string
	Gender   byte
	Birthday time.Time
	Avatar   string
	Deleted  byte
}

func QueryUserById(id float64) User {
	u := User{Id: id}
	db.DB.Model(&User{}).Last(&u)
	return u
}
