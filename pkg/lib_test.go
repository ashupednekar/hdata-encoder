package pkg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	data := DataInput{
		Str("foo"),
		DataInput{
			Str("bar"),
			I32(42),
		},
	}
	serde := HDataSerde{}
	s := serde.Encode(data)
	fmt.Printf("encoded: %s", s)
	decoded := serde.Decode(s)
	assert.Equal(t, data, decoded, "decoded data should match original data")
}
