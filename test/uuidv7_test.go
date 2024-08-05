package test

import (
	"testing"

	"github.com/samborkent/uuid"
)

func TestUUIDV7(t *testing.T) {
	t.Log(uuid.NewV7())
}

func BenchmarkUUIDV7(t *testing.B) {
	for range t.N {
		_ = uuid.NewV7()
	}
}

func BenchmarkUUIDV7CreationTime(t *testing.B) {
	uuidV7 := uuid.NewV7()

	for range t.N {
		_, _ = uuidV7.CreationTime()
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
