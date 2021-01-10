package core

import (
	"fmt"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

// GrpcServer ==> grpc server
var GrpcServer = grpc.NewServer()

func init() {

}

// RunGrpcServer ==> run grpc server
func RunGrpcServer(port int) {
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Printf("failed to listen:%v", err)
	}
	if err := GrpcServer.Serve(listen); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
