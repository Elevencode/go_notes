package server

import (
	"go_notes/database"
	"go_notes/envs"
	"log"
)

// Init env and DBs.
func InitServer() {
	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		log.Fatal("Init ENV error: ", errEnvs)

	} else {

		log.Println("Init ENV success")
	}

	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("DB connection error: ", errDatabase)
	} else {
		log.Println("DB connection success")
	}
}

// Init routes and run server.
func StartServer() {
	InitRoutes()
}
