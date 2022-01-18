package tools

import (
	"fmt"

	"github.com/joho/godotenv"
)

//Env Parser
func EnvParser() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Print(errEnv.Error())
	}
}
