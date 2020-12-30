package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
}

// Register ==> register handler
func Register(ctx *gin.Context) {
	request := new(RegisterRequest)
	ctx.ShouldBindJSON(request)
	if UsernameExist(request.Username) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg": "User already exist"
			"data": nil,
		})
	}
	userID := CreateUser(request.Username, request.Password, request.Email, request.Nickname)
	if userID == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg": "Create user failed",
			"data": nil,
		})
	}
	return ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg": "Create user success",
		"data": strconv.Itoa(userID),
	})
}

// Login ==> login handler
func Login(ctx *gin.Context) {
	ctx.ShouldBindJSON(&LoginRequest{})
}
