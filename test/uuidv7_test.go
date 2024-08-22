package test

import (
	"testing"

	"github.com/samborkent/uuid"
)

func TestUUIDV7(t *testing.T) {
	count := 0
	N := 1000000

	for range 1000000 {
		uuid1 := uuid.NewV7()
		uuid2 := uuid.NewV7()

		after, _ := uuid1.After(uuid2)

		if after {
			count++
		}
	}

	t.Errorf("count not zero: %d, percentage %f", count, (float64(count)/float64(N))*100)
}

func BenchmarkUUIDV7(t *testing.B) {
	for range t.N {
		_ = uuid.NewV7()
	}
}

func BenchmarkUUIDV7CreationTime(t *testing.B) {
	uuidV7 := uuid.NewV7()

	for range t.N {
		_ = uuidV7.CreationTime()
	}
}

func BenchmarkUUIDV7Short(t *testing.B) {
	uuidV7 := uuid.NewV7()

	for range t.N {
		_ = uuidV7.Short()
	}
}

func BenchmarkUUIDV7String(t *testing.B) {
	uuidV7 := uuid.NewV7()

	for range t.N {
		_ = uuidV7.String()
	}
}

func BenchmarkUUIDV7String2(t *testing.B) {
	for range t.N {
		_ = uuid.NewV7().String()
	}
}

func BenchmarkUUIDV7Version(t *testing.B) {
	uuidV7 := uuid.NewV7()

	for range t.N {
		_ = uuidV7.Version()
	}
}
