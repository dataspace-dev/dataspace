package types

import "gorm.io/gorm"

// Connection is a model that stores the database connection information
type Connection struct {
	gorm.Model
	Host    string
	Port    string
	Dbname  string
	User    string
	Pass    string
	SSLMode string
	UserID  uint
}
