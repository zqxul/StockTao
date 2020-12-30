package main

import (
	"stock.tao/module"
	_ "stock.tao/module/user"
)

func init() {

}

func main() {
	module.Router.Run(":8080")
}
