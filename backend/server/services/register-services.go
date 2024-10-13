package services

import (
	envvars "comfystack/services/env-vars"
	"comfystack/services/logger"
)

// Registro servizi.
func InitServices() {
	envvars.Register_env_vars()
	logger.Create_default_logger()
}
