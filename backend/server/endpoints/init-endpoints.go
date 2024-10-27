package endpoints

import (
	"github.com/gin-gonic/gin"

	sys_routes "comfystack/endpoints/sys"
)

func InitializeEndpoints(engine *gin.Engine) {
	sys_routes.InitializeSystemEndpoints(engine)
}
