package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Database represents DB configuration
type Database struct {
	Driver, Host, Port, User, Password, Name string
}

// config get values from .env file and populates the Database obj
func (db *Database) config() {
	// Setters
	db.Driver = os.Getenv("DB_DRIVER")
	db.Host = os.Getenv("DB_HOST")
	db.Port = os.Getenv("DB_PORT")
	db.User = os.Getenv("DB_USER")
	db.Password = os.Getenv("DB_PASSWORD")
	db.Name = os.Getenv("DB_NAME")
}

// getUri generates a connection string based on the Database obj
func (db *Database) getURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db.User, db.Password, db.Host, db.Port, db.Name)
}

// init attempts to connect to the db
func (db *Database) init() (*gorm.DB, error) {
	var connection, err = gorm.Open(db.Driver, db.getURI())
	if err != nil {
		fmt.Printf("✖ Cannot connect to %s database\n", db.Driver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("⚡ We are connected to the %s database\n", db.Driver)
	}
	return connection, err
}
