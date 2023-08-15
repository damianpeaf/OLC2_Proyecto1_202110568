package repl

import "main/value"

type ObjectValue struct {
	InternalScope *BaseScope
	AuxObject     interface{}
}

func (o ObjectValue) Value() interface{} {
	return o
}

func (o ObjectValue) Type() string {
	return value.IVOR_OBJECT
}
