package server

import (
	"go_notes/database"
	"go_notes/envs"
	"log"
)

// Init env and DBs.
func InitServer() {
	// Init ENVs
	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		log.Fatal("Init ENV error: ", errEnvs)

	} else {
		log.Println("Init ENV success")
	}

	// Init MongoDB
	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("DB connection error: ", errDatabase)
	} else {
		log.Println("DB connection success")
	}

	// Init Redis
	errRedis := database.InitRedis()
	if errRedis != nil {
		log.Fatal("Redis connection error: ", errRedis)
	} else {
		log.Println("Redis connection success")
	}
}

// Init routes and run server.
func StartServer() {
	InitRoutes()
}
