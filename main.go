package main

import (
	"github.com/fromcode/to-do-app/src/config"
	"github.com/fromcode/to-do-app/src/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// jalankan seluruh routes
	routes.Routes(r)

	r.Run(":3333")
}
