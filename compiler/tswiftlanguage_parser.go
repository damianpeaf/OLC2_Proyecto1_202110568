// Code generated from compiler\TSwiftLanguage.g4 by ANTLR 4.13.0. DO NOT EDIT.

package compiler // TSwiftLanguage
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TSwiftLanguage struct {
	*antlr.BaseParser
}

var TSwiftLanguageParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tswiftlanguageParserInit() {
	staticData := &TSwiftLanguageParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "';'", "", "'let'", "'var'", "'func'", "'struct'", "'if'",
		"'else'", "'switch'", "'case'", "'default'", "'for'", "'while'", "'break'",
		"'continue'", "'return'", "'Int'", "'Float'", "'String'", "'Bool'",
		"'Character'", "", "", "", "", "'nil'", "", "'+'", "'-'", "'*'", "'/'",
		"'%'", "'='", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'&&'",
		"'||'", "'!'", "'('", "')'", "'{'", "'}'", "'['", "']'", "','", "'.'",
		"':'", "'->'", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "COMMENT", "MULTILINE_COMMENT", "SEMICOLON", "NEWLINE", "LET_KW",
		"VAR_KW", "FUNC_KW", "STRUCT_KW", "IF_KW", "ELSE_KW", "SWITCH_KW", "CASE_KW",
		"DEFAULT_KW", "FOR_KW", "WHILE_KW", "BREAK_KW", "CONTINUE_KW", "RETURN_KW",
		"INTEGER_TYPE", "FLOAT_TYPE", "STRING_TYPE", "BOOL_TYPE", "CHARACTER_TYPE",
		"INTEGER_LITERAL", "FLOAT_LITERAL", "STRING_LITERAL", "BOOL_LITERAL",
		"NIL_LITERAL", "ID", "PLUS", "MINUS", "MULT", "DIV", "MOD", "EQUALS",
		"EQUALS_EQUALS", "NOT_EQUALS", "LESS_THAN", "LESS_THAN_OR_EQUAL", "GREATER_THAN",
		"GREATER_THAN_OR_EQUAL", "AND", "OR", "NOT", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "LBRACK", "RBRACK", "COMMA", "DOT", "COLON", "ARROW", "INTERROGATION",
	}
	staticData.RuleNames = []string{
		"program", "delimiter", "stmt", "decl_stmt", "var_type", "primitive_type",
		"assign_stmt", "literal", "expr", "if_stmt",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 56, 116, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 5, 0, 24, 8, 0, 10, 0, 12, 0, 27, 9, 0, 1, 1, 1, 1, 4, 1, 31,
		8, 1, 11, 1, 12, 1, 32, 1, 1, 3, 1, 36, 8, 1, 1, 2, 1, 2, 3, 2, 40, 8,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 60, 8, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 75, 8,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 86, 8, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		5, 8, 100, 8, 8, 10, 8, 12, 8, 103, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9,
		109, 8, 9, 10, 9, 12, 9, 112, 9, 9, 1, 9, 1, 9, 1, 9, 0, 1, 16, 10, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 7, 1, 0, 6, 7, 1, 0, 20, 23, 2, 0, 32,
		32, 45, 45, 1, 0, 31, 32, 1, 0, 33, 34, 1, 0, 37, 38, 1, 0, 39, 42, 124,
		0, 25, 1, 0, 0, 0, 2, 35, 1, 0, 0, 0, 4, 39, 1, 0, 0, 0, 6, 59, 1, 0, 0,
		0, 8, 61, 1, 0, 0, 0, 10, 63, 1, 0, 0, 0, 12, 65, 1, 0, 0, 0, 14, 74, 1,
		0, 0, 0, 16, 85, 1, 0, 0, 0, 18, 104, 1, 0, 0, 0, 20, 21, 3, 4, 2, 0, 21,
		22, 3, 2, 1, 0, 22, 24, 1, 0, 0, 0, 23, 20, 1, 0, 0, 0, 24, 27, 1, 0, 0,
		0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 1, 1, 0, 0, 0, 27, 25, 1,
		0, 0, 0, 28, 36, 5, 4, 0, 0, 29, 31, 5, 5, 0, 0, 30, 29, 1, 0, 0, 0, 31,
		32, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 36, 1, 0, 0,
		0, 34, 36, 5, 0, 0, 1, 35, 28, 1, 0, 0, 0, 35, 30, 1, 0, 0, 0, 35, 34,
		1, 0, 0, 0, 36, 3, 1, 0, 0, 0, 37, 40, 3, 6, 3, 0, 38, 40, 3, 12, 6, 0,
		39, 37, 1, 0, 0, 0, 39, 38, 1, 0, 0, 0, 40, 5, 1, 0, 0, 0, 41, 42, 3, 8,
		4, 0, 42, 43, 5, 30, 0, 0, 43, 44, 5, 54, 0, 0, 44, 45, 3, 10, 5, 0, 45,
		46, 5, 36, 0, 0, 46, 47, 3, 16, 8, 0, 47, 60, 1, 0, 0, 0, 48, 49, 3, 8,
		4, 0, 49, 50, 5, 30, 0, 0, 50, 51, 5, 36, 0, 0, 51, 52, 3, 16, 8, 0, 52,
		60, 1, 0, 0, 0, 53, 54, 3, 8, 4, 0, 54, 55, 5, 30, 0, 0, 55, 56, 5, 54,
		0, 0, 56, 57, 3, 10, 5, 0, 57, 58, 5, 56, 0, 0, 58, 60, 1, 0, 0, 0, 59,
		41, 1, 0, 0, 0, 59, 48, 1, 0, 0, 0, 59, 53, 1, 0, 0, 0, 60, 7, 1, 0, 0,
		0, 61, 62, 7, 0, 0, 0, 62, 9, 1, 0, 0, 0, 63, 64, 7, 1, 0, 0, 64, 11, 1,
		0, 0, 0, 65, 66, 5, 30, 0, 0, 66, 67, 5, 36, 0, 0, 67, 68, 3, 16, 8, 0,
		68, 13, 1, 0, 0, 0, 69, 75, 5, 25, 0, 0, 70, 75, 5, 26, 0, 0, 71, 75, 5,
		27, 0, 0, 72, 75, 5, 28, 0, 0, 73, 75, 5, 29, 0, 0, 74, 69, 1, 0, 0, 0,
		74, 70, 1, 0, 0, 0, 74, 71, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 73, 1,
		0, 0, 0, 75, 15, 1, 0, 0, 0, 76, 77, 6, 8, -1, 0, 77, 78, 7, 2, 0, 0, 78,
		86, 3, 16, 8, 4, 79, 80, 5, 46, 0, 0, 80, 81, 3, 16, 8, 0, 81, 82, 5, 47,
		0, 0, 82, 86, 1, 0, 0, 0, 83, 86, 5, 30, 0, 0, 84, 86, 3, 14, 7, 0, 85,
		76, 1, 0, 0, 0, 85, 79, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 84, 1, 0, 0,
		0, 86, 101, 1, 0, 0, 0, 87, 88, 10, 8, 0, 0, 88, 89, 7, 3, 0, 0, 89, 100,
		3, 16, 8, 9, 90, 91, 10, 7, 0, 0, 91, 92, 7, 4, 0, 0, 92, 100, 3, 16, 8,
		8, 93, 94, 10, 6, 0, 0, 94, 95, 7, 5, 0, 0, 95, 100, 3, 16, 8, 7, 96, 97,
		10, 5, 0, 0, 97, 98, 7, 6, 0, 0, 98, 100, 3, 16, 8, 6, 99, 87, 1, 0, 0,
		0, 99, 90, 1, 0, 0, 0, 99, 93, 1, 0, 0, 0, 99, 96, 1, 0, 0, 0, 100, 103,
		1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 17, 1, 0, 0,
		0, 103, 101, 1, 0, 0, 0, 104, 105, 5, 10, 0, 0, 105, 106, 3, 16, 8, 0,
		106, 110, 5, 48, 0, 0, 107, 109, 3, 4, 2, 0, 108, 107, 1, 0, 0, 0, 109,
		112, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 113,
		1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 113, 114, 5, 49, 0, 0, 114, 19, 1, 0,
		0, 0, 10, 25, 32, 35, 39, 59, 74, 85, 99, 101, 110,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// TSwiftLanguageInit initializes any static state used to implement TSwiftLanguage. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTSwiftLanguage(). You can call this function if you wish to initialize the static state ahead
// of time.
func TSwiftLanguageInit() {
	staticData := &TSwiftLanguageParserStaticData
	staticData.once.Do(tswiftlanguageParserInit)
}

// NewTSwiftLanguage produces a new parser instance for the optional input antlr.TokenStream.
func NewTSwiftLanguage(input antlr.TokenStream) *TSwiftLanguage {
	TSwiftLanguageInit()
	this := new(TSwiftLanguage)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TSwiftLanguageParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "TSwiftLanguage.g4"

	return this
}

// TSwiftLanguage tokens.
const (
	TSwiftLanguageEOF                   = antlr.TokenEOF
	TSwiftLanguageWS                    = 1
	TSwiftLanguageCOMMENT               = 2
	TSwiftLanguageMULTILINE_COMMENT     = 3
	TSwiftLanguageSEMICOLON             = 4
	TSwiftLanguageNEWLINE               = 5
	TSwiftLanguageLET_KW                = 6
	TSwiftLanguageVAR_KW                = 7
	TSwiftLanguageFUNC_KW               = 8
	TSwiftLanguageSTRUCT_KW             = 9
	TSwiftLanguageIF_KW                 = 10
	TSwiftLanguageELSE_KW               = 11
	TSwiftLanguageSWITCH_KW             = 12
	TSwiftLanguageCASE_KW               = 13
	TSwiftLanguageDEFAULT_KW            = 14
	TSwiftLanguageFOR_KW                = 15
	TSwiftLanguageWHILE_KW              = 16
	TSwiftLanguageBREAK_KW              = 17
	TSwiftLanguageCONTINUE_KW           = 18
	TSwiftLanguageRETURN_KW             = 19
	TSwiftLanguageINTEGER_TYPE          = 20
	TSwiftLanguageFLOAT_TYPE            = 21
	TSwiftLanguageSTRING_TYPE           = 22
	TSwiftLanguageBOOL_TYPE             = 23
	TSwiftLanguageCHARACTER_TYPE        = 24
	TSwiftLanguageINTEGER_LITERAL       = 25
	TSwiftLanguageFLOAT_LITERAL         = 26
	TSwiftLanguageSTRING_LITERAL        = 27
	TSwiftLanguageBOOL_LITERAL          = 28
	TSwiftLanguageNIL_LITERAL           = 29
	TSwiftLanguageID                    = 30
	TSwiftLanguagePLUS                  = 31
	TSwiftLanguageMINUS                 = 32
	TSwiftLanguageMULT                  = 33
	TSwiftLanguageDIV                   = 34
	TSwiftLanguageMOD                   = 35
	TSwiftLanguageEQUALS                = 36
	TSwiftLanguageEQUALS_EQUALS         = 37
	TSwiftLanguageNOT_EQUALS            = 38
	TSwiftLanguageLESS_THAN             = 39
	TSwiftLanguageLESS_THAN_OR_EQUAL    = 40
	TSwiftLanguageGREATER_THAN          = 41
	TSwiftLanguageGREATER_THAN_OR_EQUAL = 42
	TSwiftLanguageAND                   = 43
	TSwiftLanguageOR                    = 44
	TSwiftLanguageNOT                   = 45
	TSwiftLanguageLPAREN                = 46
	TSwiftLanguageRPAREN                = 47
	TSwiftLanguageLBRACE                = 48
	TSwiftLanguageRBRACE                = 49
	TSwiftLanguageLBRACK                = 50
	TSwiftLanguageRBRACK                = 51
	TSwiftLanguageCOMMA                 = 52
	TSwiftLanguageDOT                   = 53
	TSwiftLanguageCOLON                 = 54
	TSwiftLanguageARROW                 = 55
	TSwiftLanguageINTERROGATION         = 56
)

// TSwiftLanguage rules.
const (
	TSwiftLanguageRULE_program        = 0
	TSwiftLanguageRULE_delimiter      = 1
	TSwiftLanguageRULE_stmt           = 2
	TSwiftLanguageRULE_decl_stmt      = 3
	TSwiftLanguageRULE_var_type       = 4
	TSwiftLanguageRULE_primitive_type = 5
	TSwiftLanguageRULE_assign_stmt    = 6
	TSwiftLanguageRULE_literal        = 7
	TSwiftLanguageRULE_expr           = 8
	TSwiftLanguageRULE_if_stmt        = 9
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	AllDelimiter() []IDelimiterContext
	Delimiter(i int) IDelimiterContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ProgramContext) AllDelimiter() []IDelimiterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDelimiterContext); ok {
			len++
		}
	}

	tst := make([]IDelimiterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDelimiterContext); ok {
			tst[i] = t.(IDelimiterContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Delimiter(i int) IDelimiterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDelimiterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDelimiterContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TSwiftLanguageRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1073742016) != 0 {
		{
			p.SetState(20)
			p.Stmt()
		}
		{
			p.SetState(21)
			p.Delimiter()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDelimiterContext is an interface to support dynamic dispatch.
type IDelimiterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEMICOLON() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	EOF() antlr.TerminalNode

	// IsDelimiterContext differentiates from other interfaces.
	IsDelimiterContext()
}

type DelimiterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelimiterContext() *DelimiterContext {
	var p = new(DelimiterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_delimiter
	return p
}

func InitEmptyDelimiterContext(p *DelimiterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_delimiter
}

func (*DelimiterContext) IsDelimiterContext() {}

func NewDelimiterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelimiterContext {
	var p = new(DelimiterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_delimiter

	return p
}

func (s *DelimiterContext) GetParser() antlr.Parser { return s.parser }

func (s *DelimiterContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageSEMICOLON, 0)
}

func (s *DelimiterContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(TSwiftLanguageNEWLINE)
}

func (s *DelimiterContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageNEWLINE, i)
}

func (s *DelimiterContext) EOF() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageEOF, 0)
}

