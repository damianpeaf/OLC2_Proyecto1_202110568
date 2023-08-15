package repl

import "main/value"

func StringToVector(s *value.StringValue) *VectorValue {

	items := make([]value.IVOR, 0)

	for _, c := range s.InternalValue {
		items = append(items, &value.CharacterValue{InternalValue: string(c)})
	}

	return NewVectorValue(items, value.IVOR_CHARACTER)

}
