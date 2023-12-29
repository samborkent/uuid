package test

import (
	"testing"

	"github.com/samborkent/uuid"
)

func BenchmarkIsValidV4(t *testing.B) {
	uuid.SetVersion(4)
	uuidV4 := uuid.New()

	for i := 0; i < t.N; i++ {
		_ = uuid.IsValid(uuidV4[:])
	}
}

func BenchmarkIsValidV4String(t *testing.B) {
	uuid.SetVersion(4)
	uuidV4 := uuid.New().String()

	for i := 0; i < t.N; i++ {
		_ = uuid.IsValidString(uuidV4)
	}
}

func BenchmarkIsValidV7(t *testing.B) {
	uuid.SetVersion(7)
	uuidV7 := uuid.New()

	for i := 0; i < t.N; i++ {
		_ = uuid.IsValid(uuidV7[:])
	}
}

func BenchmarkIsValidV7String(t *testing.B) {
	uuid.SetVersion(7)
	uuidV7 := uuid.New().String()

	for i := 0; i < t.N; i++ {
		_ = uuid.IsValidString(uuidV7)
	}
}

func BenchmarkIsValidV8(t *testing.B) {
	uuid.SetVersion(8)
	uuidV8 := uuid.New()

	for i := 0; i < t.N; i++ {
		_ = uuid.IsValid(uuidV8[:])
	}
}

func BenchmarkIsValidV8String(t *testing.B) {
	uuid.SetVersion(8)
	uuidV8 := uuid.New().String()

	for i := 0; i < t.N; i++ {
		_ = uuid.IsValidString(uuidV8)
	}
}
