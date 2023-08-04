package repl

import "main/value"

type Variable struct {
	Name    string
	Value   value.IVOR // TODO: change to a pointer
	Type    string
	IsConst bool
}

func (v *Variable) Validate() (bool, string) {

	if v.Type != v.Value.Type() {
		return false, "Type mismatch"
	}

	return true, ""
}
