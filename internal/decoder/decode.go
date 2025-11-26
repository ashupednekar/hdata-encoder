package decoder

import (
	"fmt"

	"github.com/ashupednekar/hdata-encoder/internal/spec"
)

func Decode(data []byte) (spec.DataInput, error) {
	idx := 0
	return decodeItems(data, &idx)
}

func decodeItems(buf []byte, idx *int) (spec.DataInput, error){
	if buf[*idx] != byte(spec.ItemTag){
		return nil, fmt.Errorf("expected itemTag, got %x", buf[*idx])
	}
	*idx++
	itemCount := readI32(buf, idx)
	res := make(spec.DataInput, 0, itemCount)
	for range(itemCount){
		val, err := decodeValue(buf, idx)
		if err != nil{
			return nil, err
		}
		res = append(res, val)
	}
	return res, nil
}

func decodeValue(buf []byte, idx *int) (spec.Value, error){
	tag := buf[*idx]
	*idx++

	switch spec.Tag(tag) {
	case spec.StrTag:
		length := readI32(buf, idx)
		s := string(buf[*idx: *idx+int(length)])
		*idx += int(length)
		return spec.Str(s), nil
	case spec.I32Tag:
		return spec.I32(readI32(buf, idx)), nil
	case spec.ItemTag:
		*idx--
		return decodeItems(buf, idx)
	default:
		return nil, fmt.Errorf("error decoding unkown tag: %x", tag)
	}

}

//TODO: implement DecodeConcurrent
