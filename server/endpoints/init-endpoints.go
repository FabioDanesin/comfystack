package endpoints

import (
	"github.com/gin-gonic/gin"

	"comfystack/endpoints/api"
	sys_routes "comfystack/endpoints/api/sys"
	template_routes "comfystack/endpoints/templates"
)

func InitializeEndpoints(engine *gin.Engine) {

	// Inizializzazione degli endpoint
	engine.Group("/")
	{
		sys_routes.InitializeSystemEndpoints(engine)
		template_routes.RegisterTemplateEndpoints(engine)
	}

	// Inizializzazione degli endpoint API per uso dal frontend
	api.InitializeApiEndpoints(engine)
}
