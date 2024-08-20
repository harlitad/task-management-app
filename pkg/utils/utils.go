package utils

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func VerifyPassword(reqPassword string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(reqPassword))

	if err != nil {
		return false, err
	}

	return true, nil

}

func ParseUUID(u string) uuid.UUID {
	parsedUUID, _ := uuid.Parse(u)
	return parsedUUID
}
