package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	DATABASE_URI := "golang:golang@/devbook?charset=utf8&parseTime=true&loc=Local"

	db, error := sql.Open("mysql", DATABASE_URI)
	if error != nil {
		log.Fatal(error)
	}

	defer db.Close()

	if error = db.Ping(); error != nil {
		log.Fatal(error)
	}

	fmt.Println("Connection established.")

	rows, error := db.Query("SELECT * FROM users")
	if error != nil {
		log.Fatal(error)
	}

	defer rows.Close()

	fmt.Println(rows)
}
