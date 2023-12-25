package bootstrap

import (
	"dataspace/db"
	"dataspace/db/models"
)

func RunMigrations() {
	db := db.GetConnection()
	db.AutoMigrate(&models.User{})
}