func (s *DelimiterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelimiterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelimiterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitDelimiter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Delimiter() (localctx IDelimiterContext) {
	localctx = NewDelimiterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TSwiftLanguageRULE_delimiter)
	var _la int

	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSwiftLanguageSEMICOLON:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Match(TSwiftLanguageSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSwiftLanguageNEWLINE:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == TSwiftLanguageNEWLINE {
			{
				p.SetState(29)
				p.Match(TSwiftLanguageNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(32)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case TSwiftLanguageEOF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(34)
			p.Match(TSwiftLanguageEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Decl_stmt() IDecl_stmtContext
	Assign_stmt() IAssign_stmtContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Decl_stmt() IDecl_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecl_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecl_stmtContext)
}

func (s *StmtContext) Assign_stmt() IAssign_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_stmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TSwiftLanguageRULE_stmt)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSwiftLanguageLET_KW, TSwiftLanguageVAR_KW:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Decl_stmt()
		}

	case TSwiftLanguageID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.Assign_stmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDecl_stmtContext is an interface to support dynamic dispatch.
type IDecl_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDecl_stmtContext differentiates from other interfaces.
	IsDecl_stmtContext()
}

type Decl_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecl_stmtContext() *Decl_stmtContext {
	var p = new(Decl_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_decl_stmt
	return p
}

func InitEmptyDecl_stmtContext(p *Decl_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_decl_stmt
}

func (*Decl_stmtContext) IsDecl_stmtContext() {}

func NewDecl_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Decl_stmtContext {
	var p = new(Decl_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_decl_stmt

	return p
}

func (s *Decl_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Decl_stmtContext) CopyAll(ctx *Decl_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Decl_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decl_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeValueDeclContext struct {
	Decl_stmtContext
}

func NewTypeValueDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeValueDeclContext {
	var p = new(TypeValueDeclContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *TypeValueDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeValueDeclContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *TypeValueDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, 0)
}

func (s *TypeValueDeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageCOLON, 0)
}

func (s *TypeValueDeclContext) Primitive_type() IPrimitive_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitive_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitive_typeContext)
}

func (s *TypeValueDeclContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageEQUALS, 0)
}

