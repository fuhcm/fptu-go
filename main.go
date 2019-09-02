package main

import (
	"fptugo/db"
	"fptugo/server"
	"os"

	_ "github.com/go-sql-driver/mysql"    // MySQL Driver
	_ "github.com/joho/godotenv/autoload" // env Load
)

func main() {
	db.InitDB()
	defer db.DB.Close()

	server := server.NewServer()
	server.Run(":" + os.Getenv("PORT"))
}
