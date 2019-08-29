package db

import (
	"github.com/jmoiron/sqlx"
)

// DB ...
var DB *sqlx.DB

// InitDB ...
func InitDB() {
	dbConn, err := sqlx.Connect("mysql", "root:12345678@tcp(127.0.0.1:3306)/go_test?charset=utf8&parseTime=True")
	if err != nil {
		panic(err.Error())
	}

	DB = dbConn
}
