package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// User represents a user with name and email
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

// func (u *User) AllUsers(db *gorm.DB) (*User, error) {
// 	var err error
// 	db.First(&u, id)
// }
