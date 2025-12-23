package main

import (
	"database/sql"
	"fmt"
	"log"
)

func connectDB() *sql.DB{
	db, err := sql.Open(
		"pgx",
		"postgres://postgres:heythisismypassword@localhost:5432/",
	)
	if err != nil{
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil{
		log.Fatal(err)
	}
	fmt.Println("working")
	return db
}