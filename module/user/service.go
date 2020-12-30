package user

import (
	"time"

	"stock.tao/dao"
	"stock.tao/util"
)

func init() {

}

// UsernameExist ==> check if the username already exists
func UsernameExist(username string) (exist bool) {
	userCondition := dao.User{Username: username}
	user := dao.SelectUser(&userCondition)
	if user != nil {
		exist = true
	}
}

// CreateUser ==> create new user
func CreateUser(username, password, email, nickname string) uint64 {
	user := &dao.User{
		ID:         util.NextID(),
		Username:   username,
		Password:   password,
		Email:      email,
		Nickname:   nickname,
		CreateTime: time.Now(),
		UpdateTime: nil,
		Delete:     false,
	}
	return dao.InsertUser(user)
}
