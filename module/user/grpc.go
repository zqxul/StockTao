package user

import (
	context "context"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/anypb"
	"stock.tao/module/core"
)

var (
	grpcServiceDesc = grpc.ServiceDesc{
		ServiceName: "PbUser",
		HandlerType: (*Server)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Login",
				Handler:    grpcLogin,
			},
			{
				MethodName: "Register",
				Handler:    grpcRegister,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "grpc.proto",
	}
)

func init() {
	core.GrpcServer.RegisterService(&grpcServiceDesc, &ServerImpl{})
}

// grpcLogin ==> grpc login handler
func grpcLogin(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PbLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Server).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Server).Login(ctx, req.(*PbLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// grpcRegister ==> grpc register handler
func grpcRegister(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PbRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Server).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Server).Register(ctx, req.(*PbRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Server ==> server inteface
type Server interface {
	Login(context.Context, *PbLoginRequest) (*PbStockTao, error)
	Register(context.Context, *PbRegisterRequest) (*PbStockTao, error)
}

// ServerImpl ==> server implement
type ServerImpl struct{}

// Login ==> implement Server inteface
func (ServerImpl) Login(context.Context, *PbLoginRequest) (*PbStockTao, error) {
	return &PbStockTao{
		Code: 100,
		Msg:  codes.AlreadyExists.String(),
	}, nil
}

// Register ==> implement Server inteface
func (ServerImpl) Register(ctx context.Context, request *PbRegisterRequest) (r *PbStockTao, e error) {
	if UsernameExist(request.Username) {
		return &PbStockTao{
			Code: int32(codes.AlreadyExists),
			Msg:  codes.AlreadyExists.String(),
			Data: nil,
		}, nil
	}
	userID := CreateUser(request.Username, request.Password, request.Email, request.Nickname)
	if userID == 0 {
		return &PbStockTao{
			Code: int32(codes.Internal),
			Msg:  codes.Internal.String(),
			Data: nil,
		}, nil
	}
	return &PbStockTao{
		Code: int32(codes.OK),
		Msg:  codes.OK.String(),
		Data: &anypb.Any{
			Value: []byte(strconv.FormatUint(userID, 10)),
		},
	}, nil
}
