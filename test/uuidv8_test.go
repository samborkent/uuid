package test

import (
	"testing"

	"github.com/samborkent/uuid"
)

func BenchmarkUUIDV8(t *testing.B) {
	for range t.N {
		_ = uuid.NewV8()
	}
}

func BenchmarkUUIDV8CreationTime(t *testing.B) {
	uuidV8 := uuid.NewV8()

	for range t.N {
		_, _ = uuidV8.CreationTime()
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
