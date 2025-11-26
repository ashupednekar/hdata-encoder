package encoder

import (
	"fmt"

	"github.com/ashupednekar/hdata-encoder/internal/spec"
)

func Encode(data *spec.DataInput) ([]byte, error) {
	buf := make([]byte, 0, 5)
	buf = append(buf, byte(spec.ItemTag))
	buf = appendI32(buf, spec.I32(len(*data)))
	for _, value := range *data {
		switch v := value.(type) {
		case spec.Str:
			s := string(v)
			strBuf := make([]byte, 0, 1+4*len(s))
			strBuf = append(strBuf, byte(spec.StrTag))
			strBuf = appendI32(strBuf, spec.I32(len(s)))
			strBuf = append(strBuf, s...)
			buf = append(buf, strBuf...)
		case spec.I32:
			i32Buf := make([]byte, 0, 1+4)
			i32Buf = append(i32Buf, byte(spec.I32Tag))
			i32Buf = appendI32(i32Buf, v)
			buf = append(buf, i32Buf...)
		case spec.DataInput:
			child, err := Encode(&v)
			if err != nil {
				return nil, err
			}
			buf = append(buf, child...)
		default:
			return nil, fmt.Errorf("error encoding unknown value: %T", v)
		}
	}
	return buf, nil
}


//TODO: implement EncodeWithOffset for concurrent decoding
