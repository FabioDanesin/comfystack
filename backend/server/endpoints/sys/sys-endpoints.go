package sys_routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ritorna il numero di versione.
func getVersion(ctx *gin.Context) {
	ctx.String(http.StatusOK, "1.0.0")
}

// Registra gli endpoint di sistema
func InitializeSystemEndpoints(eng *gin.Engine) {
	eng.Group("/api")
	{
		eng.GET("/version", getVersion)
	}
}
