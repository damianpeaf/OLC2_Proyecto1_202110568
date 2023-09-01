package repl

import (
	"main/compiler"
)

type Struct struct {
	Name   string
	Fields []compiler.IStruct_propContext
}
