package dao

import "time"

// User ==> user model map to table user
type User struct {
	ID         uint64    `xorm:"id"`
	Username   string    `xorm:"username"`
	Password   string    `xorm:"password"`
	Email      string    `xorm:"email"`
	Nickname   string    `xorm:"nickname"`
	CreateTime time.Time `xorm:"create_time"`
	UpdateTime time.Time `xorm:"update_time"`
	Delete     bool      `xorm:"delete"`
}

// TableName ==> Table Name definition
func (u *User) TableName() string {
	return "st_user"
}

// UserPtr ==> user condition
type UserPtr struct {
	ID         *uint64    `column:"id"`
	Username   *string    `column:"username"`
	Password   *string    `column:"password"`
	Email      *string    `column:"email"`
	Nickname   *string    `column:"nickname"`
	CreateTime *time.Time `column:"create_time"`
	UpdateTime *time.Time `column:"update_time"`
	Delete     *bool      `column:"delete"`
}

// Condition ==> implement SQLCondition interfae
func (userPtr *UserPtr) Condition() map[string]interface{} {
	return Condition(userPtr)
}

// SelectUser ==> Select User by Condition
func SelectUser(userPtr *UserPtr) *User {
	user := new(User)
	if _, err := engine.AllCols().Where(userPtr.Condition()).Get(user); err != nil {
		panic(err)
	}
	return user
}

// ExistUser ==> Exist User by Condition
func ExistUser(userPtr *UserPtr) (exist bool) {
	var err error
	user := new(User)
	if exist, err = engine.AllCols().Where(userPtr.Condition()).Exist(user); err != nil {
		panic(err)
	}
	return exist
}

// InsertUser ==> Insert User into user table
func InsertUser(u *User) uint64 {
	if _, err := engine.Insert(u); err != nil {
		panic(err)
	}
	return u.ID
}

// UpdateUser ==> Update User
func UpdateUser(u *User) {
	if _, err := engine.Update(u); err != nil {
		panic(err)
	}
}

// DeleteUser ==> delete user
func DeleteUser(u *User) {
	user := new(User)
	if _, err := engine.Delete(user); err != nil {
		panic(err)
	}
}
