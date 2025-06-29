package utils

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateSessionId() string {
	sessionId := uuid.New().String()
	return sessionId
}

func GenerateHashedPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}