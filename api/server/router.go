package server

import (
	"shunp/api/controller"

	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	r := gin.Default()

	ping := new(controller.PingController)

	r.GET("/ping", ping.GetController)

	// GraphQL controller

	return r
}
