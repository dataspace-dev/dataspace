package bootstrap

import (
	"dataspace/db"
	db1 "dataspace/db/db"
)

func RunMigrations() {
	db := db.GetConnection()
	db.AutoMigrate(&db1.User{})
	db.AutoMigrate(&db1.Conection{})
}
