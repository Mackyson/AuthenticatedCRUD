package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secret = "gomi"

func main() {
	fmt.Println("----generate phase----")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(1 * time.Second).Unix(),
		Subject:   "1",
	})
	s, err := token.SignedString([]byte(secret))
	fmt.Printf("%s\n%v\n", s, err)
	fmt.Printf("%s\n", token.Method.Alg())

	fmt.Println("----parse phase----")

	token_parse, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() == "HS256" {
			return getKey(), nil
		} else {
			return nil, fmt.Errorf("Unexpected signing method %s", token.Method.Alg())
		}
	})
	if err == nil {
		fmt.Printf("Good!\n%v", token_parse.Claims)
	} else {
		if err == jwt.ErrInvalidKey {
			fmt.Println("Too Bad!\nInvalid signature")
		} else {
			fmt.Println("OMG!\n%s", err)
		}
	}
}
func getKey() []byte {
	return []byte(secret)
}