func (s *TypeValueDeclContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TypeValueDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitTypeValueDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeDeclContext struct {
	Decl_stmtContext
}

func NewTypeDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDeclContext {
	var p = new(TypeDeclContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *TypeDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, 0)
}

func (s *TypeDeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageCOLON, 0)
}

func (s *TypeDeclContext) Primitive_type() IPrimitive_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitive_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitive_typeContext)
}

func (s *TypeDeclContext) INTERROGATION() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageINTERROGATION, 0)
}

func (s *TypeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitTypeDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueDeclContext struct {
	Decl_stmtContext
}

func NewValueDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueDeclContext {
	var p = new(ValueDeclContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *ValueDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueDeclContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *ValueDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, 0)
}

func (s *ValueDeclContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageEQUALS, 0)
}

func (s *ValueDeclContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ValueDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitValueDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Decl_stmt() (localctx IDecl_stmtContext) {
	localctx = NewDecl_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TSwiftLanguageRULE_decl_stmt)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeValueDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Var_type()
		}
		{
			p.SetState(42)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)
			p.Match(TSwiftLanguageCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.Primitive_type()
		}
		{
			p.SetState(45)
			p.Match(TSwiftLanguageEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.expr(0)
		}

	case 2:
		localctx = NewValueDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.Var_type()
		}
		{
			p.SetState(49)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Match(TSwiftLanguageEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.expr(0)
		}

	case 3:
		localctx = NewTypeDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(53)
			p.Var_type()
		}
		{
			p.SetState(54)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Match(TSwiftLanguageCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.Primitive_type()
		}
		{
			p.SetState(57)
			p.Match(TSwiftLanguageINTERROGATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_typeContext is an interface to support dynamic dispatch.
type IVar_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_KW() antlr.TerminalNode
	LET_KW() antlr.TerminalNode

	// IsVar_typeContext differentiates from other interfaces.
	IsVar_typeContext()
}

type Var_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_var_type
	return p
}

func InitEmptyVar_typeContext(p *Var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_var_type
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_typeContext) VAR_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageVAR_KW, 0)
}

