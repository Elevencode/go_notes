package envs

import (
	"os"

	"github.com/joho/godotenv"
)

var ServerEnvs Envs

type Envs struct {
	MONGO_INITDB_ROOT_PASSWORD string
	MONGO_INITDB_ROOT_USERNAME string
	MONGO_INITDB_PORT          string
	MONGO_INITDB_HOST          string
	NOTES_PORT                 string
}

func LoadEnvs() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	ServerEnvs.MONGO_INITDB_HOST = os.Getenv("MONGO_INITDB_HOST")
	ServerEnvs.MONGO_INITDB_PORT = os.Getenv("MONGO_INITDB_PORT")
	ServerEnvs.MONGO_INITDB_ROOT_PASSWORD = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	ServerEnvs.MONGO_INITDB_ROOT_USERNAME = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	ServerEnvs.NOTES_PORT = os.Getenv("NOTES_PORT")

	return nil
}
