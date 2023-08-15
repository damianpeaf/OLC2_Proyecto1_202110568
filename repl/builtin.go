package repl

import (
	"main/value"
	"strconv"
)

type BuiltInFunction struct {
	Name string
	Exec func(context *ReplContext, args []*Argument) (value.IVOR, bool, string)
}

// implementing ivor

func (b BuiltInFunction) Type() string {
	return value.IVOR_BUILTIN_FUNCTION
}

func (b BuiltInFunction) Value() interface{} {
	return b
}

func (b BuiltInFunction) Copy() value.IVOR {
	return b
}

// * Print Function
func Print(context *ReplContext, args []*Argument) (value.IVOR, bool, string) {

	var output string

	for i, arg := range args {

		if !value.IsPrimitiveType(arg.Object.Type()) {
			return value.DefaultNilValue, false, "La función print solo acepta tipos primitivos"
		}

		switch arg.Object.Type() {

		case value.IVOR_BOOL:
			output += strconv.FormatBool(arg.Object.Value().(bool))
		case value.IVOR_INT:
			output += strconv.Itoa(arg.Object.Value().(int))
		case value.IVOR_FLOAT:
			output += strconv.FormatFloat(arg.Object.Value().(float64), 'f', 4, 64) // 4 digits of precision
		case value.IVOR_STRING:
			output += arg.Object.Value().(string)
		case value.IVOR_CHARACTER:
			output += arg.Object.Value().(string)
		case value.IVOR_NIL:
			output += "nil"
		}

		// Add a space between each argument
		if i < len(args)-1 {
			output += " "
		}
	}

	context.Console.Print(output)

	return value.DefaultNilValue, true, ""
}

var DefaultBuiltInFunctions = map[string]*BuiltInFunction{
	"print": {
		Name: "print",
		Exec: Print,
	},
}
