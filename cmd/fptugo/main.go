package main

import (
	"fptugo/configs/db"
	"fptugo/internal/confession"
	"fptugo/internal/server"
	"fptugo/internal/user"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL Driver
	_ "github.com/joho/godotenv/autoload"     // env Load
)

func main() {
	db.InitDB()

	db := db.GetDatabaseConnection()
	defer db.Close()

	migrateDatabase()

	server := server.NewServer()
	server.Run(":" + os.Getenv("PORT"))
}

func migrateDatabase() {
	db := db.GetDatabaseConnection()

	// Migrate the given tables
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&confession.Confession{})
}