func (s *Var_typeContext) LET_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLET_KW, 0)
}

func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitVar_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Var_type() (localctx IVar_typeContext) {
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TSwiftLanguageRULE_var_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TSwiftLanguageLET_KW || _la == TSwiftLanguageVAR_KW) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimitive_typeContext is an interface to support dynamic dispatch.
type IPrimitive_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER_TYPE() antlr.TerminalNode
	FLOAT_TYPE() antlr.TerminalNode
	STRING_TYPE() antlr.TerminalNode
	BOOL_TYPE() antlr.TerminalNode

	// IsPrimitive_typeContext differentiates from other interfaces.
	IsPrimitive_typeContext()
}

type Primitive_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitive_typeContext() *Primitive_typeContext {
	var p = new(Primitive_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_primitive_type
	return p
}

func InitEmptyPrimitive_typeContext(p *Primitive_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_primitive_type
}

func (*Primitive_typeContext) IsPrimitive_typeContext() {}

func NewPrimitive_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primitive_typeContext {
	var p = new(Primitive_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_primitive_type

	return p
}

func (s *Primitive_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Primitive_typeContext) INTEGER_TYPE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageINTEGER_TYPE, 0)
}

func (s *Primitive_typeContext) FLOAT_TYPE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageFLOAT_TYPE, 0)
}

