package test

import (
	"testing"

	"github.com/samborkent/uuid"
)

func TestIsValidV7(t *testing.T) {
	testUUID := "01915146-80e1-741e-a983-6cbdd7c69015"

	if err := uuid.IsValidString(testUUID); err != nil {
		t.Errorf("uuid is invalid: %s", err.Error())
	}
}

func BenchmarkIsValidV4(t *testing.B) {
	uuidV4 := uuid.NewV4()

	for range t.N {
		_ = uuid.IsValid(uuidV4[:])
	}
}

func BenchmarkIsValidV4String(t *testing.B) {
	uuidV4 := uuid.NewV4().String()

	for range t.N {
		_ = uuid.IsValidString(uuidV4)
	}
}

func BenchmarkIsValidV7(t *testing.B) {
	uuidV7 := uuid.NewV7()

	for range t.N {
		_ = uuid.IsValid(uuidV7[:])
	}
}

func BenchmarkIsValidV7String(t *testing.B) {
	uuidV7 := uuid.NewV7().String()

	for range t.N {
		_ = uuid.IsValidString(uuidV7)
	}
}

func BenchmarkIsValidV8(t *testing.B) {
	uuidV8 := uuid.NewV8()

	for range t.N {
		_ = uuid.IsValid(uuidV8[:])
	}
}

func BenchmarkIsValidV8String(t *testing.B) {
	uuidV8 := uuid.NewV8().String()

	for range t.N {
		_ = uuid.IsValidString(uuidV8)
	}
}
