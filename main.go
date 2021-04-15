package main

import (
	// _ "github.com/fullstorydev/grpcui/cmd/grpcui"
	_ "github.com/go-sql-driver/mysql"
	"stock.tao/module/core"
	_ "stock.tao/module/monitor"
	_ "stock.tao/module/user"
)

var port, grpcPort int = 8080, 8081

func init() {

}

func main() {
	// core.RunHTTPServer(port)
	core.RunGRPCServer(grpcPort)
}
