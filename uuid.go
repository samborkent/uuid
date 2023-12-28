package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"sync"
	"time"
	"unsafe"

	xsr256pp "github.com/samborkent/uuid/xrs256pp"
)

type UUID [16]byte

func ToUUID(bytes []byte) (UUID, error) {
	if err := IsValid(bytes); err != nil {
		return UUID{}, fmt.Errorf("invalid uuid: %w", err)
	}

	return UUID(bytes), nil
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
	uuid := make([]byte, 16)

	// Put unix nano timestamp into first 64 bits// Write variant bits into bits 65 and 66
	uuid[8] = (uuid[8] & 0b10111111) | 0b10000000
	binary.BigEndian.PutUint64(uuid[:8], uint64(time.Now().UnixNano()))

	// Set version 8 bits into bits 49 to 52
	uuid[6] = (uuid[6] & 0b10001111) | 0b10000000

	// Set last 64 bits to a pseudo-random number determined by xoshiro256++ algorithm
	binary.BigEndian.PutUint64(uuid[8:], xsr256pp.Next())

	// Write variant bits into bits 65 and 66
	uuid[8] = (uuid[8] & 0b10111111) | 0b10000000

	return UUID(uuid)
}

// Check if a UUID was created after another UUID
func (uuid UUID) After(other UUID) (bool, error) {
	if uuid.Version() != other.Version() {
		return false, ErrCompareVersions
	}

	switch uuid.Version() {
	case Version4:
		return false, ErrNotTimeOrdered
	case Version7:
		if uuid.creationTimeV7().Equal(other.creationTimeV7()) {
			return uuid[6] > other[6], nil
		}

		return uuid.creationTimeV7().After(other.creationTimeV7()), nil
	case Version8:
		return uuid.creationTimeV8().After(other.creationTimeV8()), nil
	default:
		return false, ErrInvalidVersion
	}
}

// Check if a UUID was created before another UUID
func (uuid UUID) Before(other UUID) (bool, error) {
	if uuid.Version() != other.Version() {
		return false, ErrCompareVersions
	}

	switch uuid.Version() {
	case Version4:
		return false, ErrNotTimeOrdered
	case Version7:
		if uuid.creationTimeV7().Equal(other.creationTimeV7()) {
			return uuid[6] < other[6], nil
		}

		return uuid.creationTimeV7().Before(other.creationTimeV7()), nil
	case Version8:
		return uuid.creationTimeV8().Before(other.creationTimeV8()), nil
	default:
		return false, ErrInvalidVersion
	}
}

// Get the version of the UUID
func (uuid UUID) Version() Version {
	return Version(uuid[6] >> 4)
}

func (uuid UUID) CreationTime() (time.Time, error) {
	switch uuid.Version() {
	case Version4:
		return time.Time{}, ErrNotTimeOrdered
	case Version7:
		return uuid.creationTimeV7(), nil
	case Version8:
		return uuid.creationTimeV8(), nil
	default:
		return time.Time{}, ErrInvalidVersion
	}
}

// Last 12 characters of UUID
// Useful for logging
func (uuid UUID) Short() string {
	buf := make([]byte, 12)

	hex.Encode(buf, uuid[10:])

	return *(*string)(unsafe.Pointer(&buf))
}

// Hyphen delimited string representation of UUID
func (uuid UUID) String() string {
	buf := make([]byte, 36)

	encodeHex(buf, uuid)

	return *(*string)(unsafe.Pointer(&buf))
}
