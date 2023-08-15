package repl

import (
	"fmt"
	"log"
	"main/value"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type BaseScope struct {
	name      string
	parent    *BaseScope
	children  []*BaseScope
	variables map[string]*Variable
	functions map[string]value.IVOR
}

func (s *BaseScope) Name() string {
	return s.name
}

func (s *BaseScope) Parent() *BaseScope {
	return s.parent
}

func (s *BaseScope) Children() []*BaseScope {
	return s.children
}

func (s *BaseScope) ValidType(_type string) bool {
	// TODO: implement struct type validation
	return value.IsPrimitiveType(_type)
}

func (s *BaseScope) AddChild(child *BaseScope) {
	s.children = append(s.children, child)
	child.parent = s
}

func (s *BaseScope) variableExists(variable *Variable) bool {

	if _, ok := s.variables[variable.Name]; ok {
		return true
	}

	return false

}

func (s *BaseScope) AddVariable(name string, varType string, value value.IVOR, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {

	variable := &Variable{
		Name:     name,
		Value:    value,
		Type:     varType,
		IsConst:  isConst,
		AllowNil: allowNil,
		Token:    token,
	}

	if s.variableExists(variable) {
		return nil, "La variable " + name + " ya existe"
	}

	typesOk, msg := variable.TypeValidation()

	// even if the variable is not valid, we add it to the scope, (internally it will be nil)
	s.variables[name] = variable

	if !typesOk {
		// report error
		return nil, msg
	}

	return variable, ""
}

func (s *BaseScope) AddVector(name string, varType string, value value.IVOR, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {

	variable := &Variable{
		Name:     name,
		Value:    value,
		Type:     varType,
		IsConst:  isConst,
		AllowNil: allowNil,
		Token:    token,
	}

	if s.variableExists(variable) {
		return nil, "La variable " + name + " ya existe"
	}

	s.variables[name] = variable

	return variable, ""
}

func (s *BaseScope) GetVariable(name string) *Variable {
	// Todo: suport for structs properties

	initialScope := s

	for {
		if variable, ok := initialScope.variables[name]; ok {

			// verify if is refering to a pointer
			if variable.Type == value.IVOR_POINTER {
				return variable.Value.(*PointerValue).AssocVariable // pointer of a pointer ?
			}

			return variable
		}

		if initialScope.parent == nil {
			break
		}

		initialScope = initialScope.parent
	}

	return nil
}

func (s *BaseScope) AddFunction(name string, function value.IVOR) {
	s.functions[name] = function
}

func (s *BaseScope) GetFunction(name string) (value.IVOR, string) {
	// Todo: suport for structs properties

	// verify if is refering to and object/struct function
	if strings.Contains(name, ".") {
		return s.searchObjectFunction(name, nil)
	}

	initialScope := s

	for {
		if function, ok := initialScope.functions[name]; ok {
			return function, ""
		}

		if initialScope.parent == nil {
			break
		}

		initialScope = initialScope.parent
	}

	return nil, "La funcion " + name + " no existe"
}

// obj1.obj2.func1()

func (s *BaseScope) searchObjectFunction(name string, lastObj value.IVOR) (value.IVOR, string) {

	// split name by dot
	parts := strings.Split(name, ".")

	if len(parts) == 0 {
		log.Fatal("idk what u did, cant split by dot")
		return nil, ""
	}

	if len(parts) == 1 {
		obj, ok := lastObj.(*ObjectValue)

		if ok {
			return obj.InternalScope.GetFunction(name)
		}

		log.Fatal("idk what u did, cant convert to object")
		return nil, ""
	}

	// then parts should be 2 or more

	if lastObj == nil {
		variable := s.GetVariable(parts[0])

		if variable == nil {
			return nil, "No se puede acceder a la propiedad " + parts[0]
		}

		obj := variable.Value

		// obj must be an object/struct or vector

		switch obj := obj.(type) {
		case *ObjectValue:
			lastObj = obj
		case *VectorValue:
			lastObj = obj.ObjectValue
		default:
			return nil, "La propiedad '" + variable.Name + "' de tipo " + obj.Type() + " no tiene propiedades"
		}

		return s.searchObjectFunction(strings.Join(parts[1:], "."), lastObj)
	}

	obj, ok := lastObj.(*ObjectValue)

	if ok {
		lastObj = obj.InternalScope.GetVariable(parts[0]).Value

		return s.searchObjectFunction(strings.Join(parts[1:], "."), lastObj)
	} else {
		log.Fatal("idk what u did, cant convert to object")
		return nil, ""
	}
}

func (s *BaseScope) Reset() {
	s.variables = make(map[string]*Variable)
	s.children = make([]*BaseScope, 0)
	s.functions = make(map[string]value.IVOR)
}

func NewGlobalScope() *BaseScope {

	// register built-in functions

	funcs := make(map[string]value.IVOR)

	for k, v := range DefaultBuiltInFunctions {
		funcs[k] = v
	}

	return &BaseScope{
		name:      "global",
		variables: make(map[string]*Variable),
		children:  make([]*BaseScope, 0),
		parent:    nil,
		functions: funcs,
	}
}

func NewLocalScope(name string) *BaseScope {
	return &BaseScope{
		name:      name,
		variables: make(map[string]*Variable),
		functions: make(map[string]value.IVOR),
		children:  make([]*BaseScope, 0),
		parent:    nil,
	}
}

type ScopeTrace struct {
	GlobalScope  *BaseScope
	CurrentScope *BaseScope
}

func (s *ScopeTrace) PushScope(name string) *BaseScope {

	newScope := NewLocalScope(name)
	s.CurrentScope.AddChild(newScope)
	s.CurrentScope = newScope

	return s.CurrentScope
}

func (s *ScopeTrace) PopScope() {
	// ? implement new method, sending scope as pointer, and make CurrentScope equals the parent
	s.CurrentScope = s.CurrentScope.Parent()
}

func (s *ScopeTrace) Reset() {
	s.CurrentScope = s.GlobalScope
}

func (s *ScopeTrace) AddVariable(name string, varType string, value value.IVOR, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {
	return s.CurrentScope.AddVariable(name, varType, value, isConst, allowNil, token)
}

func (s *ScopeTrace) AddVector(name string, varType string, value value.IVOR, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {
	return s.CurrentScope.AddVector(name, varType, value, isConst, allowNil, token)
}

func (s *ScopeTrace) GetVariable(name string) *Variable {
	return s.CurrentScope.GetVariable(name)
}

func (s *ScopeTrace) AddFunction(name string, function value.IVOR) {
	s.CurrentScope.AddFunction(name, function)
}

func (s *ScopeTrace) GetFunction(name string) (value.IVOR, string) {
	return s.CurrentScope.GetFunction(name)
}

func (s *ScopeTrace) Print() {

	fmt.Println("Global Scope")
	fmt.Println("============")

	fmt.Println("Variables")
	for k, v := range s.GlobalScope.variables {
		fmt.Println(k, v.Value.Value())
	}

	fmt.Println("Funciones")
	for k, v := range s.GlobalScope.functions {
		fmt.Println(k, v)
	}

	fmt.Println("Child Scopes")
	fmt.Println("============")
	fmt.Println("")

	for _, child := range s.GlobalScope.children {

		fmt.Println(child.name)
		fmt.Println("============")

		fmt.Println("Variables")
		for k, v := range child.variables {
			fmt.Println(k, v.Value.Value())
		}

		fmt.Println("Funciones")
		for k, v := range child.functions {
			fmt.Println(k, v)
		}

		fmt.Println("")
	}

}

func NewScopeTrace() *ScopeTrace {
	globalScope := NewGlobalScope()
	return &ScopeTrace{
		GlobalScope:  globalScope,
		CurrentScope: globalScope,
	}
}

func NewVectorScope() *BaseScope {
	var scope = &BaseScope{
		name:      "vector",
		variables: make(map[string]*Variable),
		children:  make([]*BaseScope, 0),
		functions: make(map[string]value.IVOR),
		parent:    nil,
	}

	// register object built-in functions

	return scope
}
