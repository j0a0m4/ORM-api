package main

import (
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	godotenv.Load()
	// Server Listener
	startAPI(":8080")
}

func startAPI(port string) {
	server := Server{}
	server.init()
	server.Run(port)
}
