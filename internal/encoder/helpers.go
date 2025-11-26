package encoder

import "github.com/ashupednekar/hdata-encoder/internal/spec"

func appendI32(buf []byte, n spec.I32) []byte {
	return append(
		buf,
		byte(n>>24), byte(n>>16), byte(n>>8), byte(n),
	)
}
