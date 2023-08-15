package repl

import (
	"main/value"

	"github.com/antlr4-go/antlr/v4"
)

type ObjectBuiltInFunction struct {
	*Function
	Object     *ObjectValue
	CustomExec func(builtinRef *ObjectBuiltInFunction, visitor *ReplVisitor, args map[string]*Argument, token antlr.Token)
}

// implementing ivor
func (b ObjectBuiltInFunction) Type() string {
	return value.IVOR_FUNCTION
}

func (b ObjectBuiltInFunction) Value() interface{} {
	return b
}

func (b ObjectBuiltInFunction) Copy() value.IVOR {
	return b
}

func (f *ObjectBuiltInFunction) Exec(visitor *ReplVisitor, args []*Argument, token antlr.Token) {

	context := visitor.GetReplContext()

	// validate args
	argsOk, argsMap := f.ValidateArgs(context, args, token)

	if !argsOk {
		f.ReturnValue = value.DefaultNilValue
		return
	}

	f.CustomExec(f, visitor, argsMap, token)

}

// * Vector Built In Functions

// 1. Append
// vector.append(value)

// Parameters:

var appendParams = []*Param{
	// Just one positional argument
	{
		ExternName:      "_",
		InnerName:       "_",
		Type:            value.IVOR_ANY,
		PassByReference: false,
		Token:           nil,
	},
}

func appendCustomExec(builtinRef *ObjectBuiltInFunction, visitor *ReplVisitor, args map[string]*Argument, token antlr.Token) {

	builtinRef.ReturnValue = value.DefaultNilValue

	// get the vector
	vector := builtinRef.Object.AuxObject.(*VectorValue)

	// get the value
	arg := args["_"]

	if vector.ItemType != arg.Object.Type() {
		visitor.ErrorTable.NewSemanticError(arg.Token, "No se puede agregar un valor de tipo "+arg.Object.Type()+" a un vector de tipo "+vector.ItemType)
		return
	}
	vector.InternalValue = append(vector.InternalValue, arg.Object)
}

// 2. vector.remove(at: Int) -> nil

// parameters:

var removeParams = []*Param{
	// at: Int
	{
		ExternName:      "at",
		InnerName:       "at",
		Type:            value.IVOR_INT,
		PassByReference: false,
		Token:           nil,
	},
}

func removeCustomExec(builtinRef *ObjectBuiltInFunction, visitor *ReplVisitor, args map[string]*Argument, token antlr.Token) {

	builtinRef.ReturnValue = value.DefaultNilValue

	// get the vector
	vector := builtinRef.Object.AuxObject.(*VectorValue)

	// get the value
	arg := args["at"]

	if arg.Object.Type() != value.IVOR_INT {
		visitor.ErrorTable.NewSemanticError(arg.Token, "El argumento 'at' debe ser de tipo Int")
		return
	}

	// out of bounds
	if arg.Object.Value().(int) >= vector.Size() || arg.Object.Value().(int) < 0 {
		visitor.ErrorTable.NewSemanticError(arg.Token, "El indice esta fuera de rango")
		return
	}

	// remove the element
	vector.InternalValue = append(vector.InternalValue[:arg.Object.Value().(int)], vector.InternalValue[arg.Object.Value().(int)+1:]...)
}

// 3. removeLast
// vector.removeLast() -> nil

// parameters:

var removeLastParams = []*Param{}

func removeLastCustomExec(builtinRef *ObjectBuiltInFunction, visitor *ReplVisitor, args map[string]*Argument, token antlr.Token) {

	builtinRef.ReturnValue = value.DefaultNilValue

	// get the vector
	vector := builtinRef.Object.AuxObject.(*VectorValue)

	if vector.Size() == 0 {
		visitor.ErrorTable.NewSemanticError(token, "El vector esta vacio y no se puede remover el ultimo elemento")
		return
	}

	// remove the last element
	vector.InternalValue = vector.InternalValue[:vector.Size()-1]
}

// 3. isEmpty
// vector.isEmpty() -> Bool

// parameters:

var isEmptyParams = []*Param{}

func isEmptyCustomExec(builtinRef *ObjectBuiltInFunction, visitor *ReplVisitor, args map[string]*Argument, token antlr.Token) {

	// get the vector
	vector := builtinRef.Object.AuxObject.(*VectorValue)

	if vector.Size() == 0 {
		builtinRef.ReturnValue = &value.BoolValue{
			InternalValue: true,
		}
		return
	}

	builtinRef.ReturnValue = &value.BoolValue{
		InternalValue: false,
	}
}

// 4. count
// vector.count() -> Int

// parameters:

var countParams = []*Param{}

func countCustomExec(builtinRef *ObjectBuiltInFunction, visitor *ReplVisitor, args map[string]*Argument, token antlr.Token) {

	// get the vector
	vector := builtinRef.Object.AuxObject.(*VectorValue)

	builtinRef.ReturnValue = &value.IntValue{
		InternalValue: vector.Size(),
	}
}

func AddVectorBuiltins(vectorRef *VectorValue) {

	vectorScope := NewVectorScope()

	vectorInternalObject := &ObjectValue{
		InternalScope: vectorScope,
		AuxObject:     vectorRef,
	}

	// Register built in functions
	vectorScope.AddFunction("append", &ObjectBuiltInFunction{
		Function: &Function{
			Param: appendParams,
		},
		Object:     vectorInternalObject,
		CustomExec: appendCustomExec,
	})

	vectorScope.AddFunction("remove", &ObjectBuiltInFunction{
		Function: &Function{
			Param: removeParams,
		},
		Object:     vectorInternalObject,
		CustomExec: removeCustomExec,
	})

	vectorScope.AddFunction("removeLast", &ObjectBuiltInFunction{
		Function: &Function{
			Param: removeLastParams,
		},
		Object:     vectorInternalObject,
		CustomExec: removeLastCustomExec,
	})

	vectorScope.AddFunction("isEmpty", &ObjectBuiltInFunction{
		Function: &Function{
			Param: isEmptyParams,
		},
		Object:     vectorInternalObject,
		CustomExec: isEmptyCustomExec,
	})

	vectorScope.AddFunction("count", &ObjectBuiltInFunction{
		Function: &Function{
			Param: countParams,
		},
		Object:     vectorInternalObject,
		CustomExec: countCustomExec,
	})

	vectorRef.ObjectValue = vectorInternalObject

}
