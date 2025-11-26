package decoder

import "github.com/ashupednekar/hdata-encoder/internal/spec"

func readI32(buf []byte, idx *int) spec.I32 {
	x := int32(buf[*idx])<<24 |
		int32(buf[*idx+1])<<16 |
		int32(buf[*idx+2])<<8 |
		int32(buf[*idx+3])
	*idx += 4
	return spec.I32(x)
}
