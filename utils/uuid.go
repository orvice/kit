package utils

import (
	"github.com/satori/go.uuid"
)

func NewUUID() string {
	u1 := uuid.Must(uuid.NewV4(), nil)
	return u1.String()
}
