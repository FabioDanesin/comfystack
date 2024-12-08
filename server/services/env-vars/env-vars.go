package envvars

import (
	"comfystack/types"
	"encoding/json"
	"io"
	"os"
)

var Instance *types.EnvironmentVariablesFileStructure

const environFileName string = "env.json"

func RegisterEnvVars() {
	var variables types.EnvironmentVariablesFileStructure
	jsonEnvFile, _ := os.Open(environFileName)
	buff, _ := io.ReadAll(jsonEnvFile)
	json.Unmarshal(buff, &variables)
	defer jsonEnvFile.Close()
	Instance = &variables
}
