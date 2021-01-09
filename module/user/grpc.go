package user

import (
	context "context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

var (
	server = grpc.NewServer()

	grpcServiceDesc = grpc.ServiceDesc{
		ServiceName: "PbUser",
		HandlerType: (*UserServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Login",
				Handler:    handleLogin,
			},
			{
				MethodName: "Register",
				Handler:    handleRegister,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "user.proto",
	}
)

type UserServerI interface {
	Login(context.Context, *PbLoginRequest) (*PbLoginResponse, error)
	Register(context.Context, *PbRegisterRequest) (*PbRegisterResponse, error)
}

type UserServer struct{}

func init() {
	server.RegisterService(&grpcServiceDesc, &UserServer{})

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("failed to listen:%v", err)
	}
	if err := server.Serve(listen); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}

func (UserServer) Login(context.Context, *PbLoginRequest) (*PbLoginResponse, error) {
	return nil, nil
}
func (UserServer) Register(context.Context, *PbRegisterRequest) (*PbRegisterResponse, error) {
	return nil, nil
}

func handleRegister(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {

	return nil, nil
}

func handleLogin(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	return nil, nil
}
