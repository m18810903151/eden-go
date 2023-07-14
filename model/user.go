package model

import (
	. "eden/conf"
	"time"
)

type User struct {
	Id       float64   `json:"id"`
	Name     string    `json:"name"`
	Nickname string    `json:"nickname"`
	Age      string    `json:"age"`
	Gender   byte      `json:"gender"`
	Birthday time.Time `json:"birthday"`
	Avatar   string    `json:"avatar"`
	Deleted  byte      `json:"deleted"`
}

func QueryUserById(id int64) (User, error) {
	var usr User
	first := DB.First(&usr, id)
	if first.Error != nil {
		return usr, first.Error
	}
	return usr, nil
}
