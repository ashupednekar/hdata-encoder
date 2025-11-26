package pkg

type Value interface {
	isValue()
}

type Str string
type I32 int32
type DataInput []Value

func (_ Str) isValue()       {}
func (_ I32) isValue()       {}
func (_ DataInput) isValue() {}
