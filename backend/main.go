/*
Copyright © 2023 Nixon Kosgei <nkosgey6@gmail.com>
*/
package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nixoncode/out-of-the-abyss/internal/app"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/out_of_the_abyss")
	if err != nil {
		panic("cannot connect to the database")
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	DB = db
}
func main() {
	application := app.New(DB)

	err := application.Start()
	if err != nil {
		panic("Cannot start the application")
	}

}
