package repl

import (
	"main/compiler"
	"main/value"

	"github.com/antlr4-go/antlr/v4"
)

type Struct struct {
	Name   string
	Fields []compiler.IStruct_propContext
}

type StructArg struct {
	Name  string
	Value value.IVOR
	Token antlr.Token
}
