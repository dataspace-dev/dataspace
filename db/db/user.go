package db

import "gorm.io/gorm"

// User is a model that stores the user information
type User struct {
	gorm.Model
	Name        string
	Username    string
	Email       string
	Password    string
	Connections []Conection `gorm:"foreignKey:UserID"`
}
