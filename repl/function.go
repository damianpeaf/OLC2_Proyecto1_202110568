package repl

import (
	"main/compiler"
	"main/value"

	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Function struct {
	Name        string
	Param       []*Param
	ReturnType  string
	Body        []compiler.IStmtContext
	DeclScope   *BaseScope
	ReturnValue value.IVOR
	// TODO: suport for structs, mutator, self, etc
}

func (f *Function) Value() interface{} {
	return f
}

func (f *Function) Type() string {
	return value.IVOR_FUNCTION
}

func (f *Function) Exec(visitor *ReplVisitor, args []*Argument, token antlr.Token) {

	context := visitor.GetReplContext()

	// validate args
	argsOk, argsMap := f.ValidateArgs(context, args, token)

	if !argsOk {
		f.ValidateReturn(context, value.DefaultNilValue, token)
		return
	}

	// create new scope
	initialScope := context.ScopeTrace.CurrentScope          // save current scope, scope at call time
	context.ScopeTrace.CurrentScope = f.DeclScope            // set function declaration scope as current scope
	context.ScopeTrace.PushScope("func: " + token.GetText()) // push function scope

	// push args to scope

	for varName, arg := range argsMap {
		context.ScopeTrace.CurrentScope.AddVariable(varName, arg.Object.Type(), arg.Object, false, false, arg.Token)
	}

	// push return item to callstack
	funcItem := &CallStackItem{
		ReturnValue: value.DefaultNilValue,
		Type: []string{
			ReturnItem,
		},
	}
	context.CallStack.Push(funcItem)

	// handle return from callstack

	defer func() {

		context.CallStack.Clean(funcItem)              // clean callstack
		context.ScopeTrace.PopScope()                  // pop function scope
		context.ScopeTrace.CurrentScope = initialScope // restore the call time scope

		if item, ok := recover().(*CallStackItem); item != nil && ok {

			if item != funcItem {
				context.ErrorTable.NewSemanticError(token, "Return invalido")
			}

			// validate return type
			f.ValidateReturn(context, item.ReturnValue, token) // return value from return statement
			return
		}

		f.ValidateReturn(context, value.DefaultNilValue, token)
	}()

	// evaluate body
	for _, stmt := range f.Body {
		visitor.Visit(stmt)
	}

	f.ValidateReturn(context, value.DefaultNilValue, token)
	// return
}

func (f *Function) ValidateArgs(context *ReplContext, args []*Argument, token antlr.Token) (bool, map[string]*Argument) {

	// validate arg count
	if len(args) != len(f.Param) {
		context.ErrorTable.NewSemanticError(token, "Numero de argumentos invalido")
		return false, nil
	}

	argsMap := make(map[string]*Argument)
	finalArgsMap := make(map[string]*Argument)

	for _, arg := range args {
		argsMap[arg.Name] = arg
	}

	for i, param := range f.Param {

		// determine param type
		var argToValidate *Argument = nil

		if param.ExternName == "" {
			// inner = arg name
			argToValidate = argsMap[param.InnerName]

		} else if param.ExternName == "_" {
			// positional arg
			argToValidate = args[i]
		} else {
			// extern = arg name
			argToValidate = argsMap[param.ExternName]
		}

		// validate arg exists
		if argToValidate == nil {
			context.ErrorTable.NewSemanticError(token, fmt.Sprintf("Argumento %s no especificado", param.InnerName))
			return false, nil
		}

		// validate type
		if argToValidate.Object.Type() != param.Type {
			context.ErrorTable.NewSemanticError(token, fmt.Sprintf("Tipo de argumento %s invalido", param.InnerName))
			return false, nil
		}

		// validate pass by reference
		if argToValidate.PassByReference != param.PassByReference {
			context.ErrorTable.NewSemanticError(token, fmt.Sprintf("Argumento %s no es por referencia", param.InnerName))
			return false, nil
		}

		// add to final args map
		finalArgsMap[param.InnerName] = argToValidate
	}

	return true, finalArgsMap
}

func (f *Function) ValidateReturn(context *ReplContext, val value.IVOR, token antlr.Token) {

	if val.Type() != f.ReturnType {
		context.ErrorTable.NewSemanticError(token, fmt.Sprintf("Tipo de retorno invalido, se esperaba %s, se obtuvo %s", f.ReturnType, val.Type()))

		f.ReturnValue = value.DefaultNilValue
	}

	f.ReturnValue = val
}
