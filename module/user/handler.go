package user

import "github.com/gin-gonic/gin"

func init() {

}

// Handle function
func Handle(ctx *gin.Context) {

}

// Register function
func Register(ctx *gin.Context) {
	request := new(RegisterRequest)
	ctx.ShouldBindJSON(request)

}

// Login ==> Login function
func Login(ctx *gin.Context) {
	ctx.ShouldBindJSON(&LoginRequest{})
}
