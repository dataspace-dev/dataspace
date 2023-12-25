package models

import "gorm.io/gorm"

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