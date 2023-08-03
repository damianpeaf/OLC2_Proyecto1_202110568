package value

type IntValue struct {
	InternalValue int
}

func (i IntValue) Value() interface{} {
	return i.InternalValue
}

func (i IntValue) Type() string {
	return IVOR_INT
}
