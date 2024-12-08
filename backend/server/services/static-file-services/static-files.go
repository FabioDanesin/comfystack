package staticfiles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderHtmlResponse(fileName string, dataObject any) func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, fileName, dataObject)
	}
}

const staticFilesFolder string = "./static/"

func registerStaticFile(engine *gin.Engine, htmlFileName string, staticFileName string) {
	if len(staticFileName) == 0 {
		staticFileName = htmlFileName
	}
	hasRelativePath := htmlFileName[0:1] == "./"
	if !hasRelativePath {
		htmlFileName = "./" + htmlFileName
	}
	engine.StaticFile(htmlFileName, getFromStatic(staticFileName))
}
func getFromStatic(fileName string) string {
	return staticFilesFolder + fileName
}

func InitStaticFileService(engine *gin.Engine) {
	// Inizializzazoine import file statici
	staticFiles := [3]string{
		"htmx.min.js",
		"bulma.css",
		"favicon.ico",
	}

	for _, value := range staticFiles {
		registerStaticFile(engine, value, value)
	}

	// Loader generale per file statici
	engine.LoadHTMLGlob("html/*")
}
