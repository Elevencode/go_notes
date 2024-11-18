package server

import (
	"auth/database"
	"auth/envs"
	"auth/models"
	"log"
)

func InitServer() {
	// Init ENVs
	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		log.Fatal("Init ENV error: ", errEnvs)
	} else {
		log.Println("Init ENV success")
	}

	// Init PostgreSQL
	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("DB connection error: ", errDatabase)
	} else {
		log.Println("DB connection success")
		database.DB.AutoMigrate(&models.User{})
	}
}

func StartServer() {
	InitRoutes()
}
