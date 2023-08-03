package value

type NilValue struct {
}

func (s NilValue) Value() interface{} {
	return nil
}

func (s NilValue) Type() string {
	return IVOR_NIL
}

var DefaultNilValue = &NilValue{}
