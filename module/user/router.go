package user

var group = Router.Group("/user")

// init request path to handler
func init() {
	group.POST("/register", Register)
	group.POST("/login", Login)
}
