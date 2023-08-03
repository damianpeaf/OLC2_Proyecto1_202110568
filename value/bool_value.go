package value

type BoolValue struct {
	InternalValue bool
}

func (b BoolValue) Value() interface{} {
	return b.InternalValue
}

func (b BoolValue) Type() string {
	return IVOR_BOOL
}
