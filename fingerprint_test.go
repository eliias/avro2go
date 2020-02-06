package avro2go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFingerprint64(t *testing.T) {
	assert.Equal(t, uint64(15532896163603310461), Fingerprint64([]byte("test")))
}
