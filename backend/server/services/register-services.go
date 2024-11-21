package services

import (
	envvars "comfystack/services/env-vars"
	"comfystack/services/logger"
)

// Registro servizi.
func InitServices() {
	envvars.RegisterEnvVars()
	logger.CreateDefaultLogger()
}
