package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-pass")

func createToken(userid string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username" : userid,
		"expiry" : time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)

	if err != nil{
		return "", err
	}
	return tokenString, nil
}

func verifyToken(tokenString string) error{
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		return secretKey, nil
	})

	if err != nil{
		return err
	}

	if !token.Valid{
		return fmt.Errorf("Invalid Token")
	}
	return nil
}