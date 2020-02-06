package avro2go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion_String(t *testing.T) {
	u := &Union{Types{nullType, stringType}}
	assert.Equal(t, `["null","string"]`, u.String())
}
