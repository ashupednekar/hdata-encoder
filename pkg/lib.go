package pkg

import (
	"encoding/base64"

	"github.com/ashupednekar/hdata-encoder/internal/encoder"
	"github.com/ashupednekar/hdata-encoder/internal/spec"
)

type DataInput = spec.DataInput

type Serde interface {
	Encode(toSend DataInput) ([]byte, error)
	Decode(received []byte) DataInput
  EncodeB64(toSend DataInput) (string, error)
	DecodeB64(received string) DataInput

}

type HDataSerde struct{}

func (h *HDataSerde) Encode(toSend DataInput) ([]byte, error) {
	return encoder.Encode(&toSend)
}

func (h *HDataSerde) EncodeB64(toSend DataInput) (string, error){
	buf, err := h.Encode(toSend)
  if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf), nil

}
 
func (h *HDataSerde) Decode(received []byte) DataInput {
	return DataInput{}
}

func (h *HDataSerde) DecodeB64(received string) DataInput {
	return DataInput{}
}
