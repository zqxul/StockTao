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

// RunHTTPServer ==> run http server
func RunHTTPServer(port int) {
	Router.Run(":" + strconv.Itoa(port))
}
