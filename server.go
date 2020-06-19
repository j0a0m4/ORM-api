package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Server holds the DB connection and the router
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (s *Server) init() {
	db := Database{}
	db.config()
	s.DB, _ = db.init()

	s.Router = mux.NewRouter()
	s.configRoutes()
}

func (s *Server) configRoutes() {
	// Users Controllers
	s.Router.HandleFunc("/user", s.AllUsers).Methods("GET")
	s.Router.HandleFunc("/user", s.NewUser).Methods("POST")
	s.Router.HandleFunc("/user/{id}", s.OneUser).Methods("GET")
	s.Router.HandleFunc("/user/{id}", s.UpdateUser).Methods("PUT")
	s.Router.HandleFunc("/user/{id}", s.DeleteUser).Methods("DELETE")
}

// Run starts the server
func (s *Server) Run(port string) {
	// Server Listener Setup
	fmt.Printf("🚀 Server listening @localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, s.Router))
}

// AllUsers GET all users
func (s *Server) AllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	s.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// OneUser GET a single user
func (s *Server) OneUser(w http.ResponseWriter, r *http.Request) {
	var user User
	vars := mux.Vars(r)
	id := vars["id"]
	s.DB.First(&user, id)
	json.NewEncoder(w).Encode(user)
}

// NewUser POST new user
func (s *Server) NewUser(w http.ResponseWriter, r *http.Request) {
	var user User
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	s.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

// DeleteUser DELETE a user
func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user User
	vars := mux.Vars(r)
	id := vars["id"]
	s.DB.Where("id = ?", id).Delete(&user)
	json.NewEncoder(w).Encode(&user)
}

// UpdateUser PUT a user
func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	vars := mux.Vars(r)
	id := vars["id"]
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	s.DB.Model(&user).Where("id = ?", id).Updates(user)
	s.DB.First(&user, id)
	json.NewEncoder(w).Encode(user)
}