func (s *Primitive_typeContext) STRING_TYPE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageSTRING_TYPE, 0)
}

func (s *Primitive_typeContext) BOOL_TYPE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageBOOL_TYPE, 0)
}

func (s *Primitive_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitive_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Primitive_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitPrimitive_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Primitive_type() (localctx IPrimitive_typeContext) {
	localctx = NewPrimitive_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TSwiftLanguageRULE_primitive_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15728640) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssign_stmtContext is an interface to support dynamic dispatch.
type IAssign_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssign_stmtContext differentiates from other interfaces.
	IsAssign_stmtContext()
}

type Assign_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_stmtContext() *Assign_stmtContext {
	var p = new(Assign_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_assign_stmt
	return p
}

func InitEmptyAssign_stmtContext(p *Assign_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_assign_stmt
}

func (*Assign_stmtContext) IsAssign_stmtContext() {}

func NewAssign_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_stmtContext {
	var p = new(Assign_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_assign_stmt

	return p
}

func (s *Assign_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_stmtContext) CopyAll(ctx *Assign_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Assign_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignContext struct {
	Assign_stmtContext
}

func NewAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignContext {
	var p = new(AssignContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ID() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, 0)
}

func (s *AssignContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageEQUALS, 0)
}

func (s *AssignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Assign_stmt() (localctx IAssign_stmtContext) {
	localctx = NewAssign_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TSwiftLanguageRULE_assign_stmt)
	localctx = NewAssignContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)
		p.Match(TSwiftLanguageEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringLiteralContext struct {
	LiteralContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageSTRING_LITERAL, 0)
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolLiteralContext struct {
	LiteralContext
}

func NewBoolLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolLiteralContext {
	var p = new(BoolLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *BoolLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLiteralContext) BOOL_LITERAL() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageBOOL_LITERAL, 0)
}

func (s *BoolLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitBoolLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatLiteralContext struct {
	LiteralContext
}

func NewFloatLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageFLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilLiteralContext struct {
	LiteralContext
}

func NewNilLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilLiteralContext {
	var p = new(NilLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *NilLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilLiteralContext) NIL_LITERAL() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageNIL_LITERAL, 0)
}

func (s *NilLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitNilLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntLiteralContext struct {
	LiteralContext
}

func NewIntLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLiteralContext {
	var p = new(IntLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageINTEGER_LITERAL, 0)
}

func (s *IntLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitIntLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TSwiftLanguageRULE_literal)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSwiftLanguageINTEGER_LITERAL:
		localctx = NewIntLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.Match(TSwiftLanguageINTEGER_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSwiftLanguageFLOAT_LITERAL:
		localctx = NewFloatLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.Match(TSwiftLanguageFLOAT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSwiftLanguageSTRING_LITERAL:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(71)
			p.Match(TSwiftLanguageSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSwiftLanguageBOOL_LITERAL:
		localctx = NewBoolLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(72)
			p.Match(TSwiftLanguageBOOL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSwiftLanguageNIL_LITERAL:
		localctx = NewNilLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(73)
			p.Match(TSwiftLanguageNIL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LiteralExpContext struct {
	ExprContext
}

func NewLiteralExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpContext {
	var p = new(LiteralExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LiteralExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitLiteralExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExpContext struct {
	ExprContext
}

func NewIdExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExpContext {
	var p = new(IdExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExpContext) ID() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, 0)
}

func (s *IdExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitIdExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExpContext struct {
	ExprContext
}

func NewParenExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpContext {
	var p = new(ParenExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParenExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLPAREN, 0)
}

func (s *ParenExpContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenExpContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRPAREN, 0)
}

func (s *ParenExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitParenExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExpContext struct {
	ExprContext
	op antlr.Token
}

func NewUnaryExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExpContext {
	var p = new(UnaryExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnaryExpContext) GetOp() antlr.Token { return s.op }

func (s *UnaryExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryExpContext) NOT() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageNOT, 0)
}

func (s *UnaryExpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageMINUS, 0)
}

func (s *UnaryExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitUnaryExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryExpContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewBinaryExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpContext {
	var p = new(BinaryExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BinaryExpContext) GetOp() antlr.Token { return s.op }

func (s *BinaryExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryExpContext) GetLeft() IExprContext { return s.left }

func (s *BinaryExpContext) GetRight() IExprContext { return s.right }

func (s *BinaryExpContext) SetLeft(v IExprContext) { s.left = v }

func (s *BinaryExpContext) SetRight(v IExprContext) { s.right = v }

func (s *BinaryExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExpContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BinaryExpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguagePLUS, 0)
}

func (s *BinaryExpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageMINUS, 0)
}

func (s *BinaryExpContext) MULT() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageMULT, 0)
}

func (s *BinaryExpContext) DIV() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageDIV, 0)
}

func (s *BinaryExpContext) EQUALS_EQUALS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageEQUALS_EQUALS, 0)
}

func (s *BinaryExpContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageNOT_EQUALS, 0)
}

func (s *BinaryExpContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLESS_THAN, 0)
}

