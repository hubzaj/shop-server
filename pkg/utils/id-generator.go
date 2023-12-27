package utils

import (
	"github.com/google/uuid"
	"log"
)

func CreateNewUUID() *uuid.UUID {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("error while generating id: %s", err)
	}
	return &id
}
