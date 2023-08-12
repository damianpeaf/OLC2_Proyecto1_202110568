package repl

import (
	"main/value"

	"github.com/antlr4-go/antlr/v4"
)

type Variable struct {
	Name     string
	Value    value.IVOR
	Type     string
	AuxType  string
	IsConst  bool
	AllowNil bool
	Token    antlr.Token
}

func (v *Variable) TypeValidation() (bool, string) {

	if v.Value.Type() == value.IVOR_NIL && v.AllowNil {
		return true, ""
	}

	if v.Type != v.Value.Type() {

		// try implicit conversion
		convertedValue, ok := value.ImplicitCast(v.Type, v.Value)

		if !ok {
			// Si la expresión tiene un tipo de dato diferente al definido previamente se tomará como error y la variable obtendrá el valor de nil para fines prácticos.
			msg := "No se puede asignar un valor de tipo " + v.Value.Type() + " a una variable de tipo " + v.Type
			v.Value = value.DefaultNilValue
			return false, msg
		}

		v.Value = convertedValue
	}

	return true, ""
}

func (v *Variable) Assign(val value.IVOR) (bool, string) {

	if v.IsConst {
		return false, "No se puede asignar un valor a una constante"
	}

	v.Value = val

	return v.TypeValidation()
}
