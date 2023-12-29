package test

import (
	"testing"

	"github.com/samborkent/uuid"
)

func BenchmarkUUIDV8(t *testing.B) {
	uuid.SetVersion(8)

	for i := 0; i < t.N; i++ {
		_ = uuid.New()
	}
}

func BenchmarkUUIDV8CreationTime(t *testing.B) {
	uuid.SetVersion(8)
	uuidV8 := uuid.New()

	for i := 0; i < t.N; i++ {
		_, _ = uuidV8.CreationTime()
	}
}

func BenchmarkUUIDV8Short(t *testing.B) {
	uuid.SetVersion(8)
	uuidV8 := uuid.New()

	for i := 0; i < t.N; i++ {
		_ = uuidV8.Short()
	}
}

func BenchmarkUUIDV8String(t *testing.B) {
	uuid.SetVersion(8)
	uuidV8 := uuid.New()

	for i := 0; i < t.N; i++ {
		_ = uuidV8.String()
	}
}

func BenchmarkUUIDV8Version(t *testing.B) {
	uuid.SetVersion(8)
	uuidV8 := uuid.New()

	for i := 0; i < t.N; i++ {
		_ = uuidV8.Version()
	}
}
