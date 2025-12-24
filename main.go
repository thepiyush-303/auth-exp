package main

import (
	// "encoding/json"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	// "golang.org/x/crypto/bcrypt"
)

func main() {

	db := connectDB()
	defer db.Close()
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("POST /register", createUser(db))
	mux.HandleFunc("GET /login", getUser(db))
	http.ListenAndServe(":3000", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from /")
}

func getUser(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r * http.Request){
		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "" || password == ""{
			http.Error(w, "missing credentials", http.StatusBadRequest)
		}

		var verifyEmail bool

		verifyEmail = checkUserByEmail(db, email)

		if !verifyEmail{
			w.Write([]byte("email not registered"))
			return
		}

		var verifyPass = findUserCredentials(db, email, password)

		if !verifyPass{
			w.Write([]byte("password not matched"))
			return
		}
		var url string
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	}
}

func createUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// name email password

		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if name == "" || email == "" || password == ""{
			http.Error(w, "missing fields", http.StatusBadRequest)
			return
		}

		flag := checkUserByEmail(db, email)

		if flag {
			// fmt.Fprint(w, "Email already registered")
			http.Error(w, "User already exists", http.StatusBadRequest)
			return
		}

		data := User{
			Name: name,
			Email: email,
			Password: password,
			Active: true,
		}
		insertUser(db, data)
		// fmt.Fprintf(w, "User successfully created")
		w.Write([]byte("Users succesfully registered"))
	}
}
