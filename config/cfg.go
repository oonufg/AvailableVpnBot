package cfg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type cfg struct {
	TG_API_KEY string
}

func LoadConfig() *cfg {
	workingDirrectory, _ := os.Getwd()
	error := godotenv.Load()
	if error != nil {
		log.Fatalf("Failed to read env file:\n Dirrectory - %s", workingDirrectory)
	}
	return &cfg{
		TG_API_KEY: pullEnvVariable("api_key"),
	}
}

func pullEnvVariable(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Env variable - %s - not exists", key)
	}
	return val
}
