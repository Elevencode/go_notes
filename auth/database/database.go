package database

import (
	"auth/envs"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() error {
	env := envs.ServerEnvs
	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		env.POSTGRES_HOST, env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_NAME, env.POSTGRES_PORT, env.POSTGRES_USE_SSL)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return err
	} else {
		DB = db
		return nil
	}
}
