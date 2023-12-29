package uuid

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
	"unicode"
)

// Check if a string is a valid UUID
func IsValidString(uuid string) error {
	newUUID := make([]byte, 0, 32)

	for i, r := range uuid {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			newUUID = append(newUUID, byte(r))
			continue
		}

		if r != '-' {
			return errors.New("uuid string can only contain alphanumeric characters and hyphens")
		}

		if i != 8 && i != 13 && i != 18 && i != 23 {
			return errors.New("uuid string has incorrect hyphen placement")
		}
	}

	// Check if string is the right length
	if len(newUUID) != 32 {
		return errors.New("uuid string has incorrect length")
	}

	buf := make([]byte, 16)

	// Decode string
	_, err := hex.Decode(buf, newUUID)
	if err != nil {
		return fmt.Errorf("hexadecimal encoding of uuid string: %w", err)
	}

	return IsValid(buf)
}

// Check if a byte slice is a valid UUID
func IsValid(uuid []byte) error {
	// Check if correct number of bytes are present
	if len(uuid) != 16 {
		return errors.New("incorrect number of bytes")
	}

	// Check variant bits
	if uuid[8]|0b00111111 != 0b10111111 {
		return errors.New("invalid variant bytes")
	}

	testUUID := UUID(uuid)

	switch testUUID.Version() {
	case Version4:
		return isValidV4(testUUID)
	case Version7:
		return isValidV7(testUUID)
	case Version8:
		return isValidV8(testUUID)
	default:
		return ErrInvalidVersion
	}
}

func isValidV4(uuid UUID) error {
	testUUID := uuid

	testUUID[6] = testUUID[6] & 0b10111111
	testUUID[8] = testUUID[8] & 0b01111111

	if binary.BigEndian.Uint64(testUUID[:]) == 0 {
		return errors.New("uuid v4 should have non-zero random bits")
	}

	return nil
}

func isValidV7(uuid UUID) error {
	var timestampBytes [8]byte

	copy(timestampBytes[:], uuid[:8])

	// Right shift timestamp bytes
	rightShiftTimestamp(timestampBytes[:])

	// Reject UUIDs from the future
	if binary.BigEndian.Uint64(timestampBytes[:]) > uint64(time.Now().UnixMilli()) {
		return errors.New("uuid v7 cannot have a future timestamp")
	}

	var randBytes [8]byte

	copy(randBytes[:], uuid[8:])

	randBytes[0] = randBytes[0] & 0b01111111

	// Check if random bits are filled
	if binary.BigEndian.Uint64(randBytes[:]) == 0 {
		return errors.New("uuid v7 should have non-zero random bits")
	}

	return nil
}

func isValidV8(uuid UUID) error {
	var timestampBytes [8]byte

	copy(timestampBytes[:], uuid[:8])

	// Right shift timestamp bytes
	rightShiftTimestamp(timestampBytes[:])

	timestampBytes[6] = timestampBytes[6] & 0b01111111

	// Reject UUIDs from the future
	if binary.BigEndian.Uint64(timestampBytes[:]) > uint64(time.Now().UnixNano()) {
		return errors.New("uuid v8 cannot have a future timestamp")
	}

	var randBytes [8]byte

	copy(randBytes[:], uuid[8:])

	randBytes[0] = randBytes[0] & 0b01111111

	// Check if random bits are filled
	if binary.BigEndian.Uint64(randBytes[:]) == 0 {
		return errors.New("uuid v8 should have non-zero random bits")
	}

	return nil
}
