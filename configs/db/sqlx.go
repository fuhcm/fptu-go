package db

import (
	"os"

	"github.com/jmoiron/sqlx"
)

// DB ...
var DB *sqlx.DB

// InitSQLXDB ...
func InitSQLXDB() {
	dbConn, err := sqlx.Connect("mysql", os.Getenv("DB"))
	if err != nil {
		panic(err.Error())
	}

	DB = dbConn
}
