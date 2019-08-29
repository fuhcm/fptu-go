package main

import (
	"boilerplate/db"
	"boilerplate/server"

	_ "github.com/go-sql-driver/mysql" // MySQL Driver
)

func main() {
	db.InitDB()
	defer db.DB.Close()

	server := server.NewServer()
	server.Run(":5000")
}
