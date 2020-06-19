package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	godotenv.Load()
	port := os.Getenv("SERVER_PORT")
	// Server Listener
	startAPI(port)
}

func startAPI(port string) {
	server := Server{}
	server.init()
	server.Run(port)
}
