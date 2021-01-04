package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"stock.tao/module/core"
)

func init() {
}

// Register ==> register handler
func Register(ctx *gin.Context) {
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
		Data: strconv.Itoa(int(userID)),
	})
}

// Login ==> login handler
func Login(ctx *gin.Context) {
	ctx.ShouldBindJSON(&LoginRequest{})
}
