package test

import (
	"testing"

	uuidGoogle "github.com/google/uuid"
	"github.com/samborkent/uuid"
)

func BenchmarkToUUID(t *testing.B) {
	uuidV4 := uuidGoogle.New()

	for i := 0; i < t.N; i++ {
		_, _ = uuid.ToUUID(uuidV4[:])
	}
}
