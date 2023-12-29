package test

import (
	"testing"

	"github.com/samborkent/uuid"
)

func BenchmarkUUIDV4(t *testing.B) {
	uuid.SetVersion(4)

	for i := 0; i < t.N; i++ {
		_ = uuid.New()
	}
}

func BenchmarkUUIDV4Short(t *testing.B) {
	uuid.SetVersion(4)
	uuidV4 := uuid.New()

	for i := 0; i < t.N; i++ {
		_ = uuidV4.Short()
	}
}

func BenchmarkUUIDV4String(t *testing.B) {
	uuid.SetVersion(4)
	uuidV4 := uuid.New()

	for i := 0; i < t.N; i++ {
		_ = uuidV4.String()
	}
}

func BenchmarkUUIDV4Version(t *testing.B) {
	uuid.SetVersion(4)
	uuidV4 := uuid.New()

	for i := 0; i < t.N; i++ {
		_ = uuidV4.Version()
	}
}
