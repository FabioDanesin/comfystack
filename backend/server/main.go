package main

import (
	"comfystack/endpoints"
	"comfystack/services"
	"fmt"

	envvars "comfystack/services/env-vars"
	"os"

	"github.com/gin-gonic/gin"
)

// Initializza la logica dell'applicazione.
// Ritorna l'engine inizializzato e una variabile di ok
func initialize() (*gin.Engine, error) {
	// Inizializzazione servizi.
	services.InitServices()
	// Inizializzazione engine.
	return initEngine(nil), nil
}

// Funzione per inizializzazione del gin.Engine.
// Da tenere separata in caso di logiche custom.
func initEngine(eng *gin.Engine) *gin.Engine {
	if eng == nil {
		eng = gin.Default()
	}
	// Inizializzazione routes.
	endpoints.InitializeEndpoints(eng)

	return eng
}

func main() {
	engine, initStatus := initialize()
	if initStatus != nil {
		os.Exit(1)
	} else {
		siteString := fmt.Sprintf("%s:%s", envvars.Instance.SiteOptions.Root, fmt.Sprint(envvars.Instance.SiteOptions.Port))
		engine.Run(siteString)
	}
}
