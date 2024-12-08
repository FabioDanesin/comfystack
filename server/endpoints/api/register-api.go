package api

import (
	auth_routes "comfystack/endpoints/api/auth"

	"github.com/gin-gonic/gin"
)

func InitializeApiEndpoints(engine *gin.Engine) {
	engine.Group("/api")
	{
		auth_routes.InitializeAuthEndpoints(engine)
	}
}
