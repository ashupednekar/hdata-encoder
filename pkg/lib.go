package pkg

type Serde interface {
	Encode(toSend DataInput) string
	Decode(received string) DataInput
}

type HDataSerde struct{}

func (h *HDataSerde) Encode(toSend DataInput) string {
	return ""
}

func (h *HDataSerde) Decode(received string) DataInput {
	return DataInput{}
}
