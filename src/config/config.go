package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	errorENV := godotenv.Load()
	if errorENV != nil {
		panic("Gagal menemukan env file")
	}

	dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbHost, dbName)
	db, errorDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errorDB != nil {
		panic("Gagal tersambung ke database mysql")
	}

	return db
}

// DisconnectDB untuk stop koneksi mysql database
func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Gagal memutuskan koneksi mysql database")
	}
	dbSQL.Close()
}