func (s *BinaryExpContext) LESS_THAN_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLESS_THAN_OR_EQUAL, 0)
}

func (s *BinaryExpContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageGREATER_THAN, 0)
}

func (s *BinaryExpContext) GREATER_THAN_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageGREATER_THAN_OR_EQUAL, 0)
}

func (s *BinaryExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitBinaryExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *TSwiftLanguage) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, TSwiftLanguageRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSwiftLanguageMINUS, TSwiftLanguageNOT:
		localctx = NewUnaryExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(77)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == TSwiftLanguageMINUS || _la == TSwiftLanguageNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(78)
			p.expr(4)
		}

	case TSwiftLanguageLPAREN:
		localctx = NewParenExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(79)
			p.Match(TSwiftLanguageLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.expr(0)
		}
		{
			p.SetState(81)
			p.Match(TSwiftLanguageRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSwiftLanguageID:
		localctx = NewIdExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(83)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSwiftLanguageINTEGER_LITERAL, TSwiftLanguageFLOAT_LITERAL, TSwiftLanguageSTRING_LITERAL, TSwiftLanguageBOOL_LITERAL, TSwiftLanguageNIL_LITERAL:
		localctx = NewLiteralExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(84)
			p.Literal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(99)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TSwiftLanguageRULE_expr)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(88)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TSwiftLanguagePLUS || _la == TSwiftLanguageMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(89)

					var _x = p.expr(9)

					localctx.(*BinaryExpContext).right = _x
				}

			case 2:
				localctx = NewBinaryExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TSwiftLanguageRULE_expr)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(91)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TSwiftLanguageMULT || _la == TSwiftLanguageDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(92)

					var _x = p.expr(8)

					localctx.(*BinaryExpContext).right = _x
				}

			case 3:
				localctx = NewBinaryExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TSwiftLanguageRULE_expr)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(94)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TSwiftLanguageEQUALS_EQUALS || _la == TSwiftLanguageNOT_EQUALS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(95)

					var _x = p.expr(7)

					localctx.(*BinaryExpContext).right = _x
				}

			case 4:
				localctx = NewBinaryExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TSwiftLanguageRULE_expr)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(97)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8246337208320) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(98)

					var _x = p.expr(6)

					localctx.(*BinaryExpContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF_KW() antlr.TerminalNode
	Expr() IExprContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_if_stmt
	return p
}

func InitEmptyIf_stmtContext(p *If_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_if_stmt
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) IF_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageIF_KW, 0)
}

func (s *If_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *If_stmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLBRACE, 0)
}

func (s *If_stmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRBRACE, 0)
}

func (s *If_stmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *If_stmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitIf_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TSwiftLanguageRULE_if_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(TSwiftLanguageIF_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.expr(0)
	}
	{
		p.SetState(106)
		p.Match(TSwiftLanguageLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1073742016) != 0 {
		{
			p.SetState(107)
			p.Stmt()
		}

		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(113)
		p.Match(TSwiftLanguageRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *TSwiftLanguage) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 8:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TSwiftLanguage) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
