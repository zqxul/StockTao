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
		ctx.JSON(http.StatusBadRequest, &core.StockTao{http.StatusBadRequest, "Username already exist", nil})
	}
	userID := CreateUser(request.Username, request.Password, request.Email, request.Nickname)
	if userID == nil {
		ctx.JSON(http.StatusInternalServerError, &core.StockTao{http.StatusInternalServerError, "Create user failed", nil})
	}
	return ctx.JSON(http.StatusOK, &core.StockTao{http.StatusOK, "Create user success", strconv.Itoa(userID)})
}

// Login ==> login handler
func Login(ctx *gin.Context) {
	ctx.ShouldBindJSON(&LoginRequest{})
}
