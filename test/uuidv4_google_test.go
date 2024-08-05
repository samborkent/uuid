package test

import (
	"testing"

	uuidGoogle "github.com/google/uuid"
)

func BenchmarkUUIDV4Google(t *testing.B) {
	for range t.N {
		_ = uuidGoogle.New()
	}
}

func BenchmarkUUIDV4GoogleString(t *testing.B) {
	uuid := uuidGoogle.New()

	for range t.N {
		_ = uuid.String()
	}
}
