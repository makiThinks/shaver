package domain

import "github.com/google/uuid"

func GenNextId() string {
	return uuid.New().String()
}
