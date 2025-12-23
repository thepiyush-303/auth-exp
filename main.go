package main

import (
	// "encoding/json"
	"net/http"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/jackc/pgx/v5/stdlib"
	)


func main(){

	db := connectDB()
	mux := http.NewServeMux()

	mux.HandleFunc("/", signupHandler(db))
	// mux. HandleFunc("POST /register", createUser)
	// mux.HandleFunc("GET /login", getUser)
	http.ListenAndServe(":3000", mux)

}


func signupHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "" || password == "" {
			http.Error(w, "missing fields", http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec(
			"INSERT INTO users (email, password) VALUES ($1, $2)",
			email,
			string(hashedPassword),
		)
		if err != nil {
			http.Error(w, "user already exists", http.StatusBadRequest)
			return
		}

		w.Write([]byte("signup successful"))
	}
}


// func createUser(w http.ResponseWriter, r *http.Request){

// }

// func getUser(w http.ResponseWriter, r *http.Request){

// }