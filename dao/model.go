package dao

import "time"

// User ==> user model map to table user
type User struct {
	ID         uint64    `gorm:"id"`
	Username   string    `gorm:"username"`
	Password   string    `gorm:"password"`
	Email      string    `gorm:"email"`
	Nickname   string    `gorm:"nickname"`
	CreateTime time.Time `gorm:"create_time"`
	UpdateTime time.Time `gorm:"update_time"`
	Delete     bool      `gorm:"delete"`
}
