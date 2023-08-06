package repl

import (
	"main/value"
)

type Argument struct {
	Name            string
	Object          value.IVOR
	PassByReference bool
}
