package stock

var group = Router.Group("stock")

// init request path to handler
func init() {
	group.POST("/", Handle)
}
