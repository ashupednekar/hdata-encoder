package pkg

import (
	"fmt"
	"testing"

	//"github.com/ashupednekar/hdata-encoder/internal/spec"
	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	// data := DataInput{
	// 	spec.Str("foo"),
	// 	DataInput{
	// 		spec.Str("bar"),
	// 		spec.I32(42),
	// 	},
	// }
	data := randomData(50000, 5000)
	fmt.Printf("data: %v\n", data)
	serde := HDataSerde{}
	s, err := serde.Encode(data)
	if err != nil {
		t.Errorf("error encoding data: %s", err)
	}
	//fmt.Printf("encoded: %v", s)
  sizeMB := float64(len(s)) / (1024 * 1024)
	fmt.Printf("encoded size: %.2f MB\n", sizeMB)
	decoded, err := serde.Decode(s)
	//fmt.Printf("decoded: %v", decoded)
	if err != nil {
		t.Errorf("error decoding data: %s", err)
	}
	assert.Equal(t, data, decoded, "decoded data should match original data")
}
