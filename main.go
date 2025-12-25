package main

import (
	// "encoding/json"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	db := connectDB()
	defer db.Close()
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("POST /register", register(db))
	mux.HandleFunc("POST /login", login(db))
	fmt.Println("Server running at port 3000")
	http.ListenAndServe(":3000", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from root")
}

func login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "" || password == "" {
			http.Error(w, "missing credentials", http.StatusBadRequest)
		}

		var user User

		user, err := checkUserByEmail(db, email)

		if err != nil {
			w.Write([]byte("email not registered"))
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if err != nil {
			w.Write([]byte("password not matched"))
			return
		}
		var url string
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	}
}

func register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// name email password

		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if name == "" || email == "" || password == "" {
			http.Error(w, "missing fields", http.StatusBadRequest)
			return
		}
		_, err := checkUserByEmail(db, email)

		if err == nil {
			// fmt.Fprint(w, "Email already registered")
			http.Error(w, "User already exists", http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14) //transforms the string into a byte
		if err != nil {
			http.Error(w, "error hashing password", http.StatusInternalServerError)
			return
		}
		password = string(hashedPassword)

		data := User{
			Name:     name,
			Email:    email,
			Password: password,
			Active:   true,
		}
		insertUser(db, data)
		// fmt.Fprintf(w, "User successfully created")
		w.Write([]byte("Users succesfully registered"))
	}
}
