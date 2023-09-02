package errors

import (
	"main/repl"

	"github.com/antlr4-go/antlr/v4"
)

type SyntaxErrorListener struct {
	*antlr.DefaultErrorListener
	ErrorTable *repl.ErrorTable
}

func NewSyntaxErrorListener(errorTable *repl.ErrorTable) *SyntaxErrorListener {
	return &SyntaxErrorListener{
		ErrorTable: errorTable,
	}
}

func (e *SyntaxErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {

	e.ErrorTable.AddError(
		line,
		column,
		msg,
		repl.SyntaxError,
	)

}

type LexicalErrorListener struct {
	*antlr.DefaultErrorListener
	ErrorTable *repl.ErrorTable
}

func NewLexicalErrorListener() *LexicalErrorListener {
	return &LexicalErrorListener{
		ErrorTable: repl.NewErrorTable(),
	}
}

func (e *LexicalErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {

	e.ErrorTable.AddError(
		line,
		column,
		msg,
		repl.LexicalError,
	)

}
