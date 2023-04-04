package utils

import (
	"github.com/google/uuid"
)

func GenerateUID() uuid.UUID {
	uid := uuid.New()
	return uid
}
