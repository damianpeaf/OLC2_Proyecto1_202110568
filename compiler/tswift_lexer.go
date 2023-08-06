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
		"']'", "','", "'.'", "':'", "'->'", "'?'", "'&'",
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
		"RBRACK", "COMMA", "DOT", "COLON", "ARROW", "INTERROGATION", "ANPERSAND",
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
		"RBRACK", "COMMA", "DOT", "COLON", "ARROW", "INTERROGATION", "ANPERSAND",
		"ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 60, 398, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 1, 0, 4, 0, 125, 8, 0,
		11, 0, 12, 0, 126, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 135, 8, 1,
		10, 1, 12, 1, 138, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		5, 2, 148, 8, 2, 10, 2, 12, 2, 151, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 4, 25, 282, 8, 25, 11, 25, 12, 25, 283,
		1, 26, 4, 26, 287, 8, 26, 11, 26, 12, 26, 288, 1, 26, 1, 26, 4, 26, 293,
		8, 26, 11, 26, 12, 26, 294, 1, 27, 1, 27, 1, 27, 5, 27, 300, 8, 27, 10,
		27, 12, 27, 303, 9, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 316, 8, 28, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 30, 1, 30, 5, 30, 324, 8, 30, 10, 30, 12, 30, 327, 9, 30, 1, 31,
		1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1,
		36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 40,
		1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1,
		44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48,
		1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1,
		53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 58,
		1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 60, 2, 136, 149, 0, 61, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30,
		61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39,
		79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48,
		97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111, 56, 113,
		57, 115, 58, 117, 59, 119, 60, 121, 0, 1, 0, 6, 3, 0, 9, 10, 13, 13, 32,
		32, 1, 0, 48, 57, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 3, 0, 65, 90, 95,
		95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 8, 0, 34, 34, 39, 39,
		92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 406, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0,
		0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1,
		0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71,
		1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0,
		79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0,
		0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0,
		0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1,
		0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0,
		109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0,
		0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 1, 124, 1, 0, 0, 0, 3, 130,
		1, 0, 0, 0, 5, 143, 1, 0, 0, 0, 7, 157, 1, 0, 0, 0, 9, 159, 1, 0, 0, 0,
		11, 163, 1, 0, 0, 0, 13, 167, 1, 0, 0, 0, 15, 172, 1, 0, 0, 0, 17, 179,
		1, 0, 0, 0, 19, 182, 1, 0, 0, 0, 21, 187, 1, 0, 0, 0, 23, 194, 1, 0, 0,
		0, 25, 199, 1, 0, 0, 0, 27, 207, 1, 0, 0, 0, 29, 211, 1, 0, 0, 0, 31, 217,
		1, 0, 0, 0, 33, 223, 1, 0, 0, 0, 35, 232, 1, 0, 0, 0, 37, 239, 1, 0, 0,
		0, 39, 245, 1, 0, 0, 0, 41, 248, 1, 0, 0, 0, 43, 252, 1, 0, 0, 0, 45, 258,
		1, 0, 0, 0, 47, 265, 1, 0, 0, 0, 49, 270, 1, 0, 0, 0, 51, 281, 1, 0, 0,
		0, 53, 286, 1, 0, 0, 0, 55, 296, 1, 0, 0, 0, 57, 315, 1, 0, 0, 0, 59, 317,
		1, 0, 0, 0, 61, 321, 1, 0, 0, 0, 63, 328, 1, 0, 0, 0, 65, 330, 1, 0, 0,
		0, 67, 332, 1, 0, 0, 0, 69, 334, 1, 0, 0, 0, 71, 336, 1, 0, 0, 0, 73, 338,
		1, 0, 0, 0, 75, 340, 1, 0, 0, 0, 77, 343, 1, 0, 0, 0, 79, 346, 1, 0, 0,
		0, 81, 349, 1, 0, 0, 0, 83, 352, 1, 0, 0, 0, 85, 354, 1, 0, 0, 0, 87, 357,
		1, 0, 0, 0, 89, 359, 1, 0, 0, 0, 91, 362, 1, 0, 0, 0, 93, 365, 1, 0, 0,
		0, 95, 368, 1, 0, 0, 0, 97, 370, 1, 0, 0, 0, 99, 372, 1, 0, 0, 0, 101,
		374, 1, 0, 0, 0, 103, 376, 1, 0, 0, 0, 105, 378, 1, 0, 0, 0, 107, 380,
		1, 0, 0, 0, 109, 382, 1, 0, 0, 0, 111, 384, 1, 0, 0, 0, 113, 386, 1, 0,
		0, 0, 115, 388, 1, 0, 0, 0, 117, 391, 1, 0, 0, 0, 119, 393, 1, 0, 0, 0,
		121, 395, 1, 0, 0, 0, 123, 125, 7, 0, 0, 0, 124, 123, 1, 0, 0, 0, 125,
		126, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128,
		1, 0, 0, 0, 128, 129, 6, 0, 0, 0, 129, 2, 1, 0, 0, 0, 130, 131, 5, 47,
		0, 0, 131, 132, 5, 47, 0, 0, 132, 136, 1, 0, 0, 0, 133, 135, 9, 0, 0, 0,
		134, 133, 1, 0, 0, 0, 135, 138, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 136,
		134, 1, 0, 0, 0, 137, 139, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 140,
		5, 10, 0, 0, 140, 141, 1, 0, 0, 0, 141, 142, 6, 1, 0, 0, 142, 4, 1, 0,
		0, 0, 143, 144, 5, 47, 0, 0, 144, 145, 5, 42, 0, 0, 145, 149, 1, 0, 0,
		0, 146, 148, 9, 0, 0, 0, 147, 146, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149,
		150, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150, 152, 1, 0, 0, 0, 151, 149,
		1, 0, 0, 0, 152, 153, 5, 42, 0, 0, 153, 154, 5, 47, 0, 0, 154, 155, 1,
		0, 0, 0, 155, 156, 6, 2, 0, 0, 156, 6, 1, 0, 0, 0, 157, 158, 5, 59, 0,
		0, 158, 8, 1, 0, 0, 0, 159, 160, 5, 108, 0, 0, 160, 161, 5, 101, 0, 0,
		161, 162, 5, 116, 0, 0, 162, 10, 1, 0, 0, 0, 163, 164, 5, 118, 0, 0, 164,
		165, 5, 97, 0, 0, 165, 166, 5, 114, 0, 0, 166, 12, 1, 0, 0, 0, 167, 168,
		5, 102, 0, 0, 168, 169, 5, 117, 0, 0, 169, 170, 5, 110, 0, 0, 170, 171,
		5, 99, 0, 0, 171, 14, 1, 0, 0, 0, 172, 173, 5, 115, 0, 0, 173, 174, 5,
		116, 0, 0, 174, 175, 5, 114, 0, 0, 175, 176, 5, 117, 0, 0, 176, 177, 5,
		99, 0, 0, 177, 178, 5, 116, 0, 0, 178, 16, 1, 0, 0, 0, 179, 180, 5, 105,
		0, 0, 180, 181, 5, 102, 0, 0, 181, 18, 1, 0, 0, 0, 182, 183, 5, 101, 0,
		0, 183, 184, 5, 108, 0, 0, 184, 185, 5, 115, 0, 0, 185, 186, 5, 101, 0,
		0, 186, 20, 1, 0, 0, 0, 187, 188, 5, 115, 0, 0, 188, 189, 5, 119, 0, 0,
		189, 190, 5, 105, 0, 0, 190, 191, 5, 116, 0, 0, 191, 192, 5, 99, 0, 0,
		192, 193, 5, 104, 0, 0, 193, 22, 1, 0, 0, 0, 194, 195, 5, 99, 0, 0, 195,
		196, 5, 97, 0, 0, 196, 197, 5, 115, 0, 0, 197, 198, 5, 101, 0, 0, 198,
		24, 1, 0, 0, 0, 199, 200, 5, 100, 0, 0, 200, 201, 5, 101, 0, 0, 201, 202,
		5, 102, 0, 0, 202, 203, 5, 97, 0, 0, 203, 204, 5, 117, 0, 0, 204, 205,
		5, 108, 0, 0, 205, 206, 5, 116, 0, 0, 206, 26, 1, 0, 0, 0, 207, 208, 5,
		102, 0, 0, 208, 209, 5, 111, 0, 0, 209, 210, 5, 114, 0, 0, 210, 28, 1,
		0, 0, 0, 211, 212, 5, 119, 0, 0, 212, 213, 5, 104, 0, 0, 213, 214, 5, 105,
		0, 0, 214, 215, 5, 108, 0, 0, 215, 216, 5, 101, 0, 0, 216, 30, 1, 0, 0,
		0, 217, 218, 5, 98, 0, 0, 218, 219, 5, 114, 0, 0, 219, 220, 5, 101, 0,
		0, 220, 221, 5, 97, 0, 0, 221, 222, 5, 107, 0, 0, 222, 32, 1, 0, 0, 0,
		223, 224, 5, 99, 0, 0, 224, 225, 5, 111, 0, 0, 225, 226, 5, 110, 0, 0,
		226, 227, 5, 116, 0, 0, 227, 228, 5, 105, 0, 0, 228, 229, 5, 110, 0, 0,
		229, 230, 5, 117, 0, 0, 230, 231, 5, 101, 0, 0, 231, 34, 1, 0, 0, 0, 232,
		233, 5, 114, 0, 0, 233, 234, 5, 101, 0, 0, 234, 235, 5, 116, 0, 0, 235,
		236, 5, 117, 0, 0, 236, 237, 5, 114, 0, 0, 237, 238, 5, 110, 0, 0, 238,
		36, 1, 0, 0, 0, 239, 240, 5, 103, 0, 0, 240, 241, 5, 117, 0, 0, 241, 242,
		5, 97, 0, 0, 242, 243, 5, 114, 0, 0, 243, 244, 5, 100, 0, 0, 244, 38, 1,
		0, 0, 0, 245, 246, 5, 105, 0, 0, 246, 247, 5, 110, 0, 0, 247, 40, 1, 0,
		0, 0, 248, 249, 5, 73, 0, 0, 249, 250, 5, 110, 0, 0, 250, 251, 5, 116,
		0, 0, 251, 42, 1, 0, 0, 0, 252, 253, 5, 70, 0, 0, 253, 254, 5, 108, 0,
		0, 254, 255, 5, 111, 0, 0, 255, 256, 5, 97, 0, 0, 256, 257, 5, 116, 0,
		0, 257, 44, 1, 0, 0, 0, 258, 259, 5, 83, 0, 0, 259, 260, 5, 116, 0, 0,
		260, 261, 5, 114, 0, 0, 261, 262, 5, 105, 0, 0, 262, 263, 5, 110, 0, 0,
		263, 264, 5, 103, 0, 0, 264, 46, 1, 0, 0, 0, 265, 266, 5, 66, 0, 0, 266,
		267, 5, 111, 0, 0, 267, 268, 5, 111, 0, 0, 268, 269, 5, 108, 0, 0, 269,
		48, 1, 0, 0, 0, 270, 271, 5, 67, 0, 0, 271, 272, 5, 104, 0, 0, 272, 273,
		5, 97, 0, 0, 273, 274, 5, 114, 0, 0, 274, 275, 5, 97, 0, 0, 275, 276, 5,
		99, 0, 0, 276, 277, 5, 116, 0, 0, 277, 278, 5, 101, 0, 0, 278, 279, 5,
		114, 0, 0, 279, 50, 1, 0, 0, 0, 280, 282, 7, 1, 0, 0, 281, 280, 1, 0, 0,
		0, 282, 283, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284,
		52, 1, 0, 0, 0, 285, 287, 7, 1, 0, 0, 286, 285, 1, 0, 0, 0, 287, 288, 1,
		0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 290, 1, 0, 0,
		0, 290, 292, 5, 46, 0, 0, 291, 293, 7, 1, 0, 0, 292, 291, 1, 0, 0, 0, 293,
		294, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 54, 1,
		0, 0, 0, 296, 301, 5, 34, 0, 0, 297, 300, 8, 2, 0, 0, 298, 300, 3, 121,
		60, 0, 299, 297, 1, 0, 0, 0, 299, 298, 1, 0, 0, 0, 300, 303, 1, 0, 0, 0,
		301, 299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 304, 1, 0, 0, 0, 303,
		301, 1, 0, 0, 0, 304, 305, 5, 34, 0, 0, 305, 56, 1, 0, 0, 0, 306, 307,
		5, 116, 0, 0, 307, 308, 5, 114, 0, 0, 308, 309, 5, 117, 0, 0, 309, 316,
		5, 101, 0, 0, 310, 311, 5, 102, 0, 0, 311, 312, 5, 97, 0, 0, 312, 313,
		5, 108, 0, 0, 313, 314, 5, 115, 0, 0, 314, 316, 5, 101, 0, 0, 315, 306,
		1, 0, 0, 0, 315, 310, 1, 0, 0, 0, 316, 58, 1, 0, 0, 0, 317, 318, 5, 110,
		0, 0, 318, 319, 5, 105, 0, 0, 319, 320, 5, 108, 0, 0, 320, 60, 1, 0, 0,
		0, 321, 325, 7, 3, 0, 0, 322, 324, 7, 4, 0, 0, 323, 322, 1, 0, 0, 0, 324,
		327, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 62, 1,
		0, 0, 0, 327, 325, 1, 0, 0, 0, 328, 329, 5, 43, 0, 0, 329, 64, 1, 0, 0,
		0, 330, 331, 5, 45, 0, 0, 331, 66, 1, 0, 0, 0, 332, 333, 5, 42, 0, 0, 333,
		68, 1, 0, 0, 0, 334, 335, 5, 47, 0, 0, 335, 70, 1, 0, 0, 0, 336, 337, 5,
		37, 0, 0, 337, 72, 1, 0, 0, 0, 338, 339, 5, 61, 0, 0, 339, 74, 1, 0, 0,
		0, 340, 341, 5, 43, 0, 0, 341, 342, 5, 61, 0, 0, 342, 76, 1, 0, 0, 0, 343,
		344, 5, 45, 0, 0, 344, 345, 5, 61, 0, 0, 345, 78, 1, 0, 0, 0, 346, 347,
		5, 61, 0, 0, 347, 348, 5, 61, 0, 0, 348, 80, 1, 0, 0, 0, 349, 350, 5, 33,
		0, 0, 350, 351, 5, 61, 0, 0, 351, 82, 1, 0, 0, 0, 352, 353, 5, 60, 0, 0,
		353, 84, 1, 0, 0, 0, 354, 355, 5, 60, 0, 0, 355, 356, 5, 61, 0, 0, 356,
		86, 1, 0, 0, 0, 357, 358, 5, 62, 0, 0, 358, 88, 1, 0, 0, 0, 359, 360, 5,
		62, 0, 0, 360, 361, 5, 61, 0, 0, 361, 90, 1, 0, 0, 0, 362, 363, 5, 38,
		0, 0, 363, 364, 5, 38, 0, 0, 364, 92, 1, 0, 0, 0, 365, 366, 5, 124, 0,
		0, 366, 367, 5, 124, 0, 0, 367, 94, 1, 0, 0, 0, 368, 369, 5, 33, 0, 0,
		369, 96, 1, 0, 0, 0, 370, 371, 5, 40, 0, 0, 371, 98, 1, 0, 0, 0, 372, 373,
		5, 41, 0, 0, 373, 100, 1, 0, 0, 0, 374, 375, 5, 123, 0, 0, 375, 102, 1,
		0, 0, 0, 376, 377, 5, 125, 0, 0, 377, 104, 1, 0, 0, 0, 378, 379, 5, 91,
		0, 0, 379, 106, 1, 0, 0, 0, 380, 381, 5, 93, 0, 0, 381, 108, 1, 0, 0, 0,
		382, 383, 5, 44, 0, 0, 383, 110, 1, 0, 0, 0, 384, 385, 5, 46, 0, 0, 385,
		112, 1, 0, 0, 0, 386, 387, 5, 58, 0, 0, 387, 114, 1, 0, 0, 0, 388, 389,
		5, 45, 0, 0, 389, 390, 5, 62, 0, 0, 390, 116, 1, 0, 0, 0, 391, 392, 5,
		63, 0, 0, 392, 118, 1, 0, 0, 0, 393, 394, 5, 38, 0, 0, 394, 120, 1, 0,
		0, 0, 395, 396, 5, 92, 0, 0, 396, 397, 7, 5, 0, 0, 397, 122, 1, 0, 0, 0,
		11, 0, 126, 136, 149, 283, 288, 294, 299, 301, 315, 325, 1, 6, 0, 0,
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
	TSwiftLexerANPERSAND             = 60
)
