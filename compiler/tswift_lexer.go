// Code generated from compiler\TSwiftLexer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package compiler

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type TSwiftLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var TSwiftLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tswiftlexerLexerInit() {
	staticData := &TSwiftLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "';'", "'let'", "'var'", "'func'", "'struct'", "'if'",
		"'else'", "'switch'", "'case'", "'default'", "'for'", "'while'", "'break'",
		"'continue'", "'return'", "'guard'", "'in'", "'Int'", "'Float'", "'String'",
		"'Bool'", "'Character'", "", "", "", "", "'nil'", "", "'+'", "'-'",
		"'*'", "'/'", "'%'", "'='", "'+='", "'-='", "'=='", "'!='", "'<'", "'<='",
		"'>'", "'>='", "'&&'", "'||'", "'!'", "'('", "')'", "'{'", "'}'", "'['",
		"']'", "','", "'.'", "':'", "'->'", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "COMMENT", "MULTILINE_COMMENT", "SEMICOLON", "LET_KW", "VAR_KW",
		"FUNC_KW", "STRUCT_KW", "IF_KW", "ELSE_KW", "SWITCH_KW", "CASE_KW",
		"DEFAULT_KW", "FOR_KW", "WHILE_KW", "BREAK_KW", "CONTINUE_KW", "RETURN_KW",
		"GUARD_KW", "IN_KW", "INTEGER_TYPE", "FLOAT_TYPE", "STRING_TYPE", "BOOL_TYPE",
		"CHARACTER_TYPE", "INTEGER_LITERAL", "FLOAT_LITERAL", "STRING_LITERAL",
		"BOOL_LITERAL", "NIL_LITERAL", "ID", "PLUS", "MINUS", "MULT", "DIV",
		"MOD", "EQUALS", "PLUS_EQUALS", "MINUS_EQUALS", "EQUALS_EQUALS", "NOT_EQUALS",
		"LESS_THAN", "LESS_THAN_OR_EQUAL", "GREATER_THAN", "GREATER_THAN_OR_EQUAL",
		"AND", "OR", "NOT", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK",
		"RBRACK", "COMMA", "DOT", "COLON", "ARROW", "INTERROGATION",
	}
	staticData.RuleNames = []string{
		"WS", "COMMENT", "MULTILINE_COMMENT", "SEMICOLON", "LET_KW", "VAR_KW",
		"FUNC_KW", "STRUCT_KW", "IF_KW", "ELSE_KW", "SWITCH_KW", "CASE_KW",
		"DEFAULT_KW", "FOR_KW", "WHILE_KW", "BREAK_KW", "CONTINUE_KW", "RETURN_KW",
		"GUARD_KW", "IN_KW", "INTEGER_TYPE", "FLOAT_TYPE", "STRING_TYPE", "BOOL_TYPE",
		"CHARACTER_TYPE", "INTEGER_LITERAL", "FLOAT_LITERAL", "STRING_LITERAL",
		"BOOL_LITERAL", "NIL_LITERAL", "ID", "PLUS", "MINUS", "MULT", "DIV",
		"MOD", "EQUALS", "PLUS_EQUALS", "MINUS_EQUALS", "EQUALS_EQUALS", "NOT_EQUALS",
		"LESS_THAN", "LESS_THAN_OR_EQUAL", "GREATER_THAN", "GREATER_THAN_OR_EQUAL",
		"AND", "OR", "NOT", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK",
		"RBRACK", "COMMA", "DOT", "COLON", "ARROW", "INTERROGATION", "ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 59, 394, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 1, 0, 4, 0, 123, 8, 0, 11, 0, 12, 0,
		124, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 133, 8, 1, 10, 1, 12, 1,
		136, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 146, 8,
		2, 10, 2, 12, 2, 149, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 25, 4, 25, 280, 8, 25, 11, 25, 12, 25, 281, 1, 26, 4, 26,
		285, 8, 26, 11, 26, 12, 26, 286, 1, 26, 1, 26, 4, 26, 291, 8, 26, 11, 26,
		12, 26, 292, 1, 27, 1, 27, 1, 27, 5, 27, 298, 8, 27, 10, 27, 12, 27, 301,
		9, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 3, 28, 314, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30,
		5, 30, 322, 8, 30, 10, 30, 12, 30, 325, 9, 30, 1, 31, 1, 31, 1, 32, 1,
		32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37,
		1, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1,
		41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 45,
		1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1,
		49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54,
		1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1,
		59, 1, 59, 2, 134, 147, 0, 60, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25,
		51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34,
		69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43,
		87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52,
		105, 53, 107, 54, 109, 55, 111, 56, 113, 57, 115, 58, 117, 59, 119, 0,
		1, 0, 6, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 4, 0, 10, 10, 13, 13,
		34, 34, 92, 92, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 8, 0, 34, 34, 39, 39, 92, 92, 98, 98, 102, 102, 110, 110,
		114, 114, 116, 116, 402, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0,
		0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1,
		0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67,
		1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0,
		75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0,
		0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0,
		0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0,
		0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105,
		1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0,
		0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 1, 122, 1,
		0, 0, 0, 3, 128, 1, 0, 0, 0, 5, 141, 1, 0, 0, 0, 7, 155, 1, 0, 0, 0, 9,
		157, 1, 0, 0, 0, 11, 161, 1, 0, 0, 0, 13, 165, 1, 0, 0, 0, 15, 170, 1,
		0, 0, 0, 17, 177, 1, 0, 0, 0, 19, 180, 1, 0, 0, 0, 21, 185, 1, 0, 0, 0,
		23, 192, 1, 0, 0, 0, 25, 197, 1, 0, 0, 0, 27, 205, 1, 0, 0, 0, 29, 209,
		1, 0, 0, 0, 31, 215, 1, 0, 0, 0, 33, 221, 1, 0, 0, 0, 35, 230, 1, 0, 0,
		0, 37, 237, 1, 0, 0, 0, 39, 243, 1, 0, 0, 0, 41, 246, 1, 0, 0, 0, 43, 250,
		1, 0, 0, 0, 45, 256, 1, 0, 0, 0, 47, 263, 1, 0, 0, 0, 49, 268, 1, 0, 0,
		0, 51, 279, 1, 0, 0, 0, 53, 284, 1, 0, 0, 0, 55, 294, 1, 0, 0, 0, 57, 313,
		1, 0, 0, 0, 59, 315, 1, 0, 0, 0, 61, 319, 1, 0, 0, 0, 63, 326, 1, 0, 0,
		0, 65, 328, 1, 0, 0, 0, 67, 330, 1, 0, 0, 0, 69, 332, 1, 0, 0, 0, 71, 334,
		1, 0, 0, 0, 73, 336, 1, 0, 0, 0, 75, 338, 1, 0, 0, 0, 77, 341, 1, 0, 0,
		0, 79, 344, 1, 0, 0, 0, 81, 347, 1, 0, 0, 0, 83, 350, 1, 0, 0, 0, 85, 352,
		1, 0, 0, 0, 87, 355, 1, 0, 0, 0, 89, 357, 1, 0, 0, 0, 91, 360, 1, 0, 0,
		0, 93, 363, 1, 0, 0, 0, 95, 366, 1, 0, 0, 0, 97, 368, 1, 0, 0, 0, 99, 370,
		1, 0, 0, 0, 101, 372, 1, 0, 0, 0, 103, 374, 1, 0, 0, 0, 105, 376, 1, 0,
		0, 0, 107, 378, 1, 0, 0, 0, 109, 380, 1, 0, 0, 0, 111, 382, 1, 0, 0, 0,
		113, 384, 1, 0, 0, 0, 115, 386, 1, 0, 0, 0, 117, 389, 1, 0, 0, 0, 119,
		391, 1, 0, 0, 0, 121, 123, 7, 0, 0, 0, 122, 121, 1, 0, 0, 0, 123, 124,
		1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126, 1, 0,
		0, 0, 126, 127, 6, 0, 0, 0, 127, 2, 1, 0, 0, 0, 128, 129, 5, 47, 0, 0,
		129, 130, 5, 47, 0, 0, 130, 134, 1, 0, 0, 0, 131, 133, 9, 0, 0, 0, 132,
		131, 1, 0, 0, 0, 133, 136, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 134, 132,
		1, 0, 0, 0, 135, 137, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 137, 138, 5, 10,
		0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 6, 1, 0, 0, 140, 4, 1, 0, 0, 0, 141,
		142, 5, 47, 0, 0, 142, 143, 5, 42, 0, 0, 143, 147, 1, 0, 0, 0, 144, 146,
		9, 0, 0, 0, 145, 144, 1, 0, 0, 0, 146, 149, 1, 0, 0, 0, 147, 148, 1, 0,
		0, 0, 147, 145, 1, 0, 0, 0, 148, 150, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0,
		150, 151, 5, 42, 0, 0, 151, 152, 5, 47, 0, 0, 152, 153, 1, 0, 0, 0, 153,
		154, 6, 2, 0, 0, 154, 6, 1, 0, 0, 0, 155, 156, 5, 59, 0, 0, 156, 8, 1,
		0, 0, 0, 157, 158, 5, 108, 0, 0, 158, 159, 5, 101, 0, 0, 159, 160, 5, 116,
		0, 0, 160, 10, 1, 0, 0, 0, 161, 162, 5, 118, 0, 0, 162, 163, 5, 97, 0,
		0, 163, 164, 5, 114, 0, 0, 164, 12, 1, 0, 0, 0, 165, 166, 5, 102, 0, 0,
		166, 167, 5, 117, 0, 0, 167, 168, 5, 110, 0, 0, 168, 169, 5, 99, 0, 0,
		169, 14, 1, 0, 0, 0, 170, 171, 5, 115, 0, 0, 171, 172, 5, 116, 0, 0, 172,
		173, 5, 114, 0, 0, 173, 174, 5, 117, 0, 0, 174, 175, 5, 99, 0, 0, 175,
		176, 5, 116, 0, 0, 176, 16, 1, 0, 0, 0, 177, 178, 5, 105, 0, 0, 178, 179,
		5, 102, 0, 0, 179, 18, 1, 0, 0, 0, 180, 181, 5, 101, 0, 0, 181, 182, 5,
		108, 0, 0, 182, 183, 5, 115, 0, 0, 183, 184, 5, 101, 0, 0, 184, 20, 1,
		0, 0, 0, 185, 186, 5, 115, 0, 0, 186, 187, 5, 119, 0, 0, 187, 188, 5, 105,
		0, 0, 188, 189, 5, 116, 0, 0, 189, 190, 5, 99, 0, 0, 190, 191, 5, 104,
		0, 0, 191, 22, 1, 0, 0, 0, 192, 193, 5, 99, 0, 0, 193, 194, 5, 97, 0, 0,
		194, 195, 5, 115, 0, 0, 195, 196, 5, 101, 0, 0, 196, 24, 1, 0, 0, 0, 197,
		198, 5, 100, 0, 0, 198, 199, 5, 101, 0, 0, 199, 200, 5, 102, 0, 0, 200,
		201, 5, 97, 0, 0, 201, 202, 5, 117, 0, 0, 202, 203, 5, 108, 0, 0, 203,
		204, 5, 116, 0, 0, 204, 26, 1, 0, 0, 0, 205, 206, 5, 102, 0, 0, 206, 207,
		5, 111, 0, 0, 207, 208, 5, 114, 0, 0, 208, 28, 1, 0, 0, 0, 209, 210, 5,
		119, 0, 0, 210, 211, 5, 104, 0, 0, 211, 212, 5, 105, 0, 0, 212, 213, 5,
		108, 0, 0, 213, 214, 5, 101, 0, 0, 214, 30, 1, 0, 0, 0, 215, 216, 5, 98,
		0, 0, 216, 217, 5, 114, 0, 0, 217, 218, 5, 101, 0, 0, 218, 219, 5, 97,
		0, 0, 219, 220, 5, 107, 0, 0, 220, 32, 1, 0, 0, 0, 221, 222, 5, 99, 0,
		0, 222, 223, 5, 111, 0, 0, 223, 224, 5, 110, 0, 0, 224, 225, 5, 116, 0,
		0, 225, 226, 5, 105, 0, 0, 226, 227, 5, 110, 0, 0, 227, 228, 5, 117, 0,
		0, 228, 229, 5, 101, 0, 0, 229, 34, 1, 0, 0, 0, 230, 231, 5, 114, 0, 0,
		231, 232, 5, 101, 0, 0, 232, 233, 5, 116, 0, 0, 233, 234, 5, 117, 0, 0,
		234, 235, 5, 114, 0, 0, 235, 236, 5, 110, 0, 0, 236, 36, 1, 0, 0, 0, 237,
		238, 5, 103, 0, 0, 238, 239, 5, 117, 0, 0, 239, 240, 5, 97, 0, 0, 240,
		241, 5, 114, 0, 0, 241, 242, 5, 100, 0, 0, 242, 38, 1, 0, 0, 0, 243, 244,
		5, 105, 0, 0, 244, 245, 5, 110, 0, 0, 245, 40, 1, 0, 0, 0, 246, 247, 5,
		73, 0, 0, 247, 248, 5, 110, 0, 0, 248, 249, 5, 116, 0, 0, 249, 42, 1, 0,
		0, 0, 250, 251, 5, 70, 0, 0, 251, 252, 5, 108, 0, 0, 252, 253, 5, 111,
		0, 0, 253, 254, 5, 97, 0, 0, 254, 255, 5, 116, 0, 0, 255, 44, 1, 0, 0,
		0, 256, 257, 5, 83, 0, 0, 257, 258, 5, 116, 0, 0, 258, 259, 5, 114, 0,
		0, 259, 260, 5, 105, 0, 0, 260, 261, 5, 110, 0, 0, 261, 262, 5, 103, 0,
		0, 262, 46, 1, 0, 0, 0, 263, 264, 5, 66, 0, 0, 264, 265, 5, 111, 0, 0,
		265, 266, 5, 111, 0, 0, 266, 267, 5, 108, 0, 0, 267, 48, 1, 0, 0, 0, 268,
		269, 5, 67, 0, 0, 269, 270, 5, 104, 0, 0, 270, 271, 5, 97, 0, 0, 271, 272,
		5, 114, 0, 0, 272, 273, 5, 97, 0, 0, 273, 274, 5, 99, 0, 0, 274, 275, 5,
		116, 0, 0, 275, 276, 5, 101, 0, 0, 276, 277, 5, 114, 0, 0, 277, 50, 1,
		0, 0, 0, 278, 280, 7, 1, 0, 0, 279, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0,
		0, 281, 279, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 52, 1, 0, 0, 0, 283,
		285, 7, 1, 0, 0, 284, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 284,
		1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 290, 5, 46,
		0, 0, 289, 291, 7, 1, 0, 0, 290, 289, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0,
		292, 290, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 54, 1, 0, 0, 0, 294, 299,
		5, 34, 0, 0, 295, 298, 8, 2, 0, 0, 296, 298, 3, 119, 59, 0, 297, 295, 1,
		0, 0, 0, 297, 296, 1, 0, 0, 0, 298, 301, 1, 0, 0, 0, 299, 297, 1, 0, 0,
		0, 299, 300, 1, 0, 0, 0, 300, 302, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 302,
		303, 5, 34, 0, 0, 303, 56, 1, 0, 0, 0, 304, 305, 5, 116, 0, 0, 305, 306,
		5, 114, 0, 0, 306, 307, 5, 117, 0, 0, 307, 314, 5, 101, 0, 0, 308, 309,
		5, 102, 0, 0, 309, 310, 5, 97, 0, 0, 310, 311, 5, 108, 0, 0, 311, 312,
		5, 115, 0, 0, 312, 314, 5, 101, 0, 0, 313, 304, 1, 0, 0, 0, 313, 308, 1,
		0, 0, 0, 314, 58, 1, 0, 0, 0, 315, 316, 5, 110, 0, 0, 316, 317, 5, 105,
		0, 0, 317, 318, 5, 108, 0, 0, 318, 60, 1, 0, 0, 0, 319, 323, 7, 3, 0, 0,
		320, 322, 7, 4, 0, 0, 321, 320, 1, 0, 0, 0, 322, 325, 1, 0, 0, 0, 323,
		321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 62, 1, 0, 0, 0, 325, 323, 1,
		0, 0, 0, 326, 327, 5, 43, 0, 0, 327, 64, 1, 0, 0, 0, 328, 329, 5, 45, 0,
		0, 329, 66, 1, 0, 0, 0, 330, 331, 5, 42, 0, 0, 331, 68, 1, 0, 0, 0, 332,
		333, 5, 47, 0, 0, 333, 70, 1, 0, 0, 0, 334, 335, 5, 37, 0, 0, 335, 72,
		1, 0, 0, 0, 336, 337, 5, 61, 0, 0, 337, 74, 1, 0, 0, 0, 338, 339, 5, 43,
		0, 0, 339, 340, 5, 61, 0, 0, 340, 76, 1, 0, 0, 0, 341, 342, 5, 45, 0, 0,
		342, 343, 5, 61, 0, 0, 343, 78, 1, 0, 0, 0, 344, 345, 5, 61, 0, 0, 345,
		346, 5, 61, 0, 0, 346, 80, 1, 0, 0, 0, 347, 348, 5, 33, 0, 0, 348, 349,
		5, 61, 0, 0, 349, 82, 1, 0, 0, 0, 350, 351, 5, 60, 0, 0, 351, 84, 1, 0,
		0, 0, 352, 353, 5, 60, 0, 0, 353, 354, 5, 61, 0, 0, 354, 86, 1, 0, 0, 0,
		355, 356, 5, 62, 0, 0, 356, 88, 1, 0, 0, 0, 357, 358, 5, 62, 0, 0, 358,
		359, 5, 61, 0, 0, 359, 90, 1, 0, 0, 0, 360, 361, 5, 38, 0, 0, 361, 362,
		5, 38, 0, 0, 362, 92, 1, 0, 0, 0, 363, 364, 5, 124, 0, 0, 364, 365, 5,
		124, 0, 0, 365, 94, 1, 0, 0, 0, 366, 367, 5, 33, 0, 0, 367, 96, 1, 0, 0,
		0, 368, 369, 5, 40, 0, 0, 369, 98, 1, 0, 0, 0, 370, 371, 5, 41, 0, 0, 371,
		100, 1, 0, 0, 0, 372, 373, 5, 123, 0, 0, 373, 102, 1, 0, 0, 0, 374, 375,
		5, 125, 0, 0, 375, 104, 1, 0, 0, 0, 376, 377, 5, 91, 0, 0, 377, 106, 1,
		0, 0, 0, 378, 379, 5, 93, 0, 0, 379, 108, 1, 0, 0, 0, 380, 381, 5, 44,
		0, 0, 381, 110, 1, 0, 0, 0, 382, 383, 5, 46, 0, 0, 383, 112, 1, 0, 0, 0,
		384, 385, 5, 58, 0, 0, 385, 114, 1, 0, 0, 0, 386, 387, 5, 45, 0, 0, 387,
		388, 5, 62, 0, 0, 388, 116, 1, 0, 0, 0, 389, 390, 5, 63, 0, 0, 390, 118,
		1, 0, 0, 0, 391, 392, 5, 92, 0, 0, 392, 393, 7, 5, 0, 0, 393, 120, 1, 0,
		0, 0, 11, 0, 124, 134, 147, 281, 286, 292, 297, 299, 313, 323, 1, 6, 0,
		0,
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

// TSwiftLexerInit initializes any static state used to implement TSwiftLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewTSwiftLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TSwiftLexerInit() {
	staticData := &TSwiftLexerLexerStaticData
	staticData.once.Do(tswiftlexerLexerInit)
}

// NewTSwiftLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewTSwiftLexer(input antlr.CharStream) *TSwiftLexer {
	TSwiftLexerInit()
	l := new(TSwiftLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &TSwiftLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "TSwiftLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TSwiftLexer tokens.
const (
	TSwiftLexerWS                    = 1
	TSwiftLexerCOMMENT               = 2
	TSwiftLexerMULTILINE_COMMENT     = 3
	TSwiftLexerSEMICOLON             = 4
	TSwiftLexerLET_KW                = 5
	TSwiftLexerVAR_KW                = 6
	TSwiftLexerFUNC_KW               = 7
	TSwiftLexerSTRUCT_KW             = 8
	TSwiftLexerIF_KW                 = 9
	TSwiftLexerELSE_KW               = 10
	TSwiftLexerSWITCH_KW             = 11
	TSwiftLexerCASE_KW               = 12
	TSwiftLexerDEFAULT_KW            = 13
	TSwiftLexerFOR_KW                = 14
	TSwiftLexerWHILE_KW              = 15
	TSwiftLexerBREAK_KW              = 16
	TSwiftLexerCONTINUE_KW           = 17
	TSwiftLexerRETURN_KW             = 18
	TSwiftLexerGUARD_KW              = 19
	TSwiftLexerIN_KW                 = 20
	TSwiftLexerINTEGER_TYPE          = 21
	TSwiftLexerFLOAT_TYPE            = 22
	TSwiftLexerSTRING_TYPE           = 23
	TSwiftLexerBOOL_TYPE             = 24
	TSwiftLexerCHARACTER_TYPE        = 25
	TSwiftLexerINTEGER_LITERAL       = 26
	TSwiftLexerFLOAT_LITERAL         = 27
	TSwiftLexerSTRING_LITERAL        = 28
	TSwiftLexerBOOL_LITERAL          = 29
	TSwiftLexerNIL_LITERAL           = 30
	TSwiftLexerID                    = 31
	TSwiftLexerPLUS                  = 32
	TSwiftLexerMINUS                 = 33
	TSwiftLexerMULT                  = 34
	TSwiftLexerDIV                   = 35
	TSwiftLexerMOD                   = 36
	TSwiftLexerEQUALS                = 37
	TSwiftLexerPLUS_EQUALS           = 38
	TSwiftLexerMINUS_EQUALS          = 39
	TSwiftLexerEQUALS_EQUALS         = 40
	TSwiftLexerNOT_EQUALS            = 41
	TSwiftLexerLESS_THAN             = 42
	TSwiftLexerLESS_THAN_OR_EQUAL    = 43
	TSwiftLexerGREATER_THAN          = 44
	TSwiftLexerGREATER_THAN_OR_EQUAL = 45
	TSwiftLexerAND                   = 46
	TSwiftLexerOR                    = 47
	TSwiftLexerNOT                   = 48
	TSwiftLexerLPAREN                = 49
	TSwiftLexerRPAREN                = 50
	TSwiftLexerLBRACE                = 51
	TSwiftLexerRBRACE                = 52
	TSwiftLexerLBRACK                = 53
	TSwiftLexerRBRACK                = 54
	TSwiftLexerCOMMA                 = 55
	TSwiftLexerDOT                   = 56
	TSwiftLexerCOLON                 = 57
	TSwiftLexerARROW                 = 58
	TSwiftLexerINTERROGATION         = 59
)
