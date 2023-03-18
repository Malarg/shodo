package app

import (
	"shodo/internal/transport"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/ping", transport.PingHandler)
	r.Run()
}
