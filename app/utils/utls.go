package utils

import (
	"github.com/google/uuid"
)

func GenerateSessionId() string {
	sessionId := uuid.New().String()
	return sessionId
}