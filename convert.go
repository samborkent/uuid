package uuid

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// ToUUID safely converts a byte slice to a UUID with validation.
func ToUUID(bytes []byte) (UUID, error) {
	if err := IsValid(bytes); err != nil {
		return UUID{}, fmt.Errorf("invalid uuid: %w", err)
	}

	return UUID(bytes), nil
}

// StringToUUID safely converts a byte slice to a UUID with validation.
func StringToUUID(uuid string) (UUID, error) {
	if err := IsValidString(uuid); err != nil {
		return UUID{}, fmt.Errorf("invalid uuid: %w", err)
	}

	bytes, err := hex.DecodeString(strings.Replace(uuid, "-", "", 4))
	if err != nil {
		return UUID{}, fmt.Errorf("failed to decode uuid string: %w", err)
	}

	return UUID(bytes), nil
}
