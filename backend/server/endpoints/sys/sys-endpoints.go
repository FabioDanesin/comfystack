package sys_routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getVersion(ctx *gin.Context) {
	ctx.String(http.StatusAccepted, "1.0.0")
}

func InitializeSystemEndpoints(eng *gin.Engine) []gin.HandlerFunc {
	return []gin.HandlerFunc{getVersion}
}
