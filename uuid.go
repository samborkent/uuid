package uuid

import (
	"crypto/rand"
	"errors"
	"math"
	"sync"
	"time"
)

type UUID [16]byte

type Version uint8

const defaultVersion = Version7

const (
	Version4 Version = 4
	Version7 Version = 7
	Version8 Version = 8
)

var currentVersion = defaultVersion

var ErrInvalidVersion = errors.New("version not supported. must be 4, 7, or 8")

// SetVersion sets the version of the application.
//
// It takes an integer value representing the version and returns an error.
func SetVersion(version int) error {
	newVersion := Version(version)

	if newVersion != Version4 && newVersion != Version7 && newVersion != Version8 {
		return ErrInvalidVersion
	}

	currentVersion = newVersion

	return nil
}

func New() UUID {
	switch currentVersion {
	case Version4:
		return NewV4()
	case Version7:
		return NewV7()
	case Version8:
		return NewV8()
	default:
		return NewV7()
	}
}

var version4Pool = sync.Pool{
	New: func() any {
		randBuf := make([]byte, 16)
		return &randBuf
	},
}

func NewV4() UUID {
	uuid := make([]byte, 16)

	// Get buffer to store random bits from pool
	randBufPtr, _ := version4Pool.Get().(*[]byte)
	randBuf := *randBufPtr

	// Generate 128-bit pseudo-random number
	// Randomness is determined by crypto/rand package
	_, _ = rand.Read(randBuf)

	// Copy random bits to UUID
	copy(uuid, randBuf)

	// Return buffer to pool
	version4Pool.Put(randBufPtr)

	// Set version 4 bits
	uuid[6] = (uuid[6] & 0b01001111) | 0b01000000

	// Set variant bits
	uuid[8] = (uuid[6] & 0b10111111) | 0b10000000

	return UUID(uuid)
}

var version7Pool = sync.Pool{
	New: func() any {
		randBuf := make([]byte, 8)
		return &randBuf
	},
}

func NewV7() UUID {
	uuid := make([]byte, 16)

	ms, milliFraction := math.Modf(1e-6 * float64(time.Now().UnixNano()))
	unixMilli := uint64(ms)

	// Big endian binary encoding on Unix millisecond timestamp into the first 48 bits
	uuid[0] = byte(unixMilli >> 40)
	uuid[1] = byte(unixMilli >> 32)
	uuid[2] = byte(unixMilli >> 24)
	uuid[3] = byte(unixMilli >> 16)
	uuid[4] = byte(unixMilli >> 8)
	uuid[5] = byte(unixMilli)

	sequenceNumber := uint16(4096 * milliFraction)

	// Big endian binary encoding of sequence number into bits 49 to 64
	uuid[6] = byte(sequenceNumber >> 8)
	uuid[7] = byte(sequenceNumber)

	// Set version 7 bits into bits 49 to 52
	uuid[6] = (uuid[6] & 0b01111111) | 0b01110000

	// Get buffer to store random bits from pool
	randBufPtr, _ := version7Pool.Get().(*[]byte)
	randBuf := *randBufPtr

	// Generate 64-bit pseudo-random number
	// Randomness is determined by crypto/rand package
	_, _ = rand.Read(randBuf)

	// Copy random bits to UUID
	copy(uuid[8:], randBuf)

	// Return buffer to pool
	version7Pool.Put(randBufPtr)

	// Write variant bits into bits 65 and 66
	uuid[8] = (uuid[8] & 0b10111111) | 0b10000000

	return UUID(uuid)
}

func NewV8() UUID {
	return UUID{}
}
