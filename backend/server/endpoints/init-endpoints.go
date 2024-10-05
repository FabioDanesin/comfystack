package endpoints

import (
	"github.com/gin-gonic/gin"

	sys_routes "comfystack/endpoints/sys"
)

type Endpoint = func(ctx *gin.Context)

func InitializeEndpoints(engine *gin.Engine) {
	engine.Group("/api")
	{
		sys_routes.InitializeSystemEndpoints(engine)
	}
}
