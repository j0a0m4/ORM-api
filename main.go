package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// PORT where the server will listen
const PORT = ":8080"

func server() {
	// Setup Environment Variables
	godotenv.Load()
	// Router Setup
	router := mux.NewRouter().StrictSlash(true)
	// Controllers
	router.HandleFunc("/", hello).Methods("GET")
	// Users Controllers
	router.HandleFunc("/user", AllUsers).Methods("GET")
	router.HandleFunc("/user", NewUser).Methods("POST")
	router.HandleFunc("/user/{id}", OneUser).Methods("GET")
	router.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
	// Database Setup
	InitialMigration(os.Getenv("DB_DRIVER"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	defer DB.Close()
	// Server Listener Setup
	fmt.Printf("🚀 Server listening @localhost%s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	server()
}
