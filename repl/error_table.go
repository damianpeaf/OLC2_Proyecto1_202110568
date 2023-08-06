package repl

const (
	LexicalError  = "Error léxico"
	SyntaxError   = "Error sintáctico"
	SemanticError = "Error semántico"
)

type Error struct {
	Line   int
	Column int
	Msg    string
	Type   int
}

type ErrorTable struct {
	Errors []Error
}

func (et *ErrorTable) AddError(line int, column int, msg string, errorType int) {
	et.Errors = append(et.Errors, Error{line, column, msg, errorType})
}

func NewErrorTable() *ErrorTable {
	return &ErrorTable{}
}
