package util

import (
	"github.com/google/uuid"
)

func GenUuidV4() string {
	uu, _ := uuid.NewRandom()
	return uu.String()
}
