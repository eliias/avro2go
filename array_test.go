package avro2go

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArraySerialization(t *testing.T) {
	tests := []struct {
		a *Array
		j string
	}{
		{
			&Array{
				Type:  "array",
				Items: longType,
			},
			`{"type":"array","items":"long"}`,
		},
		{
			&Array{
				Type:  "array",
				Items: stringType,
			},
			`{"type":"array","items":"string"}`,
		},
	}

	for i, test := range tests {
		test := test
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			j, _ := json.Marshal(test.a)
			assert.Equal(t, test.j, string(j))
		})
	}
}
