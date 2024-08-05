package test

import (
	"testing"

	"github.com/samborkent/uuid"
)

func BenchmarkUUIDV4(t *testing.B) {
	for range t.N {
		_ = uuid.NewV4()
	}
}

func BenchmarkUUIDV4Short(t *testing.B) {
	uuidV4 := uuid.NewV4()

	for range t.N {
		_ = uuidV4.Short()
	}
}

func BenchmarkUUIDV4String(t *testing.B) {
	uuidV4 := uuid.NewV4()

	for range t.N {
		_ = uuidV4.String()
	}
}

func BenchmarkUUIDV4String2(t *testing.B) {
	for range t.N {
		_ = uuid.NewV4().String()
	}
}

func BenchmarkUUIDV4Version(t *testing.B) {
	uuidV4 := uuid.NewV4()

	for range t.N {
		_ = uuidV4.Version()
	}
}
