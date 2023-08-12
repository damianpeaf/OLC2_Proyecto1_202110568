package value

// implements ivor interface

type VectorValue struct {
	InternalValue []IVOR
	CurrentIndex  int
	Builtins      map[string]IVOR
	ItemType      string
}

func (v VectorValue) Value() interface{} {
	return v
}

func (v VectorValue) Type() string {
	return IVOR_VECTOR
}

func (v VectorValue) Size() int {
	return len(v.InternalValue)
}

func (v *VectorValue) Next() bool {
	if v.CurrentIndex < len(v.InternalValue) {
		v.CurrentIndex++
		return true
	}
	return false
}

func (v *VectorValue) Current() IVOR {
	return v.InternalValue[v.CurrentIndex]
}

func (v *VectorValue) Reset() {
	v.CurrentIndex = 0
}

var DefaultEmptyVectorValue = &VectorValue{
	InternalValue: []IVOR{},
	CurrentIndex:  0,
	Builtins:      nil,
	ItemType:      IVOR_NIL,
}
