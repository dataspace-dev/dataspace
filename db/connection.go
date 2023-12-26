package db

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once     sync.Once
	instance *gorm.DB
)

// GetConnection returns a gorm.DB instance
func GetConnection() *gorm.DB {
	once.Do(func() { // Singleton, only one instance of the database connection will be created
		instance = connect()
	})
	return instance
}

// connects to the internal app database and returns a gorm.DB instance
func connect() *gorm.DB {
	dsn := os.Getenv("INTERNAL_DB_DSN")
	if dsn == "" {
		fmt.Println("INTERNAL_DB_DSN is not set")
		os.Exit(1)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
