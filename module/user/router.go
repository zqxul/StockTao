package user

import "stock.tao/common"

var group = common.Router.Group("user")

// init request path to handler
func init() {
	group.POST("/register", Register)
	group.POST("/login", Login)
}
