package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(plainPassword),
		bcrypt.DefaultCost,
	)

	return string(bytes), err
}

func CheckPassword(hash, userPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(userPassword))
}
