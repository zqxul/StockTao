package common

import (
	"github.com/gin-gonic/gin"
)

// Router => the global router for register request path
var Router = gin.Default()

func init() {
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())
}
