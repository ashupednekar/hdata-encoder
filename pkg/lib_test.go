package pkg

import (
	"fmt"
	"testing"

	"github.com/ashupednekar/hdata-encoder/internal/spec"
	//"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	data := DataInput{
		spec.Str("foo"),
		DataInput{
			spec.Str("bar"),
			spec.I32(42),
		},
	}
	serde := HDataSerde{}
	s, err := serde.Encode(data)
	if err != nil {
		t.Errorf("error encoding data: %s", err)
	}
	fmt.Printf("encoded: %v", s)
	//decoded := serde.Decode(s)
	//assert.Equal(t, data, decoded, "decoded data should match original data")
}
