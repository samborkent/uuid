package test

import (
	"testing"

	"github.com/samborkent/uuid"
)

func TestUUIDV8(t *testing.T) {
	count := 0
	N := 1000000

	for range 1000000 {
		uuid1 := uuid.NewV8()
		uuid2 := uuid.NewV8()

		after, _ := uuid1.After(uuid2)

		if after {
			count++
		}
	}

	t.Errorf("count not zero: %d, percentage %f", count, (float64(count)/float64(N))*100)
}

func BenchmarkUUIDV8(t *testing.B) {
	for range t.N {
		_ = uuid.NewV8()
	}
}

func BenchmarkUUIDV8CreationTime(t *testing.B) {
	uuidV8 := uuid.NewV8()

	for range t.N {
		_ = uuidV8.CreationTime()
	}
}

func BenchmarkUUIDV8Short(t *testing.B) {
	uuidV8 := uuid.NewV8()

	for range t.N {
		_ = uuidV8.Short()
	}
}

func BenchmarkUUIDV8String(t *testing.B) {
	uuidV8 := uuid.NewV8()

	for range t.N {
		_ = uuidV8.String()
	}
}

func BenchmarkUUIDV8String2(t *testing.B) {
	for range t.N {
		_ = uuid.NewV8().String()
	}
}

func BenchmarkUUIDV8Version(t *testing.B) {
	uuidV8 := uuid.NewV8()

	for range t.N {
		_ = uuidV8.Version()
	}
}
