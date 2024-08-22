package uuid

import (
	"encoding/binary"
	"encoding/hex"
	"sync"
	"time"
)

var version4Pool = sync.Pool{
	New: func() any {
		randBuf := make([]byte, 16)
		return &randBuf
	},
}

var version7Pool = sync.Pool{
	New: func() any {
		randBuf := make([]byte, 8)
		return &randBuf
	},
}

// Only has an millisecond accuracy as defined by UUID v7 proposal
func (uuid UUID) creationTimeV7() time.Time {
	// Right shift timestamp bytes
	rightShiftTimestamp(uuid[:8])

	return time.UnixMilli(int64(binary.BigEndian.Uint64(uuid[:8]))).UTC()
}

// Only has an microsecond accuracy as time package does not provide UnixNano function,
// as accuracy of Unix nano timestamp is determined by the OS,
// and it is always greater than nanosecond precision
func (uuid UUID) creationTimeV8() time.Time {
	uuid[6] |= 0b0111_1111

	return time.UnixMicro(int64(binary.BigEndian.Uint64(uuid[:8])) / 1000).UTC()
}

// From github.com/google/uuid
func encodeHex(dst []byte, uuid [16]byte) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}

// Right shift timestamp bytes
func rightShiftTimestamp(uuid []byte) {
	uuid[7] = uuid[5]
	uuid[6] = uuid[4]
	uuid[5] = uuid[3]
	uuid[4] = uuid[2]
	uuid[3] = uuid[1]
	uuid[2] = uuid[0]
	uuid[1] = 0
	uuid[0] = 0
}

func setVariantBits(b byte) byte {
	return (b | 0b1000_0000) & 0b1011_1111
}

func setVersion4Bits(b byte) byte {
	return (b | 0b0100_0000) & 0b0100_1111
}

func setVersion7Bits(b byte) byte {
	return (b | 0b0111_0000) & 0b0111_1111
}

func setVersion8Bits(b byte) byte {
	return (b | 0b1000_0000) & 0b1000_1111
}
