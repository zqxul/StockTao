package user

import (
	"time"

	"stock.tao/dao"
	"stock.tao/util"
)

func init() {

}

// UsernameExist ==> check if the username already exists
func UsernameExist(username string) bool {
	userCondition := dao.UserCondition{Username: &username}
	return dao.Exist(&userCondition)
}

// CreateUser ==> create new user
func CreateUser(username, password, email, nickname string) uint64 {
	user := &dao.User{
		ID:         util.NextID(),
		Username:   username,
		Password:   password,
		Salt:       util.Salt(32),
		Email:      email,
		Nickname:   nickname,
		CreateTime: time.Now(),
		UpdateTime: time.Time{},
		Delete:     false,
	}
	return dao.InsertUser(user)
}

// VerifyUser ==> verify user
func VerifyUser(username, password string) bool {
	userCondition := dao.UserCondition{
		Username: &username,
	}
	user := dao.SelectOne(&userCondition)
	if user == nil {
		return false
	}
	return user.Password == util.Encrypt([]byte(user.Salt), []byte(password))
}
