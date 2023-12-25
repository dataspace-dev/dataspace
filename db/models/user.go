package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	Username    string
	Email       string
	Password    string
	Connections []Conection `gorm:"foreignKey:UserID"`
}
