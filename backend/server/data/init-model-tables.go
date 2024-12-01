package models

import (
	"comfystack/data/models"
	utils "comfystack/services/database"
	"comfystack/services/logger"
	"context"
	"reflect"
)

func initModelTable(model interface{}) {
	conn := utils.GetConnectionString()
	name := reflect.TypeOf(model).Elem().Name()

	// Droppa la tabella
	_, err :=
		conn.NewDropTable().
			IfExists().
			Model(model).
			Exec(context.Background())
	if err != nil {
		logger.Instance.LogWrite("Drop failed for table " + name)
		logger.Instance.LogWrite(err.Error())
		return
	}

	// Ricrea la tabella.
	_, err = conn.NewCreateTable().
		Model(model).
		Exec(context.Background())

	if err != nil {
		logger.Instance.LogWrite("Creation failed for table " + name)
		logger.Instance.LogWrite(err.Error())
		return
	}

	// Conclusa
	logger.Instance.LogWrite("Creation for table " + name + " successful")
}

// Inizializzazione delle tabelle.
func InitializeDatabase() {
	logger.Instance.LogWrite("Initializing database models...")

	models := []interface{}{
		(*models.Utente)(nil),
		(*models.Token)(nil),
	}

	for _, item := range models {
		initModelTable(item)
	}

	logger.Instance.LogWrite("Finished initializing database models!")
}
