package db

import "gorm.io/gorm"

// Conection is a model that stores the database connection information
type Conection struct {
	gorm.Model
	Host    string
	Port    string
	Dbname  string
	User    string
	Pass    string
	SSLMode string
	UserID  uint
}