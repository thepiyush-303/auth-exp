package main

import (
	"database/sql"
	"fmt"
	"log"
)

func connectDB() *sql.DB {
	db, err := sql.Open(
		"pgx",
		"postgres://postgres:heythisismypassword@localhost:5432/authdb",
	)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func checkUserByEmail(db *sql.DB, email string) (User, error) {
	query := `SELECT ROW FROM TABLE WHERE email = $1`

	var user User

	err := db.QueryRow(query, email).Scan(&user)

	if err != nil {
		log.Fatal(err)
		return user, err
	}
	return user, nil
}

func findUserCredentials(db *sql.DB, email string, password string) bool {
	query := `SELECT password FROM users WHERE email = $1`

	var pass string
	err := db.QueryRow(query, email).Scan(&pass)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return password == pass
}

func fetchUser(db *sql.DB, pk int) User {
	query := `SELECT name, email, password, active FROM users WHERE id = $1`

	var user User

	err := db.QueryRow(query, pk).Scan(&user.Name, &user.Email, &user.Password, &user.Active)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("No row with the id %d", pk)
		}
		log.Fatal(err)
	}
	return user
}

func insertUser(db *sql.DB, data User) int {
	query := `INSERT INTO users	(
		name, email, password, active)
		VALUES ($1, $2, $3, $4) RETURNING id`

	var pk int
	err := db.QueryRow(query, data.Name, data.Email, data.Password, data.Active).Scan(&pk)

	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func createUserTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) UNIQUE NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password VARCHAR(100) NOT NULL,
		active BOOLEAN
	)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
