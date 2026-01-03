package utils

import (
	"github.com/google/uuid"
)

// from string to uuid
func ParseUUID(param string) uuid.UUID {
	id, err := uuid.Parse(param)
	if err != nil {
		return uuid.Nil
	}
	return id
}
