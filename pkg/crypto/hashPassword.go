package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte, cost int) []byte {
	hash, _ := bcrypt.GenerateFromPassword(password, cost)
	return hash
}

func IsValidPassword(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
