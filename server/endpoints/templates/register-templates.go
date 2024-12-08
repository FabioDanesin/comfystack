package template_routes

import (
	staticfiles "comfystack/services/static-file-services"

	"github.com/gin-gonic/gin"
)

func RegisterTemplateEndpoints(engine *gin.Engine) *gin.RouterGroup {
	templateEndpoints := engine.Group("/")

	templateEndpoints.GET("/", staticfiles.RenderHtmlResponse("index.html", nil))

	return templateEndpoints
}
