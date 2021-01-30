package user

import (
	context "context"
	"errors"
	"strconv"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"stock.tao/module/core"
)

var (
	grpcServiceDesc = grpc.ServiceDesc{
		ServiceName: "User",
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
func (ServerImpl) Login(ctx context.Context, request *PbLoginRequest) (*PbStockTao, error) {
	err := func(request *PbLoginRequest) error {
		if len(request.Username) == 0 {
			return errors.New("Username can not be blank")
		} else if len(request.Password) == 0 {
			return errors.New("Password can not be blank")
		} else if len(request.VerifyCode) == 0 {
			return errors.New("VerifyCode can not be blank")
		}
		return nil
	}(request)
	if err != nil {
		return &PbStockTao{
			Code: int32(codes.DataLoss),
			Msg:  codes.DataLoss.String(),
		}, nil
	}
	if !VerifyUser(request.Username, request.Password) {
		return &PbStockTao{
			Code: int32(codes.PermissionDenied),
			Msg:  codes.PermissionDenied.String(),
		}, nil
	}
	return &PbStockTao{
		Code: int32(codes.OK),
		Msg:  codes.OK.String(),
	}, nil
}

// Register ==> implement Server inteface
func (ServerImpl) Register(ctx context.Context, request *PbRegisterRequest) (r *PbStockTao, e error) {
	if UsernameExist(request.Username) {
		return &PbStockTao{
			Code: int32(codes.AlreadyExists),
			Msg:  codes.AlreadyExists.String(),
		}, nil
	}
	userID := CreateUser(request.Username, request.Password, request.Email, request.Nickname)
	if userID == 0 {
		return &PbStockTao{
			Code: int32(codes.Internal),
			Msg:  codes.Internal.String(),
		}, nil
	}
	data, _ := ptypes.MarshalAny(&PbRegisterResponse{
		UserID: strconv.FormatInt(int64(userID), 10),
	})
	return &PbStockTao{
		Code: int32(codes.OK),
		Msg:  codes.OK.String(),
		Data: data,
	}, nil
}
