package test

import (
	"testing"

	uuidGoogle "github.com/google/uuid"
	"github.com/samborkent/uuid"
)

func BenchmarkToUUID(t *testing.B) {
	uuidV4 := uuidGoogle.New()

	for range t.N {
		_, _ = uuid.Parse(uuidV4[:])
	}
}
