package repl

import (
	"main/value"

	"github.com/antlr4-go/antlr/v4"
)

type Argument struct {
	Name            string
	Object          value.IVOR
	PassByReference bool
	Token           antlr.Token
}
