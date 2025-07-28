package main

import (
	"github.com/fromcode/to-do-app/src/config"
	"github.com/fromcode/to-do-app/src/routes"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)

	// jalankan seluruh routes
	routes.Routes()
}
