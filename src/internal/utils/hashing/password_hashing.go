package hashing

import (
	"golang.org/x/crypto/bcrypt"
)

// HashAndSaltPassword hashes a given password
func HashAndSaltPassword(password string) (string, error) {
	asBytes := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(asBytes, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// VerifyPassword checks whether two given passwords match
func VerifyPassword(hashedPassword, plainPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)) == nil
}