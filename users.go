package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// DB Connection
var DB *gorm.DB
var err error

// User represents a user with name and email
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

// InitialMigration setups the DB
func InitialMigration(DBDriver, DBHost, DBPort, DBUser, DBPassword, DBName string) {
	URI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)
	DB, err = gorm.Open(DBDriver, URI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DB")
	}
	DB.AutoMigrate(&User{})
}

// AllUsers GET all users
func AllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// OneUser GET a single user
func OneUser(w http.ResponseWriter, r *http.Request) {
	var user User
	vars := mux.Vars(r)
	id := vars["id"]
	DB.First(&user, id)
	json.NewEncoder(w).Encode(user)
}

// NewUser POST new user
func NewUser(w http.ResponseWriter, r *http.Request) {
	var user User
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

// DeleteUser DELETE a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user User
	vars := mux.Vars(r)
	id := vars["id"]
	DB.Where("id = ?", id).Delete(&user)
	json.NewEncoder(w).Encode(&user)
}

// UpdateUser PUT a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	vars := mux.Vars(r)
	id := vars["id"]
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	DB.Model(&user).Where("id = ?", id).Updates(user)
	DB.First(&user, id)
	json.NewEncoder(w).Encode(user)
}
