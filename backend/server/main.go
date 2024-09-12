package main

import (
	"comfystack/connections/utils"
	"comfystack/constants"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const environFileName string = "env.json"

type EnvironmentVariablesFileStructure struct {
	Dbconn      []utils.PostgresqlConnString `json:"dbconn"`
	SiteOptions utils.SiteConnectionType     `json:"siteopts"`
}

func pErrorAndExit(err string) {
	fmt.Errorf(err)
	os.Exit(1)
}

func readOsEnvironFile() EnvironmentVariablesFileStructure {
	jsonEnvFile, err := os.Open(environFileName)
	if err != nil {
		pErrorAndExit("File variabili d'ambiente assente. Inserire un file di nome '" + environFileName + "' accanto all'eseguibile.")
	}

	buff, err := io.ReadAll(jsonEnvFile)
	if err != nil {
		pErrorAndExit("Errore lettura file variabili d'ambiente.")
	}

	var variables EnvironmentVariablesFileStructure
	json.Unmarshal(buff, &variables)
	fmt.Printf("variables: %v\n", variables)
	defer jsonEnvFile.Close()
	return variables
}

func initialize() {
	// Registro variabili d'ambiente
	envs := readOsEnvironFile()
	utils.RegisterNewSingleton(&envs, constants.Vars)
}

func initEngine(eng *gin.Engine) {

}

func main() {
	initialize()
	var engine *gin.Engine = gin.Default()
	engine.GET("/version", func(ctx *gin.Context) {
		ctx.String(http.StatusAccepted, "1.0.0")
	})
	engine.Run("localhost:8000")
}
