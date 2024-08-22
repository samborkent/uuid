package uuid

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// Parse safely converts a byte slice to a UUID with validation.
func Parse(bytes []byte) (UUID, error) {
	if err := IsValid(bytes); err != nil {
		return UUID{}, fmt.Errorf("invalid uuid: %w", err)
	}

	return UUID(bytes), nil
}

// FromBytes converts a byte slice to a UUID without validation, it will return a Nil UUID in case of an error.
func FromBytes(bytes []byte) UUID {
	if err := IsValid(bytes); err != nil {
		return Nil
	}

	return UUID(bytes)
}

// ParseString safely converts a byte slice to a UUID with validation.
func ParseString(uuid string) (UUID, error) {
	if err := IsValidString(uuid); err != nil {
		return UUID{}, fmt.Errorf("invalid uuid: %w", err)
	}

	bytes, err := hex.DecodeString(strings.Replace(uuid, "-", "", 4))
	if err != nil {
		return UUID{}, fmt.Errorf("failed to decode uuid string: %w", err)
	}

	return UUID(bytes), nil
}

// FromString converts a byte slice to a UUID without validation, it will return a Nil UUID in case of an error.
func FromString(uuid string) UUID {
	if err := IsValidString(uuid); err != nil {
		return Nil
	}

	bytes, err := hex.DecodeString(strings.Replace(uuid, "-", "", 4))
	if err != nil {
		return Nil
	}

	return UUID(bytes)
}
