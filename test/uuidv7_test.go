package test

import (
	"testing"

	"github.com/samborkent/uuid"
)

func BenchmarkUUIDV7(t *testing.B) {
	uuid.SetVersion(7)

	for i := 0; i < t.N; i++ {
		_ = uuid.New()
	}
}

func BenchmarkUUIDV7CreationTime(t *testing.B) {
	uuid.SetVersion(7)
	uuidV7 := uuid.New()

	for i := 0; i < t.N; i++ {
		_, _ = uuidV7.CreationTime()
	}
}

func BenchmarkUUIDV7Short(t *testing.B) {
	uuid.SetVersion(7)
	uuidV7 := uuid.New()

	for i := 0; i < t.N; i++ {
		_ = uuidV7.Short()
	}
}

func BenchmarkUUIDV7String(t *testing.B) {
	uuid.SetVersion(7)
	uuidV7 := uuid.New()

	for i := 0; i < t.N; i++ {
		_ = uuidV7.String()
	}
}

func BenchmarkUUIDV7Version(t *testing.B) {
	uuid.SetVersion(7)
	uuidV7 := uuid.New()

	for i := 0; i < t.N; i++ {
		_ = uuidV7.Version()
	}
}
