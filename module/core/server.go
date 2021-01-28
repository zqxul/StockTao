package core

import (
	"fmt"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GrpcServer ==> grpc server
var GrpcServer = grpc.NewServer()

func init() {
	reflection.Register(GrpcServer)
}

// RunGRPCServer ==> run grpc server
func RunGRPCServer(port int) {
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Printf("failed to listen:%v", err)
	}
	if err := GrpcServer.Serve(listen); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
