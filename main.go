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

	// createUserTable(db)

	// data := User{
	// 	Name: "Piyush3",
	// 	Email: "piyush3@gmail.com",
	// 	Password: "secretpassword2",
	// 	Active: true,
	// }

	// pk := insertUser(db, data)

	// var fetchedUser User

	// fetchedUser = fetchUser(db, pk)
	// fmt.Printf("User -> %s", fetchedUser.Name)

	// fmt.Printf("Id %d", pk)

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
