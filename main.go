package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server := Server{}
	server.init()
	server.configRoutes()
	server.Run(os.Getenv("SERVER_PORT"))
}
