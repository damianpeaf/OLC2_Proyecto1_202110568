package repl

import (
	"fmt"
	"main/value"

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

func (s *BaseScope) AddVector(name string, varType string, auxType string, value value.IVOR, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {

	variable := &Variable{
		Name:     name,
		Value:    value,
		Type:     varType,
		IsConst:  isConst,
		AuxType:  auxType,
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

func (s *BaseScope) GetFunction(name string) value.IVOR {
	// Todo: suport for structs properties

	initialScope := s

	for {
		if function, ok := initialScope.functions[name]; ok {
			return function
		}

		if initialScope.parent == nil {
			break
		}

		initialScope = initialScope.parent
	}

	return nil
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

func (s *ScopeTrace) AddVector(name string, varType string, auxType string, value value.IVOR, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {
	return s.CurrentScope.AddVector(name, varType, auxType, value, isConst, allowNil, token)
}

func (s *ScopeTrace) GetVariable(name string) *Variable {
	return s.CurrentScope.GetVariable(name)
}

func (s *ScopeTrace) AddFunction(name string, function value.IVOR) {
	s.CurrentScope.AddFunction(name, function)
}

func (s *ScopeTrace) GetFunction(name string) value.IVOR {
	return s.CurrentScope.GetFunction(name)
}

func (s *ScopeTrace) Print() {

	fmt.Println("Global Scope")
	fmt.Println("============")

	for k, v := range s.GlobalScope.variables {
		fmt.Println(k, v.Value.Value())
	}

	fmt.Println("Child Scopes")
	fmt.Println("============")
	fmt.Println("")

	for _, child := range s.GlobalScope.children {

		fmt.Println(child.name)
		fmt.Println("============")

		for k, v := range child.variables {
			fmt.Println(k, v.Value.Value())
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
