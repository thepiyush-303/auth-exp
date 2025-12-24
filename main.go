package main

import (
	// "encoding/json"
	"fmt"
	// "net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	// "golang.org/x/crypto/bcrypt"
)


type User struct{
	Name string `json: name`
	Email string `json: email`
	Password string `json: password`
	Active bool `json: active`
}

func main(){

	db := connectDB()
	defer db.Close()
	
	createUserTable(db)

	data := User{
		Name: "Piyush3",
		Email: "piyush3@gmail.com",
		Password: "secretpassword2",
		Active: true,
	}

	pk := insertUser(db, data)

	var fetchedUser User

	fetchedUser = fetchUser(db, pk)
	fmt.Printf("User -> %s", fetchedUser.Name)

	// fmt.Printf("Id %d", pk)

	// mux := http.NewServeMux()


	// mux.HandleFunc("/", )
	// mux. HandleFunc("POST /register", createUser)
	// mux.HandleFunc("GET /login", getUser)
	// http.ListenAndServe(":3000", mux)

}
