package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var engine *gin.Engine = gin.Default()
	engine.GET("version", func(ctx *gin.Context) {
		ctx.String(http.StatusAccepted, "1.0.0")
	})
	engine.Run("localhost:8000")
}
