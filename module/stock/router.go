package stock

import "stock.tao/module/core"

var group = core.Router.Group("stock")

// init request path to handler
func init() {
	group.POST("/", Handle)
}
