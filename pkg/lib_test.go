package pkg

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestEncodeDecode(t *testing.T) {
	data := RandomData(100, 20)
	serde := HDataSerde{}

	encoded, err := serde.Encode(data)
	assert.NoError(t, err)

	decoded, err := serde.Decode(encoded)
	assert.NoError(t, err)

	assert.Equal(t, data, decoded)
}
