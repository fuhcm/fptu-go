package db

import (
	"os"

	"github.com/jmoiron/sqlx"
)

// DB ...
var DB *sqlx.DB

// InitDB ...
func InitDB() {
	dbConn, err := sqlx.Connect("mysql", os.Getenv("DB"))
	if err != nil {
		panic(err.Error())
	}

	DB = dbConn
}
