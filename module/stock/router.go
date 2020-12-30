package stock

import "stock.tao/common"

var group = common.Router.Group("stock")

// init request path to handler
func init() {
	group.POST("/", Handle)
}
