package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	DATABASE_URI := "golang:golang@/devbook?charset=utf8&parseTime=true&loc=Local"

	db, error := sql.Open("mysql", DATABASE_URI)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		return nil, error
	}

	fmt.Println("Connection established.")

	return db, nil
}
