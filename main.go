package main

import (
	"go_notes/server"
)

func init() {
	server.InitServer()
}

func main() {
	server.StartServer()
}
