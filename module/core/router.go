package core

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Router => the global router for register request path
var Router = gin.Default()

func init() {
	gin.Recovery()
}

// RunServer ==> run server
func RunServer(port int) {
	Router.Run(":" + strconv.Itoa(port))
}
