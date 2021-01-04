package main

import (
	_ "github.com/go-sql-driver/mysql"
	"stock.tao/module/core"
	_ "stock.tao/module/user"
)

func init() {

}

func main() {
	core.Router.Run(":8080")
}
