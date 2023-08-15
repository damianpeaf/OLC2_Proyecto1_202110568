package repl

import "main/value"

// implements ivor interface

type VectorValue struct {
	*ObjectValue
	InternalValue []value.IVOR
	CurrentIndex  int
	ItemType      string
}

func (v VectorValue) Value() interface{} {
	return v
}

func (v VectorValue) Type() string {
	return value.IVOR_VECTOR
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

func (v *VectorValue) Current() value.IVOR {
	return v.InternalValue[v.CurrentIndex]
}

func (v *VectorValue) Reset() {
	v.CurrentIndex = 0
}

var DefaultVectorInternalScope = &BaseScope{
	name: "vector",
}

func NewVectorValue(vectorItems []value.IVOR, itemType string) *VectorValue {
	vector := &VectorValue{
		InternalValue: vectorItems,
		CurrentIndex:  0,
		ItemType:      itemType,
	}

	AddVectorBuiltins(vector)

	return vector
}

var DefaultEmptyVectorValue = &VectorValue{
	InternalValue: []value.IVOR{},
	CurrentIndex:  0,
	ItemType:      value.IVOR_NIL,
}
