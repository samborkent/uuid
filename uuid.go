package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"math"
	mathrand "math/rand/v2"
	"sync/atomic"
	"time"
	"unsafe"
)

type UUID [16]byte

var (
	Max = UUID{
		math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8,
		math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8,
		math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8,
		math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8,
	}
	Nil = UUID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
)

// New returns a UUID based on the current version.
// The current version can be set with SetVersion. It defaults to version 7.
// Alternatively, the explicit functions NewV4, NewV7, and NewV8 can be used.
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

// NewV4 generates a new version 4 UUID.
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
	uuid[6] = setVersion4Bits(uuid[6])

	// Set variant bits
	uuid[8] = setVariantBits(uuid[8])

	return UUID(uuid)
}

var (
	prevUnixMilli      = new(atomic.Uint64)
	prevSequenceNumber = new(atomic.Uint32)
	prevCount          = new(atomic.Uint32)
)

// NewV7 generates a new version 7 UUID.
//
// It generates a UUID based on the current Unix millisecond timestamp and a
// sequence number derived from the fraction to the next millisecond.
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

	sequenceNumber := uint32(4096 * milliFraction)

	// Big endian binary encoding of sequence number into bits 49 to 64
	uuid[6] = byte(sequenceNumber >> 8)
	uuid[7] = byte(sequenceNumber)

	// Set version 7 bits into bits 49 to 52
	uuid[6] = setVersion7Bits(uuid[6])

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

	// Clear variant bits and one extra for overflow.
	uuid[8] |= 0b0001_1111

	// If UUID was created in the same timeframe as previous
	if unixMilli == prevUnixMilli.Swap(unixMilli) {
		// If sequence number of UUID is less than or equal to sequence number of previous UUID
		prevSeqNum := prevSequenceNumber.Swap(sequenceNumber)
		if sequenceNumber <= prevSeqNum {
			// Set sequence number to same as previous
			uuid[6] = byte(prevSeqNum >> 8)
			uuid[7] = byte(prevSeqNum)
			uuid[6] = setVersion7Bits(uuid[6])

			// Set overflow rand bits to same as previous + 1.
			count := prevCount.Swap(uint32(binary.BigEndian.Uint16(uuid[8:10])+1)) + 1
			uuid[8] = byte(count >> 8)
			uuid[9] = byte(count)
		}
	} else {
		// Update previous sequence number if UUID was not created in the same timeframe as previous.
		prevSequenceNumber.Store(sequenceNumber)
	}

	// Write variant bits into bits 65 and 66
	uuid[8] = setVariantBits(uuid[8])

	return UUID(uuid)
}

// NewV8 generates a new V8 UUID.
//
// It inserts the Unix nano timestamp into the first 64 bits. The precision of this timestamp is platform dependend.
// It generates fast pseodo-random number using math/rand/v2. This is NOT crpytographically secure.
func NewV8() UUID {
	uuid := make([]byte, 16)

	// Put unix nano timestamp into first 64 bits
	binary.BigEndian.PutUint64(uuid[:8], uint64(time.Now().UnixNano()))

	// Set version 8 bits into bits 49 to 52
	uuid[6] = setVersion8Bits(uuid[6])

	// Set last 64 bits to a pseudo-random number determined by xoshiro256++ algorithm
	binary.BigEndian.PutUint64(uuid[8:], mathrand.Uint64())

	// Write variant bits into bits 65 and 66
	uuid[8] = setVariantBits(uuid[8])

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
			// Compare sequence numbers if Unix milli timestamps are equal.

			var uuidSequence [2]byte
			copy(uuidSequence[:], uuid[6:8])
			uuidSequence[0] |= 0b0000_1111

			var otherSequence [2]byte
			copy(otherSequence[:], other[6:8])
			otherSequence[0] |= 0b0000_1111

			uuidSeqNum := binary.BigEndian.Uint16(uuidSequence[:])
			otherSeqNum := binary.BigEndian.Uint16(otherSequence[:])

			if uuidSeqNum == otherSeqNum {
				return binary.BigEndian.Uint16(uuid[8:10]) > binary.BigEndian.Uint16(otherSequence[:]), nil
			}

			return binary.BigEndian.Uint16(uuidSequence[:]) > binary.BigEndian.Uint16(otherSequence[:]), nil
		}

		return uuid.creationTimeV7().After(other.creationTimeV7()), nil
	case Version8:
		// This uses Unix nano timestamp, which will never be equal for two UUIDs,
		// because it takes ~50 ns to generate a UUID v8.
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

func (uuid UUID) CreationTime() time.Time {
	switch uuid.Version() {
	case Version7:
		return uuid.creationTimeV7()
	case Version8:
		return uuid.creationTimeV8()
	default:
		return time.Time{}
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
