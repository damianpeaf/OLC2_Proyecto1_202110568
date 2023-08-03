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
	// todo: inspect parent scopes

	variable, ok := s.variables[name]

	if ok {
		return variable
	}

	return nil
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

func (s *ScopeTrace) PushScope(name string) {

	newScope := NewLocalScope(name)
	s.CurrentScope.AddChild(newScope)
	s.CurrentScope = newScope

}

func (s *ScopeTrace) PopScope() {
	s.CurrentScope = s.CurrentScope.Parent()
}

func (s *ScopeTrace) Print() {

	fmt.Println("Global Scope")
	fmt.Println("============")

	for k, v := range s.GlobalScope.variables {
		fmt.Println(k, v)
	}

}

func NewScopeTrace() *ScopeTrace {
	globalScope := NewGlobalScope()
	return &ScopeTrace{
		GlobalScope:  globalScope,
		CurrentScope: globalScope,
	}
}
