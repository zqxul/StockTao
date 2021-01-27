package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"stock.tao/module/core"
)

func init() {
}

// register ==> register handler
func register(ctx *gin.Context) {
	request := new(RegisterRequest)
	err := ctx.BindJSON(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &core.StockTao{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	if UsernameExist(request.Username) {
		ctx.JSON(http.StatusBadRequest, &core.StockTao{
			Code: http.StatusBadRequest,
			Msg:  "Username already exist",
			Data: nil,
		})
		return
	}
	userID := CreateUser(request.Username, request.Password, request.Email, request.Nickname)
	if userID == 0 {
		ctx.JSON(http.StatusInternalServerError, &core.StockTao{
			Code: http.StatusInternalServerError,
			Msg:  "Create user failed",
		})
		return
	}
	ctx.JSON(http.StatusOK, &core.StockTao{
		Code: http.StatusOK,
		Msg:  "Register success",
		Data: strconv.FormatUint(userID, 10),
	})
}

// login ==> login handler
func login(ctx *gin.Context) {
	request := new(LoginRequest)
	err := ctx.BindJSON(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &core.StockTao{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	if !VerifyUser(request.Username, request.Password) {
		ctx.JSON(http.StatusOK, &core.StockTao{
			Code: http.StatusBadRequest,
			Msg:  "Username Or Password invalid",
		})
		return
	}
	ctx.JSON(http.StatusOK, &core.StockTao{
		Code: http.StatusOK,
		Msg:  "Login Success",
	})
}
