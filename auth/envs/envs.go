package envs

var ServerEnvs Envs

type Envs struct {
	POSTGRES_PASSWORD string
	POSTGRES_USER     string
	POSTGRES_PORT     string
	JWT_SECRET        string
	AUTH_PORT         string
	POSTGRES_NAME     string
	POSTGRES_HOST     string
	POSTGRES_USE_SSL  string
}

func LoadEnvs() error {
	ServerEnvs.JWT_SECRET = ""

	ServerEnvs.POSTGRES_USER = "arthur"
	ServerEnvs.POSTGRES_PASSWORD = "arthur"
	ServerEnvs.POSTGRES_PORT = "9203"
	ServerEnvs.POSTGRES_NAME = "postgres"
	ServerEnvs.POSTGRES_HOST = "localhost"
	ServerEnvs.POSTGRES_USE_SSL = "disable"
	ServerEnvs.AUTH_PORT = "9204"

	return nil
}
