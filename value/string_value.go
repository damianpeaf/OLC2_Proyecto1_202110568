package value

type StringValue struct {
	InternalValue string
}

func (s StringValue) Value() interface{} {
	return s.InternalValue
}

func (s StringValue) Type() string {
	return IVOR_STRING
}

func (s StringValue) ToVector() *VectorValue {

	items := make([]IVOR, 0)

	for _, c := range s.InternalValue {
		items = append(items, &CharacterValue{InternalValue: string(c)})
	}

	return &VectorValue{InternalValue: items, CurrentIndex: 0, Builtins: nil, ItemType: IVOR_STRING}

}
