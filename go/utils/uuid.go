package utils

import (
	"github.com/google/uuid"
	"log"
)

func CreateUuid() uuid.UUID {
	var random, err = uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	return random
}
