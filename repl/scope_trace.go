package repl

import (
	"fmt"
	"main/value"
)

type BaseScope struct {
	name      string
	parent    *BaseScope
	children  []*BaseScope
	variables map[string]*Variable
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

func (s *BaseScope) AddVariable(name string, varType string, value value.IVOR, isConst bool) *Variable {

	variable := &Variable{
		Name:    name,
		Value:   value,
		Type:    varType,
		IsConst: isConst,
	}

	// TODO: check if variable already exists
	s.variables[name] = variable

	return variable
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

func (s *BaseScope) Reset() {
	s.variables = make(map[string]*Variable)
	s.children = make([]*BaseScope, 0)
}

func NewGlobalScope() *BaseScope {
	return &BaseScope{
		name:      "global",
		variables: make(map[string]*Variable),
		children:  make([]*BaseScope, 0),
		parent:    nil,
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

func (s *ScopeTrace) AddVariable(name string, varType string, value value.IVOR, isConst bool) *Variable {
	return s.CurrentScope.AddVariable(name, varType, value, isConst)
}

func (s *ScopeTrace) GetVariable(name string) *Variable {
	return s.CurrentScope.GetVariable(name)
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
