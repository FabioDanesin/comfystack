package main

import (
	"comfystack/connections/utils"
	"comfystack/constants"
	"comfystack/endpoints"
	"comfystack/types"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

const environFileName string = "env.json"

type EnvironmentVariablesFileStructure struct {
	Dbconn      []types.PostgresqlConnString `json:"dbconn"`
	SiteOptions types.SiteConnectionType     `json:"siteopts"`
}

func readOsEnvironFile() (EnvironmentVariablesFileStructure, error) {
	jsonEnvFile, err := os.Open(environFileName)
	if err != nil {
		return EnvironmentVariablesFileStructure{}, fmt.Errorf("File variabili d'ambiente assente. Inserire un file di nome '" + environFileName + "' accanto all'eseguibile.")
	}

	buff, err := io.ReadAll(jsonEnvFile)
	if err != nil {
		return EnvironmentVariablesFileStructure{}, fmt.Errorf("Errore lettura file variabili d'ambiente.")
	}

	var variables EnvironmentVariablesFileStructure
	json.Unmarshal(buff, &variables)
	fmt.Printf("variables: %v\n", variables)
	defer jsonEnvFile.Close()
	return variables, nil
}

// Initializza la logica dell'applicazione.
// Ritorna l'engine inizializzato e una variabile di ok
func initialize() (*gin.Engine, error) {
	// Registro variabili d'ambiente.
	envs, err := readOsEnvironFile()
	if err != nil {
		return nil, err
	}
	// Registro servizi.
	utils.RegisterNewSingleton(&envs, constants.Vars)
	// Inizializzazione engine.
	engine := gin.Default()
	initEngine(engine)

	return engine, nil
}

// Funzione per inizializzazione del gin.Engine.
// Da tenere separata in caso di logiche custom.
func initEngine(eng *gin.Engine) {
	// Inizializzazione routes.
	endpoints.InitializeEndpoints(eng)
}

func main() {
	engine, initStatus := initialize()
	if initStatus != nil {
		os.Exit(1)
	} else {
		engine.Run("localhost:8000")
	}
}
