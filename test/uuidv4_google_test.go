package test

import (
	"testing"

	uuidGoogle "github.com/google/uuid"
)

func BenchmarkUUIDV4Google(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = uuidGoogle.New()
	}
}

func BenchmarkUUIDV4GoogleString(t *testing.B) {
	uuid := uuidGoogle.New()

	for i := 0; i < t.N; i++ {
		_ = uuid.String()
	}
}
