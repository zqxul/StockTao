package dao

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"stock.tao/util"
)

// User ==> user model map to table user
type User struct {
	ID         uint64    `xorm:"id"`
	Username   string    `xorm:"username"`
	Password   string    `xorm:"password"`
	Salt       string    `xorm:"salt"`
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

// UserCondition ==> user condition
type UserCondition struct {
	ID         *uint64    `column:"id"`
	Username   *string    `column:"username"`
	Password   *string    `column:"password"`
	Salt       *string    `column:"salt"`
	Email      *string    `column:"email"`
	Nickname   *string    `column:"nickname"`
	CreateTime *time.Time `column:"create_time"`
	UpdateTime *time.Time `column:"update_time"`
	Delete     *bool      `column:"delete"`
}

// Build ==> implement SQLCondition interfae
func (userCondition UserCondition) Build() map[string]interface{} {
	c := make(map[string]interface{})
	t, v := reflect.TypeOf(userCondition), reflect.Indirect(reflect.ValueOf(userCondition))
	if t.NumField() == 0 || v.NumField() == 0 {
		return c
	}
	for i := 0; i < t.NumField(); i++ {
		if v.Field(i).IsNil() {
			continue
		}
		f := t.Field(i)
		if column, ok := f.Tag.Lookup("column"); ok {
			c[column] = util.GetValue(v.Field(i))
		}
	}

	bs, _ := json.Marshal(c)
	fmt.Printf(string(bs))
	return c
}

// Exist ==> Select one User by Condition
func Exist(userCondition *UserCondition) bool {
	exist, err := engine.Table("st_user").AllCols().Where(userCondition.Build()).Exist()
	if err != nil {
		panic(err)
	}
	return exist
}

// SelectOne ==> Select one User by Condition
func SelectOne(userCondition *UserCondition) *User {
	users := make([]*User, 0)
	if err := engine.Table("st_user").AllCols().Where(userCondition.Build()).Find(users); err != nil {
		panic(err)
	}
	if len(users) == 0 {
		return nil
	}
	return users[0]
}

// SelectList ==> Select users by condition
func SelectList(userCondition *UserCondition) []*User {
	users := make([]*User, 0)
	if err := engine.AllCols().Where(userCondition.Build(), make([]interface{}, 0)).Find(users); err != nil {
		panic(err)
	}
	return users

}

// ExistUser ==> Exist User by Condition
func ExistUser(userCondition *UserCondition) (exist bool) {
	var err error
	// user := new(User)
	if exist, err = engine.AllCols().Where(userCondition.Build()).Exist(); err != nil {
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
