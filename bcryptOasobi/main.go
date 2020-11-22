package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const (
	secret      = "gomi"
	STRETCH_NUM = 5
)

func main() {
	hash, _ := bcrypt.GenerateFromPassword([]byte(secret), STRETCH_NUM)
	fmt.Printf("hash:%s\ntype:%T\n", hash, hash)

	err := bcrypt.CompareHashAndPassword(hash, []byte(secret))
	fmt.Println(err)

	err = bcrypt.CompareHashAndPassword(hash, []byte(secret+"a"))
	fmt.Println(err)

}
