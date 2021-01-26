package user

import (
	context "context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"stock.tao/module/core"
)

func init() {
}

// register ==> register handler
func register(ctx *gin.Context) {
	request := new(RegisterRequest)
	ctx.ShouldBindJSON(request)
	if UsernameExist(request.Username) {
		ctx.JSON(http.StatusBadRequest, &core.StockTao{
			Code: http.StatusBadRequest,
			Msg:  "Username already exist",
			Data: nil,
		})
	}
	userID := CreateUser(request.Username, request.Password, request.Email, request.Nickname)
	if userID == 0 {
		ctx.JSON(http.StatusInternalServerError, &core.StockTao{
			Code: http.StatusInternalServerError,
			Msg:  "Create user failed",
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, &core.StockTao{
		Code: http.StatusOK,
		Msg:  "Create user success",
		Data: strconv.FormatUint(userID, 10),
	})
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

// login ==> login handler
func login(ctx *gin.Context) {
	ctx.ShouldBindJSON(&LoginRequest{})
}

// grpcLogin ==> grpc login handler
func grpcLogin(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	fmt.Printf("-------------")
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
