package repl

import (
	"main/value"

	"github.com/antlr4-go/antlr/v4"
)

type Variable struct {
	Name     string
	Value    value.IVOR
	Type     string
	IsConst  bool
	AllowNil bool
	Token    antlr.Token
}

func (v *Variable) TypeValidation() (bool, string) {

	if v.Value.Type() == value.IVOR_NIL && v.AllowNil {
		return true, ""
	}

	if v.Type != v.Value.Type() {

		// vector type validation
		if IsVectorType(v.Type) && IsVectorType(v.Value.Type()) {

			// empty vector cast
			if v.Value.Type() == "[]" {

				// modify vector type
				v.Value.(*VectorValue).ItemType = RemoveBrackets(v.Type)
				v.Value.(*VectorValue).FullType = v.Type

				return true, ""
			}

			// implicit vector conversion

			targetType := RemoveBrackets(v.Type) // inner type
			newConvertedItems := make([]value.IVOR, 0)

			for _, item := range v.Value.(*VectorValue).InternalValue {
				convertedValue, ok := value.ImplicitCast(targetType, item)

				if !ok {
					break
				}
				newConvertedItems = append(newConvertedItems, convertedValue)
			}

			if len(newConvertedItems) == len(v.Value.(*VectorValue).InternalValue) {
				return true, ""
			}

			msg := "No se puede asignar un vector de tipo " + v.Value.Type() + " a una vector de tipo " + v.Type
			v.Value = value.DefaultNilValue
			return false, msg
		}

		// try implicit primitive conversion
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
