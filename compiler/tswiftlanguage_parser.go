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
		"", "", "", "", "';'", "'let'", "'var'", "'func'", "'struct'", "'if'",
		"'else'", "'switch'", "'case'", "'default'", "'for'", "'while'", "'break'",
		"'continue'", "'return'", "'guard'", "'inout'", "'mutating'", "'in'",
		"", "", "", "", "'nil'", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'='",
		"'+='", "'-='", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'&&'",
		"'||'", "'!'", "'('", "')'", "'{'", "'}'", "'['", "']'", "','", "'.'",
		"':'", "'->'", "'?'", "'&'",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "COMMENT", "MULTILINE_COMMENT", "SEMICOLON", "LET_KW", "VAR_KW",
		"FUNC_KW", "STRUCT_KW", "IF_KW", "ELSE_KW", "SWITCH_KW", "CASE_KW",
		"DEFAULT_KW", "FOR_KW", "WHILE_KW", "BREAK_KW", "CONTINUE_KW", "RETURN_KW",
		"GUARD_KW", "INOUT_KW", "MUTATING_KW", "IN_KW", "INTEGER_LITERAL", "FLOAT_LITERAL",
		"STRING_LITERAL", "BOOL_LITERAL", "NIL_LITERAL", "ID", "PLUS", "MINUS",
		"MULT", "DIV", "MOD", "EQUALS", "PLUS_EQUALS", "MINUS_EQUALS", "EQUALS_EQUALS",
		"NOT_EQUALS", "LESS_THAN", "LESS_THAN_OR_EQUAL", "GREATER_THAN", "GREATER_THAN_OR_EQUAL",
		"AND", "OR", "NOT", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK",
		"RBRACK", "COMMA", "DOT", "COLON", "ARROW", "INTERROGATION", "ANPERSAND",
	}
	staticData.RuleNames = []string{
		"program", "delimiter", "stmt", "decl_stmt", "vector_expr", "vector_item",
		"repeating", "var_type", "type", "vector_type", "matrix_type", "aux_matrix_type",
		"assign_stmt", "id_pattern", "literal", "expr", "if_stmt", "if_chain",
		"else_stmt", "switch_stmt", "switch_case", "default_case", "while_stmt",
		"for_stmt", "range", "guard_stmt", "transfer_stmt", "func_call", "arg_list",
		"func_arg", "func_dcl", "param_list", "func_param", "strct_dcl", "struct_prop",
		"struct_vector",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 57, 469, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0, 5, 0,
		74, 8, 0, 10, 0, 12, 0, 77, 9, 0, 1, 0, 3, 0, 80, 8, 0, 1, 1, 3, 1, 83,
		8, 1, 1, 1, 3, 1, 86, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 107, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 127, 8, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 5, 4, 133, 8, 4, 10, 4, 12, 4, 136, 9, 4, 3, 4, 138,
		8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 4, 5, 147, 8, 5, 11, 5,
		12, 5, 148, 1, 6, 1, 6, 3, 6, 153, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 3, 8, 170,
		8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		3, 10, 182, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 200,
		8, 12, 1, 13, 1, 13, 1, 13, 5, 13, 205, 8, 13, 10, 13, 12, 13, 208, 9,
		13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 215, 8, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 3, 15, 231, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 5, 15, 251, 8, 15, 10, 15, 12, 15, 254, 9, 15, 1, 16, 1, 16,
		1, 16, 5, 16, 259, 8, 16, 10, 16, 12, 16, 262, 9, 16, 1, 16, 3, 16, 265,
		8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 271, 8, 17, 10, 17, 12, 17, 274,
		9, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 5, 18, 281, 8, 18, 10, 18, 12,
		18, 284, 9, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 292, 8,
		19, 10, 19, 12, 19, 295, 9, 19, 1, 19, 3, 19, 298, 8, 19, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 306, 8, 20, 10, 20, 12, 20, 309, 9,
		20, 1, 21, 1, 21, 1, 21, 5, 21, 314, 8, 21, 10, 21, 12, 21, 317, 9, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 323, 8, 22, 10, 22, 12, 22, 326, 9,
		22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 335, 8, 23,
		1, 23, 1, 23, 5, 23, 339, 8, 23, 10, 23, 12, 23, 342, 9, 23, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 5, 25, 357, 8, 25, 10, 25, 12, 25, 360, 9, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 3, 26, 366, 8, 26, 1, 26, 1, 26, 3, 26, 370, 8, 26, 1, 27, 1,
		27, 1, 27, 3, 27, 375, 8, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 5, 28,
		382, 8, 28, 10, 28, 12, 28, 385, 9, 28, 1, 29, 1, 29, 3, 29, 389, 8, 29,
		1, 29, 3, 29, 392, 8, 29, 1, 29, 1, 29, 3, 29, 396, 8, 29, 1, 30, 1, 30,
		1, 30, 1, 30, 3, 30, 402, 8, 30, 1, 30, 1, 30, 1, 30, 3, 30, 407, 8, 30,
		1, 30, 1, 30, 5, 30, 411, 8, 30, 10, 30, 12, 30, 414, 9, 30, 1, 30, 1,
		30, 1, 31, 1, 31, 1, 31, 5, 31, 421, 8, 31, 10, 31, 12, 31, 424, 9, 31,
		1, 32, 3, 32, 427, 8, 32, 1, 32, 1, 32, 1, 32, 3, 32, 432, 8, 32, 1, 32,
		1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 5, 33, 440, 8, 33, 10, 33, 12, 33, 443,
		9, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 451, 8, 34, 1,
		34, 1, 34, 3, 34, 455, 8, 34, 1, 34, 3, 34, 458, 8, 34, 1, 34, 3, 34, 461,
		8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 0, 1, 30, 36, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 0, 8, 1,
		0, 5, 6, 1, 0, 35, 36, 1, 0, 34, 36, 2, 0, 30, 30, 45, 45, 1, 0, 31, 33,
		1, 0, 29, 30, 1, 0, 39, 42, 1, 0, 37, 38, 507, 0, 75, 1, 0, 0, 0, 2, 85,
		1, 0, 0, 0, 4, 106, 1, 0, 0, 0, 6, 126, 1, 0, 0, 0, 8, 128, 1, 0, 0, 0,
		10, 141, 1, 0, 0, 0, 12, 152, 1, 0, 0, 0, 14, 164, 1, 0, 0, 0, 16, 169,
		1, 0, 0, 0, 18, 171, 1, 0, 0, 0, 20, 181, 1, 0, 0, 0, 22, 183, 1, 0, 0,
		0, 24, 199, 1, 0, 0, 0, 26, 201, 1, 0, 0, 0, 28, 214, 1, 0, 0, 0, 30, 230,
		1, 0, 0, 0, 32, 255, 1, 0, 0, 0, 34, 266, 1, 0, 0, 0, 36, 277, 1, 0, 0,
		0, 38, 287, 1, 0, 0, 0, 40, 301, 1, 0, 0, 0, 42, 310, 1, 0, 0, 0, 44, 318,
		1, 0, 0, 0, 46, 329, 1, 0, 0, 0, 48, 345, 1, 0, 0, 0, 50, 351, 1, 0, 0,
		0, 52, 369, 1, 0, 0, 0, 54, 371, 1, 0, 0, 0, 56, 378, 1, 0, 0, 0, 58, 388,
		1, 0, 0, 0, 60, 397, 1, 0, 0, 0, 62, 417, 1, 0, 0, 0, 64, 426, 1, 0, 0,
		0, 66, 435, 1, 0, 0, 0, 68, 460, 1, 0, 0, 0, 70, 462, 1, 0, 0, 0, 72, 74,
		3, 4, 2, 0, 73, 72, 1, 0, 0, 0, 74, 77, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0,
		75, 76, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 78, 80, 5,
		0, 0, 1, 79, 78, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 1, 1, 0, 0, 0, 81,
		83, 5, 4, 0, 0, 82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 86, 1, 0, 0,
		0, 84, 86, 5, 0, 0, 1, 85, 82, 1, 0, 0, 0, 85, 84, 1, 0, 0, 0, 86, 3, 1,
		0, 0, 0, 87, 88, 3, 6, 3, 0, 88, 89, 3, 2, 1, 0, 89, 107, 1, 0, 0, 0, 90,
		91, 3, 24, 12, 0, 91, 92, 3, 2, 1, 0, 92, 107, 1, 0, 0, 0, 93, 94, 3, 52,
		26, 0, 94, 95, 3, 2, 1, 0, 95, 107, 1, 0, 0, 0, 96, 107, 3, 32, 16, 0,
		97, 107, 3, 38, 19, 0, 98, 107, 3, 44, 22, 0, 99, 107, 3, 46, 23, 0, 100,
		107, 3, 50, 25, 0, 101, 102, 3, 54, 27, 0, 102, 103, 3, 2, 1, 0, 103, 107,
		1, 0, 0, 0, 104, 107, 3, 60, 30, 0, 105, 107, 3, 66, 33, 0, 106, 87, 1,
		0, 0, 0, 106, 90, 1, 0, 0, 0, 106, 93, 1, 0, 0, 0, 106, 96, 1, 0, 0, 0,
		106, 97, 1, 0, 0, 0, 106, 98, 1, 0, 0, 0, 106, 99, 1, 0, 0, 0, 106, 100,
		1, 0, 0, 0, 106, 101, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 105, 1, 0,
		0, 0, 107, 5, 1, 0, 0, 0, 108, 109, 3, 14, 7, 0, 109, 110, 5, 28, 0, 0,
		110, 111, 5, 54, 0, 0, 111, 112, 3, 16, 8, 0, 112, 113, 5, 34, 0, 0, 113,
		114, 3, 30, 15, 0, 114, 127, 1, 0, 0, 0, 115, 116, 3, 14, 7, 0, 116, 117,
		5, 28, 0, 0, 117, 118, 5, 34, 0, 0, 118, 119, 3, 30, 15, 0, 119, 127, 1,
		0, 0, 0, 120, 121, 3, 14, 7, 0, 121, 122, 5, 28, 0, 0, 122, 123, 5, 54,
		0, 0, 123, 124, 3, 16, 8, 0, 124, 125, 5, 56, 0, 0, 125, 127, 1, 0, 0,
		0, 126, 108, 1, 0, 0, 0, 126, 115, 1, 0, 0, 0, 126, 120, 1, 0, 0, 0, 127,
		7, 1, 0, 0, 0, 128, 137, 5, 50, 0, 0, 129, 134, 3, 30, 15, 0, 130, 131,
		5, 52, 0, 0, 131, 133, 3, 30, 15, 0, 132, 130, 1, 0, 0, 0, 133, 136, 1,
		0, 0, 0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 138, 1, 0, 0,
		0, 136, 134, 1, 0, 0, 0, 137, 129, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138,
		139, 1, 0, 0, 0, 139, 140, 5, 51, 0, 0, 140, 9, 1, 0, 0, 0, 141, 146, 3,
		26, 13, 0, 142, 143, 5, 50, 0, 0, 143, 144, 3, 30, 15, 0, 144, 145, 5,
		51, 0, 0, 145, 147, 1, 0, 0, 0, 146, 142, 1, 0, 0, 0, 147, 148, 1, 0, 0,
		0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 11, 1, 0, 0, 0, 150,
		153, 3, 18, 9, 0, 151, 153, 3, 20, 10, 0, 152, 150, 1, 0, 0, 0, 152, 151,
		1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 155, 5, 46, 0, 0, 155, 156, 5, 28,
		0, 0, 156, 157, 5, 54, 0, 0, 157, 158, 3, 30, 15, 0, 158, 159, 5, 52, 0,
		0, 159, 160, 5, 28, 0, 0, 160, 161, 5, 54, 0, 0, 161, 162, 3, 30, 15, 0,
		162, 163, 5, 47, 0, 0, 163, 13, 1, 0, 0, 0, 164, 165, 7, 0, 0, 0, 165,
		15, 1, 0, 0, 0, 166, 170, 5, 28, 0, 0, 167, 170, 3, 18, 9, 0, 168, 170,
		3, 20, 10, 0, 169, 166, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 168, 1,
		0, 0, 0, 170, 17, 1, 0, 0, 0, 171, 172, 5, 50, 0, 0, 172, 173, 5, 28, 0,
		0, 173, 174, 5, 51, 0, 0, 174, 19, 1, 0, 0, 0, 175, 182, 3, 22, 11, 0,
		176, 177, 5, 50, 0, 0, 177, 178, 5, 50, 0, 0, 178, 179, 5, 28, 0, 0, 179,
		180, 5, 51, 0, 0, 180, 182, 5, 51, 0, 0, 181, 175, 1, 0, 0, 0, 181, 176,
		1, 0, 0, 0, 182, 21, 1, 0, 0, 0, 183, 184, 5, 50, 0, 0, 184, 185, 3, 20,
		10, 0, 185, 186, 5, 51, 0, 0, 186, 23, 1, 0, 0, 0, 187, 188, 3, 26, 13,
		0, 188, 189, 5, 34, 0, 0, 189, 190, 3, 30, 15, 0, 190, 200, 1, 0, 0, 0,
		191, 192, 3, 26, 13, 0, 192, 193, 7, 1, 0, 0, 193, 194, 3, 30, 15, 0, 194,
		200, 1, 0, 0, 0, 195, 196, 3, 10, 5, 0, 196, 197, 7, 2, 0, 0, 197, 198,
		3, 30, 15, 0, 198, 200, 1, 0, 0, 0, 199, 187, 1, 0, 0, 0, 199, 191, 1,
		0, 0, 0, 199, 195, 1, 0, 0, 0, 200, 25, 1, 0, 0, 0, 201, 206, 5, 28, 0,
		0, 202, 203, 5, 53, 0, 0, 203, 205, 5, 28, 0, 0, 204, 202, 1, 0, 0, 0,
		205, 208, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207,
		27, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 209, 215, 5, 23, 0, 0, 210, 215,
		5, 24, 0, 0, 211, 215, 5, 25, 0, 0, 212, 215, 5, 26, 0, 0, 213, 215, 5,
		27, 0, 0, 214, 209, 1, 0, 0, 0, 214, 210, 1, 0, 0, 0, 214, 211, 1, 0, 0,
		0, 214, 212, 1, 0, 0, 0, 214, 213, 1, 0, 0, 0, 215, 29, 1, 0, 0, 0, 216,
		217, 6, 15, -1, 0, 217, 218, 5, 46, 0, 0, 218, 219, 3, 30, 15, 0, 219,
		220, 5, 47, 0, 0, 220, 231, 1, 0, 0, 0, 221, 231, 3, 54, 27, 0, 222, 231,
		3, 26, 13, 0, 223, 231, 3, 10, 5, 0, 224, 231, 3, 28, 14, 0, 225, 231,
		3, 8, 4, 0, 226, 231, 3, 12, 6, 0, 227, 231, 3, 70, 35, 0, 228, 229, 7,
		3, 0, 0, 229, 231, 3, 30, 15, 7, 230, 216, 1, 0, 0, 0, 230, 221, 1, 0,
		0, 0, 230, 222, 1, 0, 0, 0, 230, 223, 1, 0, 0, 0, 230, 224, 1, 0, 0, 0,
		230, 225, 1, 0, 0, 0, 230, 226, 1, 0, 0, 0, 230, 227, 1, 0, 0, 0, 230,
		228, 1, 0, 0, 0, 231, 252, 1, 0, 0, 0, 232, 233, 10, 6, 0, 0, 233, 234,
		7, 4, 0, 0, 234, 251, 3, 30, 15, 7, 235, 236, 10, 5, 0, 0, 236, 237, 7,
		5, 0, 0, 237, 251, 3, 30, 15, 6, 238, 239, 10, 4, 0, 0, 239, 240, 7, 6,
		0, 0, 240, 251, 3, 30, 15, 5, 241, 242, 10, 3, 0, 0, 242, 243, 7, 7, 0,
		0, 243, 251, 3, 30, 15, 4, 244, 245, 10, 2, 0, 0, 245, 246, 5, 43, 0, 0,
		246, 251, 3, 30, 15, 3, 247, 248, 10, 1, 0, 0, 248, 249, 5, 44, 0, 0, 249,
		251, 3, 30, 15, 2, 250, 232, 1, 0, 0, 0, 250, 235, 1, 0, 0, 0, 250, 238,
		1, 0, 0, 0, 250, 241, 1, 0, 0, 0, 250, 244, 1, 0, 0, 0, 250, 247, 1, 0,
		0, 0, 251, 254, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0,
		253, 31, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 255, 260, 3, 34, 17, 0, 256,
		257, 5, 10, 0, 0, 257, 259, 3, 34, 17, 0, 258, 256, 1, 0, 0, 0, 259, 262,
		1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 264, 1, 0,
		0, 0, 262, 260, 1, 0, 0, 0, 263, 265, 3, 36, 18, 0, 264, 263, 1, 0, 0,
		0, 264, 265, 1, 0, 0, 0, 265, 33, 1, 0, 0, 0, 266, 267, 5, 9, 0, 0, 267,
		268, 3, 30, 15, 0, 268, 272, 5, 48, 0, 0, 269, 271, 3, 4, 2, 0, 270, 269,
		1, 0, 0, 0, 271, 274, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 272, 273, 1, 0,
		0, 0, 273, 275, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 275, 276, 5, 49, 0, 0,
		276, 35, 1, 0, 0, 0, 277, 278, 5, 10, 0, 0, 278, 282, 5, 48, 0, 0, 279,
		281, 3, 4, 2, 0, 280, 279, 1, 0, 0, 0, 281, 284, 1, 0, 0, 0, 282, 280,
		1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 285, 1, 0, 0, 0, 284, 282, 1, 0,
		0, 0, 285, 286, 5, 49, 0, 0, 286, 37, 1, 0, 0, 0, 287, 288, 5, 11, 0, 0,
		288, 289, 3, 30, 15, 0, 289, 293, 5, 48, 0, 0, 290, 292, 3, 40, 20, 0,
		291, 290, 1, 0, 0, 0, 292, 295, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 293,
		294, 1, 0, 0, 0, 294, 297, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 296, 298,
		3, 42, 21, 0, 297, 296, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 299, 1,
		0, 0, 0, 299, 300, 5, 49, 0, 0, 300, 39, 1, 0, 0, 0, 301, 302, 5, 12, 0,
		0, 302, 303, 3, 30, 15, 0, 303, 307, 5, 54, 0, 0, 304, 306, 3, 4, 2, 0,
		305, 304, 1, 0, 0, 0, 306, 309, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 307,
		308, 1, 0, 0, 0, 308, 41, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 310, 311, 5,
		13, 0, 0, 311, 315, 5, 54, 0, 0, 312, 314, 3, 4, 2, 0, 313, 312, 1, 0,
		0, 0, 314, 317, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0,
		316, 43, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 318, 319, 5, 15, 0, 0, 319,
		320, 3, 30, 15, 0, 320, 324, 5, 48, 0, 0, 321, 323, 3, 4, 2, 0, 322, 321,
		1, 0, 0, 0, 323, 326, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 324, 325, 1, 0,
		0, 0, 325, 327, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 327, 328, 5, 49, 0, 0,
		328, 45, 1, 0, 0, 0, 329, 330, 5, 14, 0, 0, 330, 331, 5, 28, 0, 0, 331,
		334, 5, 22, 0, 0, 332, 335, 3, 30, 15, 0, 333, 335, 3, 48, 24, 0, 334,
		332, 1, 0, 0, 0, 334, 333, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 340,
		5, 48, 0, 0, 337, 339, 3, 4, 2, 0, 338, 337, 1, 0, 0, 0, 339, 342, 1, 0,
		0, 0, 340, 338, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 343, 1, 0, 0, 0,
		342, 340, 1, 0, 0, 0, 343, 344, 5, 49, 0, 0, 344, 47, 1, 0, 0, 0, 345,
		346, 3, 30, 15, 0, 346, 347, 5, 53, 0, 0, 347, 348, 5, 53, 0, 0, 348, 349,
		5, 53, 0, 0, 349, 350, 3, 30, 15, 0, 350, 49, 1, 0, 0, 0, 351, 352, 5,
		19, 0, 0, 352, 353, 3, 30, 15, 0, 353, 354, 5, 10, 0, 0, 354, 358, 5, 48,
		0, 0, 355, 357, 3, 4, 2, 0, 356, 355, 1, 0, 0, 0, 357, 360, 1, 0, 0, 0,
		358, 356, 1, 0, 0, 0, 358, 359, 1, 0, 0, 0, 359, 361, 1, 0, 0, 0, 360,
		358, 1, 0, 0, 0, 361, 362, 5, 49, 0, 0, 362, 51, 1, 0, 0, 0, 363, 365,
		5, 18, 0, 0, 364, 366, 3, 30, 15, 0, 365, 364, 1, 0, 0, 0, 365, 366, 1,
		0, 0, 0, 366, 370, 1, 0, 0, 0, 367, 370, 5, 16, 0, 0, 368, 370, 5, 17,
		0, 0, 369, 363, 1, 0, 0, 0, 369, 367, 1, 0, 0, 0, 369, 368, 1, 0, 0, 0,
		370, 53, 1, 0, 0, 0, 371, 372, 3, 26, 13, 0, 372, 374, 5, 46, 0, 0, 373,
		375, 3, 56, 28, 0, 374, 373, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 376,
		1, 0, 0, 0, 376, 377, 5, 47, 0, 0, 377, 55, 1, 0, 0, 0, 378, 383, 3, 58,
		29, 0, 379, 380, 5, 52, 0, 0, 380, 382, 3, 58, 29, 0, 381, 379, 1, 0, 0,
		0, 382, 385, 1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384,
		57, 1, 0, 0, 0, 385, 383, 1, 0, 0, 0, 386, 387, 5, 28, 0, 0, 387, 389,
		5, 54, 0, 0, 388, 386, 1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389, 391, 1, 0,
		0, 0, 390, 392, 5, 57, 0, 0, 391, 390, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0,
		392, 395, 1, 0, 0, 0, 393, 396, 3, 26, 13, 0, 394, 396, 3, 30, 15, 0, 395,
		393, 1, 0, 0, 0, 395, 394, 1, 0, 0, 0, 396, 59, 1, 0, 0, 0, 397, 398, 5,
		7, 0, 0, 398, 399, 5, 28, 0, 0, 399, 401, 5, 46, 0, 0, 400, 402, 3, 62,
		31, 0, 401, 400, 1, 0, 0, 0, 401, 402, 1, 0, 0, 0, 402, 403, 1, 0, 0, 0,
		403, 406, 5, 47, 0, 0, 404, 405, 5, 55, 0, 0, 405, 407, 3, 16, 8, 0, 406,
		404, 1, 0, 0, 0, 406, 407, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408, 412,
		5, 48, 0, 0, 409, 411, 3, 4, 2, 0, 410, 409, 1, 0, 0, 0, 411, 414, 1, 0,
		0, 0, 412, 410, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 415, 1, 0, 0, 0,
		414, 412, 1, 0, 0, 0, 415, 416, 5, 49, 0, 0, 416, 61, 1, 0, 0, 0, 417,
		422, 3, 64, 32, 0, 418, 419, 5, 52, 0, 0, 419, 421, 3, 64, 32, 0, 420,
		418, 1, 0, 0, 0, 421, 424, 1, 0, 0, 0, 422, 420, 1, 0, 0, 0, 422, 423,
		1, 0, 0, 0, 423, 63, 1, 0, 0, 0, 424, 422, 1, 0, 0, 0, 425, 427, 5, 28,
		0, 0, 426, 425, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 428, 1, 0, 0, 0,
		428, 429, 5, 28, 0, 0, 429, 431, 5, 54, 0, 0, 430, 432, 5, 20, 0, 0, 431,
		430, 1, 0, 0, 0, 431, 432, 1, 0, 0, 0, 432, 433, 1, 0, 0, 0, 433, 434,
		3, 16, 8, 0, 434, 65, 1, 0, 0, 0, 435, 436, 5, 8, 0, 0, 436, 437, 5, 28,
		0, 0, 437, 441, 5, 48, 0, 0, 438, 440, 3, 68, 34, 0, 439, 438, 1, 0, 0,
		0, 440, 443, 1, 0, 0, 0, 441, 439, 1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442,
		444, 1, 0, 0, 0, 443, 441, 1, 0, 0, 0, 444, 445, 5, 49, 0, 0, 445, 67,
		1, 0, 0, 0, 446, 447, 3, 14, 7, 0, 447, 450, 5, 28, 0, 0, 448, 449, 5,
		54, 0, 0, 449, 451, 3, 16, 8, 0, 450, 448, 1, 0, 0, 0, 450, 451, 1, 0,
		0, 0, 451, 454, 1, 0, 0, 0, 452, 453, 5, 34, 0, 0, 453, 455, 3, 30, 15,
		0, 454, 452, 1, 0, 0, 0, 454, 455, 1, 0, 0, 0, 455, 461, 1, 0, 0, 0, 456,
		458, 5, 21, 0, 0, 457, 456, 1, 0, 0, 0, 457, 458, 1, 0, 0, 0, 458, 459,
		1, 0, 0, 0, 459, 461, 3, 60, 30, 0, 460, 446, 1, 0, 0, 0, 460, 457, 1,
		0, 0, 0, 461, 69, 1, 0, 0, 0, 462, 463, 5, 50, 0, 0, 463, 464, 5, 28, 0,
		0, 464, 465, 5, 51, 0, 0, 465, 466, 5, 46, 0, 0, 466, 467, 5, 47, 0, 0,
		467, 71, 1, 0, 0, 0, 48, 75, 79, 82, 85, 106, 126, 134, 137, 148, 152,
		169, 181, 199, 206, 214, 230, 250, 252, 260, 264, 272, 282, 293, 297, 307,
		315, 324, 334, 340, 358, 365, 369, 374, 383, 388, 391, 395, 401, 406, 412,
		422, 426, 431, 441, 450, 454, 457, 460,
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
	TSwiftLanguageLET_KW                = 5
	TSwiftLanguageVAR_KW                = 6
	TSwiftLanguageFUNC_KW               = 7
	TSwiftLanguageSTRUCT_KW             = 8
	TSwiftLanguageIF_KW                 = 9
	TSwiftLanguageELSE_KW               = 10
	TSwiftLanguageSWITCH_KW             = 11
	TSwiftLanguageCASE_KW               = 12
	TSwiftLanguageDEFAULT_KW            = 13
	TSwiftLanguageFOR_KW                = 14
	TSwiftLanguageWHILE_KW              = 15
	TSwiftLanguageBREAK_KW              = 16
	TSwiftLanguageCONTINUE_KW           = 17
	TSwiftLanguageRETURN_KW             = 18
	TSwiftLanguageGUARD_KW              = 19
	TSwiftLanguageINOUT_KW              = 20
	TSwiftLanguageMUTATING_KW           = 21
	TSwiftLanguageIN_KW                 = 22
	TSwiftLanguageINTEGER_LITERAL       = 23
	TSwiftLanguageFLOAT_LITERAL         = 24
	TSwiftLanguageSTRING_LITERAL        = 25
	TSwiftLanguageBOOL_LITERAL          = 26
	TSwiftLanguageNIL_LITERAL           = 27
	TSwiftLanguageID                    = 28
	TSwiftLanguagePLUS                  = 29
	TSwiftLanguageMINUS                 = 30
	TSwiftLanguageMULT                  = 31
	TSwiftLanguageDIV                   = 32
	TSwiftLanguageMOD                   = 33
	TSwiftLanguageEQUALS                = 34
	TSwiftLanguagePLUS_EQUALS           = 35
	TSwiftLanguageMINUS_EQUALS          = 36
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
	TSwiftLanguageANPERSAND             = 57
)

// TSwiftLanguage rules.
const (
	TSwiftLanguageRULE_program         = 0
	TSwiftLanguageRULE_delimiter       = 1
	TSwiftLanguageRULE_stmt            = 2
	TSwiftLanguageRULE_decl_stmt       = 3
	TSwiftLanguageRULE_vector_expr     = 4
	TSwiftLanguageRULE_vector_item     = 5
	TSwiftLanguageRULE_repeating       = 6
	TSwiftLanguageRULE_var_type        = 7
	TSwiftLanguageRULE_type            = 8
	TSwiftLanguageRULE_vector_type     = 9
	TSwiftLanguageRULE_matrix_type     = 10
	TSwiftLanguageRULE_aux_matrix_type = 11
	TSwiftLanguageRULE_assign_stmt     = 12
	TSwiftLanguageRULE_id_pattern      = 13
	TSwiftLanguageRULE_literal         = 14
	TSwiftLanguageRULE_expr            = 15
	TSwiftLanguageRULE_if_stmt         = 16
	TSwiftLanguageRULE_if_chain        = 17
	TSwiftLanguageRULE_else_stmt       = 18
	TSwiftLanguageRULE_switch_stmt     = 19
	TSwiftLanguageRULE_switch_case     = 20
	TSwiftLanguageRULE_default_case    = 21
	TSwiftLanguageRULE_while_stmt      = 22
	TSwiftLanguageRULE_for_stmt        = 23
	TSwiftLanguageRULE_range           = 24
	TSwiftLanguageRULE_guard_stmt      = 25
	TSwiftLanguageRULE_transfer_stmt   = 26
	TSwiftLanguageRULE_func_call       = 27
	TSwiftLanguageRULE_arg_list        = 28
	TSwiftLanguageRULE_func_arg        = 29
	TSwiftLanguageRULE_func_dcl        = 30
	TSwiftLanguageRULE_param_list      = 31
	TSwiftLanguageRULE_func_param      = 32
	TSwiftLanguageRULE_strct_dcl       = 33
	TSwiftLanguageRULE_struct_prop     = 34
	TSwiftLanguageRULE_struct_vector   = 35
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	EOF() antlr.TerminalNode

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

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageEOF, 0)
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
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&269470688) != 0 {
		{
			p.SetState(72)
			p.Stmt()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(78)
			p.Match(TSwiftLanguageEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
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

// IDelimiterContext is an interface to support dynamic dispatch.
type IDelimiterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEMICOLON() antlr.TerminalNode
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

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TSwiftLanguageSEMICOLON {
			{
				p.SetState(81)
				p.Match(TSwiftLanguageSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.Match(TSwiftLanguageEOF)
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

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Decl_stmt() IDecl_stmtContext
	Delimiter() IDelimiterContext
	Assign_stmt() IAssign_stmtContext
	Transfer_stmt() ITransfer_stmtContext
	If_stmt() IIf_stmtContext
	Switch_stmt() ISwitch_stmtContext
	While_stmt() IWhile_stmtContext
	For_stmt() IFor_stmtContext
	Guard_stmt() IGuard_stmtContext
	Func_call() IFunc_callContext
	Func_dcl() IFunc_dclContext
	Strct_dcl() IStrct_dclContext

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

func (s *StmtContext) Delimiter() IDelimiterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDelimiterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDelimiterContext)
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

func (s *StmtContext) Transfer_stmt() ITransfer_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransfer_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransfer_stmtContext)
}

func (s *StmtContext) If_stmt() IIf_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *StmtContext) Switch_stmt() ISwitch_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_stmtContext)
}

func (s *StmtContext) While_stmt() IWhile_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_stmtContext)
}

func (s *StmtContext) For_stmt() IFor_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_stmtContext)
}

func (s *StmtContext) Guard_stmt() IGuard_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuard_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuard_stmtContext)
}

func (s *StmtContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *StmtContext) Func_dcl() IFunc_dclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_dclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_dclContext)
}

func (s *StmtContext) Strct_dcl() IStrct_dclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrct_dclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrct_dclContext)
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
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Decl_stmt()
		}
		{
			p.SetState(88)
			p.Delimiter()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.Assign_stmt()
		}
		{
			p.SetState(91)
			p.Delimiter()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
			p.Transfer_stmt()
		}
		{
			p.SetState(94)
			p.Delimiter()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(96)
			p.If_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(97)
			p.Switch_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(98)
			p.While_stmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(99)
			p.For_stmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(100)
			p.Guard_stmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(101)
			p.Func_call()
		}
		{
			p.SetState(102)
			p.Delimiter()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(104)
			p.Func_dcl()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(105)
			p.Strct_dcl()
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

func (s *TypeValueDeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
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

func (s *TypeDeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
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
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeValueDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Var_type()
		}
		{
			p.SetState(109)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.Match(TSwiftLanguageCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Type_()
		}
		{
			p.SetState(112)
			p.Match(TSwiftLanguageEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.expr(0)
		}

	case 2:
		localctx = NewValueDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Var_type()
		}
		{
			p.SetState(116)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(TSwiftLanguageEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.expr(0)
		}

	case 3:
		localctx = NewTypeDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(120)
			p.Var_type()
		}
		{
			p.SetState(121)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Match(TSwiftLanguageCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Type_()
		}
		{
			p.SetState(124)
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

// IVector_exprContext is an interface to support dynamic dispatch.
type IVector_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVector_exprContext differentiates from other interfaces.
	IsVector_exprContext()
}

type Vector_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVector_exprContext() *Vector_exprContext {
	var p = new(Vector_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_vector_expr
	return p
}

func InitEmptyVector_exprContext(p *Vector_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_vector_expr
}

func (*Vector_exprContext) IsVector_exprContext() {}

func NewVector_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vector_exprContext {
	var p = new(Vector_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_vector_expr

	return p
}

func (s *Vector_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Vector_exprContext) CopyAll(ctx *Vector_exprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Vector_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vector_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VectorItemListContext struct {
	Vector_exprContext
}

func NewVectorItemListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorItemListContext {
	var p = new(VectorItemListContext)

	InitEmptyVector_exprContext(&p.Vector_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Vector_exprContext))

	return p
}

func (s *VectorItemListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorItemListContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLBRACK, 0)
}

func (s *VectorItemListContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRBRACK, 0)
}

func (s *VectorItemListContext) AllExpr() []IExprContext {
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

func (s *VectorItemListContext) Expr(i int) IExprContext {
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

func (s *VectorItemListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TSwiftLanguageCOMMA)
}

func (s *VectorItemListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageCOMMA, i)
}

func (s *VectorItemListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitVectorItemList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Vector_expr() (localctx IVector_exprContext) {
	localctx = NewVector_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TSwiftLanguageRULE_vector_expr)
	var _la int

	localctx = NewVectorItemListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(TSwiftLanguageLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1231454625333248) != 0 {
		{
			p.SetState(129)
			p.expr(0)
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TSwiftLanguageCOMMA {
			{
				p.SetState(130)
				p.Match(TSwiftLanguageCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(131)
				p.expr(0)
			}

			p.SetState(136)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(139)
		p.Match(TSwiftLanguageRBRACK)
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

// IVector_itemContext is an interface to support dynamic dispatch.
type IVector_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVector_itemContext differentiates from other interfaces.
	IsVector_itemContext()
}

type Vector_itemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVector_itemContext() *Vector_itemContext {
	var p = new(Vector_itemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_vector_item
	return p
}

func InitEmptyVector_itemContext(p *Vector_itemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_vector_item
}

func (*Vector_itemContext) IsVector_itemContext() {}

func NewVector_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vector_itemContext {
	var p = new(Vector_itemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_vector_item

	return p
}

func (s *Vector_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Vector_itemContext) CopyAll(ctx *Vector_itemContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Vector_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vector_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VectorItemContext struct {
	Vector_itemContext
}

func NewVectorItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorItemContext {
	var p = new(VectorItemContext)

	InitEmptyVector_itemContext(&p.Vector_itemContext)
	p.parser = parser
	p.CopyAll(ctx.(*Vector_itemContext))

	return p
}

func (s *VectorItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorItemContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *VectorItemContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(TSwiftLanguageLBRACK)
}

func (s *VectorItemContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLBRACK, i)
}

func (s *VectorItemContext) AllExpr() []IExprContext {
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

func (s *VectorItemContext) Expr(i int) IExprContext {
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

func (s *VectorItemContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(TSwiftLanguageRBRACK)
}

func (s *VectorItemContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRBRACK, i)
}

func (s *VectorItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitVectorItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Vector_item() (localctx IVector_itemContext) {
	localctx = NewVector_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TSwiftLanguageRULE_vector_item)
	var _alt int

	localctx = NewVectorItemContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Id_pattern()
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(142)
				p.Match(TSwiftLanguageLBRACK)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(143)
				p.expr(0)
			}
			{
				p.SetState(144)
				p.Match(TSwiftLanguageRBRACK)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRepeatingContext is an interface to support dynamic dispatch.
type IRepeatingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	COMMA() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Vector_type() IVector_typeContext
	Matrix_type() IMatrix_typeContext

	// IsRepeatingContext differentiates from other interfaces.
	IsRepeatingContext()
}

type RepeatingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepeatingContext() *RepeatingContext {
	var p = new(RepeatingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_repeating
	return p
}

func InitEmptyRepeatingContext(p *RepeatingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_repeating
}

func (*RepeatingContext) IsRepeatingContext() {}

func NewRepeatingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepeatingContext {
	var p = new(RepeatingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_repeating

	return p
}

func (s *RepeatingContext) GetParser() antlr.Parser { return s.parser }

func (s *RepeatingContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLPAREN, 0)
}

func (s *RepeatingContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TSwiftLanguageID)
}

func (s *RepeatingContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, i)
}

func (s *RepeatingContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(TSwiftLanguageCOLON)
}

func (s *RepeatingContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageCOLON, i)
}

func (s *RepeatingContext) AllExpr() []IExprContext {
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

func (s *RepeatingContext) Expr(i int) IExprContext {
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

func (s *RepeatingContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageCOMMA, 0)
}

func (s *RepeatingContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRPAREN, 0)
}

func (s *RepeatingContext) Vector_type() IVector_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_typeContext)
}

func (s *RepeatingContext) Matrix_type() IMatrix_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_typeContext)
}

func (s *RepeatingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepeatingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitRepeating(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Repeating() (localctx IRepeatingContext) {
	localctx = NewRepeatingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TSwiftLanguageRULE_repeating)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(150)
			p.Vector_type()
		}

	case 2:
		{
			p.SetState(151)
			p.Matrix_type()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(154)
		p.Match(TSwiftLanguageLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(TSwiftLanguageCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.expr(0)
	}
	{
		p.SetState(158)
		p.Match(TSwiftLanguageCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.Match(TSwiftLanguageCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.expr(0)
	}
	{
		p.SetState(162)
		p.Match(TSwiftLanguageRPAREN)
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
	p.EnterRule(localctx, 14, TSwiftLanguageRULE_var_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Vector_type() IVector_typeContext
	Matrix_type() IMatrix_typeContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) ID() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, 0)
}

func (s *TypeContext) Vector_type() IVector_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_typeContext)
}

func (s *TypeContext) Matrix_type() IMatrix_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_typeContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TSwiftLanguageRULE_type)
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Vector_type()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(168)
			p.Matrix_type()
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

// IVector_typeContext is an interface to support dynamic dispatch.
type IVector_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	ID() antlr.TerminalNode
	RBRACK() antlr.TerminalNode

	// IsVector_typeContext differentiates from other interfaces.
	IsVector_typeContext()
}

type Vector_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVector_typeContext() *Vector_typeContext {
	var p = new(Vector_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_vector_type
	return p
}

func InitEmptyVector_typeContext(p *Vector_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_vector_type
}

func (*Vector_typeContext) IsVector_typeContext() {}

func NewVector_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vector_typeContext {
	var p = new(Vector_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_vector_type

	return p
}

func (s *Vector_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Vector_typeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLBRACK, 0)
}

func (s *Vector_typeContext) ID() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, 0)
}

func (s *Vector_typeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRBRACK, 0)
}

func (s *Vector_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vector_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vector_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitVector_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Vector_type() (localctx IVector_typeContext) {
	localctx = NewVector_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TSwiftLanguageRULE_vector_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(TSwiftLanguageLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Match(TSwiftLanguageRBRACK)
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

// IMatrix_typeContext is an interface to support dynamic dispatch.
type IMatrix_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Aux_matrix_type() IAux_matrix_typeContext
	AllLBRACK() []antlr.TerminalNode
	LBRACK(i int) antlr.TerminalNode
	ID() antlr.TerminalNode
	AllRBRACK() []antlr.TerminalNode
	RBRACK(i int) antlr.TerminalNode

	// IsMatrix_typeContext differentiates from other interfaces.
	IsMatrix_typeContext()
}

type Matrix_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrix_typeContext() *Matrix_typeContext {
	var p = new(Matrix_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_matrix_type
	return p
}

func InitEmptyMatrix_typeContext(p *Matrix_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_matrix_type
}

func (*Matrix_typeContext) IsMatrix_typeContext() {}

func NewMatrix_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Matrix_typeContext {
	var p = new(Matrix_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_matrix_type

	return p
}

func (s *Matrix_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Matrix_typeContext) Aux_matrix_type() IAux_matrix_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAux_matrix_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAux_matrix_typeContext)
}

func (s *Matrix_typeContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(TSwiftLanguageLBRACK)
}

func (s *Matrix_typeContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLBRACK, i)
}

func (s *Matrix_typeContext) ID() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, 0)
}

func (s *Matrix_typeContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(TSwiftLanguageRBRACK)
}

func (s *Matrix_typeContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRBRACK, i)
}

func (s *Matrix_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Matrix_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Matrix_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitMatrix_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Matrix_type() (localctx IMatrix_typeContext) {
	localctx = NewMatrix_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TSwiftLanguageRULE_matrix_type)
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)
			p.Aux_matrix_type()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)
			p.Match(TSwiftLanguageLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Match(TSwiftLanguageLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.Match(TSwiftLanguageRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.Match(TSwiftLanguageRBRACK)
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

// IAux_matrix_typeContext is an interface to support dynamic dispatch.
type IAux_matrix_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	Matrix_type() IMatrix_typeContext
	RBRACK() antlr.TerminalNode

	// IsAux_matrix_typeContext differentiates from other interfaces.
	IsAux_matrix_typeContext()
}

type Aux_matrix_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAux_matrix_typeContext() *Aux_matrix_typeContext {
	var p = new(Aux_matrix_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_aux_matrix_type
	return p
}

func InitEmptyAux_matrix_typeContext(p *Aux_matrix_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_aux_matrix_type
}

func (*Aux_matrix_typeContext) IsAux_matrix_typeContext() {}

func NewAux_matrix_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Aux_matrix_typeContext {
	var p = new(Aux_matrix_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_aux_matrix_type

	return p
}

func (s *Aux_matrix_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Aux_matrix_typeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLBRACK, 0)
}

func (s *Aux_matrix_typeContext) Matrix_type() IMatrix_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_typeContext)
}

func (s *Aux_matrix_typeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRBRACK, 0)
}

func (s *Aux_matrix_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Aux_matrix_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Aux_matrix_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitAux_matrix_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Aux_matrix_type() (localctx IAux_matrix_typeContext) {
	localctx = NewAux_matrix_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TSwiftLanguageRULE_aux_matrix_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(TSwiftLanguageLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.Matrix_type()
	}
	{
		p.SetState(185)
		p.Match(TSwiftLanguageRBRACK)
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

type DirectAssignContext struct {
	Assign_stmtContext
}

func NewDirectAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DirectAssignContext {
	var p = new(DirectAssignContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *DirectAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectAssignContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *DirectAssignContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageEQUALS, 0)
}

func (s *DirectAssignContext) Expr() IExprContext {
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

func (s *DirectAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitDirectAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorAssignContext struct {
	Assign_stmtContext
	op antlr.Token
}

func NewVectorAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorAssignContext {
	var p = new(VectorAssignContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *VectorAssignContext) GetOp() antlr.Token { return s.op }

func (s *VectorAssignContext) SetOp(v antlr.Token) { s.op = v }

func (s *VectorAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorAssignContext) Vector_item() IVector_itemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_itemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_itemContext)
}

func (s *VectorAssignContext) Expr() IExprContext {
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

func (s *VectorAssignContext) PLUS_EQUALS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguagePLUS_EQUALS, 0)
}

func (s *VectorAssignContext) MINUS_EQUALS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageMINUS_EQUALS, 0)
}

func (s *VectorAssignContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageEQUALS, 0)
}

func (s *VectorAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitVectorAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArithmeticAssignContext struct {
	Assign_stmtContext
	op antlr.Token
}

func NewArithmeticAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeticAssignContext {
	var p = new(ArithmeticAssignContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *ArithmeticAssignContext) GetOp() antlr.Token { return s.op }

func (s *ArithmeticAssignContext) SetOp(v antlr.Token) { s.op = v }

func (s *ArithmeticAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticAssignContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *ArithmeticAssignContext) Expr() IExprContext {
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

func (s *ArithmeticAssignContext) PLUS_EQUALS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguagePLUS_EQUALS, 0)
}

func (s *ArithmeticAssignContext) MINUS_EQUALS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageMINUS_EQUALS, 0)
}

func (s *ArithmeticAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitArithmeticAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Assign_stmt() (localctx IAssign_stmtContext) {
	localctx = NewAssign_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TSwiftLanguageRULE_assign_stmt)
	var _la int

	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDirectAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.Id_pattern()
		}
		{
			p.SetState(188)
			p.Match(TSwiftLanguageEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.expr(0)
		}

	case 2:
		localctx = NewArithmeticAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)
			p.Id_pattern()
		}
		{
			p.SetState(192)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ArithmeticAssignContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == TSwiftLanguagePLUS_EQUALS || _la == TSwiftLanguageMINUS_EQUALS) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ArithmeticAssignContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(193)
			p.expr(0)
		}

	case 3:
		localctx = NewVectorAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(195)
			p.Vector_item()
		}
		{
			p.SetState(196)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*VectorAssignContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&120259084288) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*VectorAssignContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(197)
			p.expr(0)
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

// IId_patternContext is an interface to support dynamic dispatch.
type IId_patternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsId_patternContext differentiates from other interfaces.
	IsId_patternContext()
}

type Id_patternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_patternContext() *Id_patternContext {
	var p = new(Id_patternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_id_pattern
	return p
}

func InitEmptyId_patternContext(p *Id_patternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_id_pattern
}

func (*Id_patternContext) IsId_patternContext() {}

func NewId_patternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_patternContext {
	var p = new(Id_patternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_id_pattern

	return p
}

func (s *Id_patternContext) GetParser() antlr.Parser { return s.parser }

func (s *Id_patternContext) CopyAll(ctx *Id_patternContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Id_patternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_patternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdPatternContext struct {
	Id_patternContext
}

func NewIdPatternContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdPatternContext {
	var p = new(IdPatternContext)

	InitEmptyId_patternContext(&p.Id_patternContext)
	p.parser = parser
	p.CopyAll(ctx.(*Id_patternContext))

	return p
}

func (s *IdPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdPatternContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TSwiftLanguageID)
}

func (s *IdPatternContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, i)
}

func (s *IdPatternContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TSwiftLanguageDOT)
}

func (s *IdPatternContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageDOT, i)
}

func (s *IdPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitIdPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Id_pattern() (localctx IId_patternContext) {
	localctx = NewId_patternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TSwiftLanguageRULE_id_pattern)
	var _alt int

	localctx = NewIdPatternContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(202)
				p.Match(TSwiftLanguageDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(203)
				p.Match(TSwiftLanguageID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 28, TSwiftLanguageRULE_literal)
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSwiftLanguageINTEGER_LITERAL:
		localctx = NewIntLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
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
			p.SetState(210)
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
			p.SetState(211)
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
			p.SetState(212)
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
			p.SetState(213)
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

func (s *IdExpContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *IdExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitIdExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructVectorExpContext struct {
	ExprContext
}

func NewStructVectorExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructVectorExpContext {
	var p = new(StructVectorExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StructVectorExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructVectorExpContext) Struct_vector() IStruct_vectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_vectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_vectorContext)
}

func (s *StructVectorExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitStructVectorExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type RepeatingExpContext struct {
	ExprContext
}

func NewRepeatingExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RepeatingExpContext {
	var p = new(RepeatingExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RepeatingExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatingExpContext) Repeating() IRepeatingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepeatingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepeatingContext)
}

func (s *RepeatingExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitRepeatingExp(s)

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

type VectorItemExpContext struct {
	ExprContext
}

func NewVectorItemExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorItemExpContext {
	var p = new(VectorItemExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VectorItemExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorItemExpContext) Vector_item() IVector_itemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_itemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_itemContext)
}

func (s *VectorItemExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitVectorItemExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncCallExpContext struct {
	ExprContext
}

func NewFuncCallExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallExpContext {
	var p = new(FuncCallExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FuncCallExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallExpContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *FuncCallExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitFuncCallExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorExpContext struct {
	ExprContext
}

func NewVectorExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorExpContext {
	var p = new(VectorExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VectorExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorExpContext) Vector_expr() IVector_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_exprContext)
}

func (s *VectorExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitVectorExp(s)

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

func (s *BinaryExpContext) MULT() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageMULT, 0)
}

func (s *BinaryExpContext) DIV() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageDIV, 0)
}

func (s *BinaryExpContext) MOD() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageMOD, 0)
}

func (s *BinaryExpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguagePLUS, 0)
}

func (s *BinaryExpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageMINUS, 0)
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

func (s *BinaryExpContext) EQUALS_EQUALS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageEQUALS_EQUALS, 0)
}

func (s *BinaryExpContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageNOT_EQUALS, 0)
}

func (s *BinaryExpContext) AND() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageAND, 0)
}

func (s *BinaryExpContext) OR() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageOR, 0)
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
	_startState := 30
	p.EnterRecursionRule(localctx, 30, TSwiftLanguageRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(217)
			p.Match(TSwiftLanguageLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.expr(0)
		}
		{
			p.SetState(219)
			p.Match(TSwiftLanguageRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFuncCallExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(221)
			p.Func_call()
		}

	case 3:
		localctx = NewIdExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(222)
			p.Id_pattern()
		}

	case 4:
		localctx = NewVectorItemExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(223)
			p.Vector_item()
		}

	case 5:
		localctx = NewLiteralExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(224)
			p.Literal()
		}

	case 6:
		localctx = NewVectorExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(225)
			p.Vector_expr()
		}

	case 7:
		localctx = NewRepeatingExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(226)
			p.Repeating()
		}

	case 8:
		localctx = NewStructVectorExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(227)
			p.Struct_vector()
		}

	case 9:
		localctx = NewUnaryExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(228)

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
			p.SetState(229)
			p.expr(7)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(250)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TSwiftLanguageRULE_expr)
				p.SetState(232)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(233)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15032385536) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(234)

					var _x = p.expr(7)

					localctx.(*BinaryExpContext).right = _x
				}

			case 2:
				localctx = NewBinaryExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TSwiftLanguageRULE_expr)
				p.SetState(235)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(236)

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
					p.SetState(237)

					var _x = p.expr(6)

					localctx.(*BinaryExpContext).right = _x
				}

			case 3:
				localctx = NewBinaryExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TSwiftLanguageRULE_expr)
				p.SetState(238)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(239)

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
					p.SetState(240)

					var _x = p.expr(5)

					localctx.(*BinaryExpContext).right = _x
				}

			case 4:
				localctx = NewBinaryExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TSwiftLanguageRULE_expr)
				p.SetState(241)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(242)

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
					p.SetState(243)

					var _x = p.expr(4)

					localctx.(*BinaryExpContext).right = _x
				}

			case 5:
				localctx = NewBinaryExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TSwiftLanguageRULE_expr)
				p.SetState(244)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(245)

					var _m = p.Match(TSwiftLanguageAND)

					localctx.(*BinaryExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(246)

					var _x = p.expr(3)

					localctx.(*BinaryExpContext).right = _x
				}

			case 6:
				localctx = NewBinaryExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TSwiftLanguageRULE_expr)
				p.SetState(247)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(248)

					var _m = p.Match(TSwiftLanguageOR)

					localctx.(*BinaryExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(249)

					var _x = p.expr(2)

					localctx.(*BinaryExpContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
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

func (s *If_stmtContext) CopyAll(ctx *If_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfStmtContext struct {
	If_stmtContext
}

func NewIfStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStmtContext {
	var p = new(IfStmtContext)

	InitEmptyIf_stmtContext(&p.If_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_stmtContext))

	return p
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) AllIf_chain() []IIf_chainContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIf_chainContext); ok {
			len++
		}
	}

	tst := make([]IIf_chainContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIf_chainContext); ok {
			tst[i] = t.(IIf_chainContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) If_chain(i int) IIf_chainContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_chainContext); ok {
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

	return t.(IIf_chainContext)
}

func (s *IfStmtContext) AllELSE_KW() []antlr.TerminalNode {
	return s.GetTokens(TSwiftLanguageELSE_KW)
}

func (s *IfStmtContext) ELSE_KW(i int) antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageELSE_KW, i)
}

func (s *IfStmtContext) Else_stmt() IElse_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_stmtContext)
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TSwiftLanguageRULE_if_stmt)
	var _la int

	var _alt int

	localctx = NewIfStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.If_chain()
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(256)
				p.Match(TSwiftLanguageELSE_KW)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(257)
				p.If_chain()
			}

		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSwiftLanguageELSE_KW {
		{
			p.SetState(263)
			p.Else_stmt()
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

// IIf_chainContext is an interface to support dynamic dispatch.
type IIf_chainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_chainContext differentiates from other interfaces.
	IsIf_chainContext()
}

type If_chainContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_chainContext() *If_chainContext {
	var p = new(If_chainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_if_chain
	return p
}

func InitEmptyIf_chainContext(p *If_chainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_if_chain
}

func (*If_chainContext) IsIf_chainContext() {}

func NewIf_chainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_chainContext {
	var p = new(If_chainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_if_chain

	return p
}

func (s *If_chainContext) GetParser() antlr.Parser { return s.parser }

func (s *If_chainContext) CopyAll(ctx *If_chainContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_chainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_chainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfChainContext struct {
	If_chainContext
}

func NewIfChainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfChainContext {
	var p = new(IfChainContext)

	InitEmptyIf_chainContext(&p.If_chainContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_chainContext))

	return p
}

func (s *IfChainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfChainContext) IF_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageIF_KW, 0)
}

func (s *IfChainContext) Expr() IExprContext {
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

func (s *IfChainContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLBRACE, 0)
}

func (s *IfChainContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRBRACE, 0)
}

func (s *IfChainContext) AllStmt() []IStmtContext {
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

func (s *IfChainContext) Stmt(i int) IStmtContext {
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

func (s *IfChainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitIfChain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) If_chain() (localctx IIf_chainContext) {
	localctx = NewIf_chainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TSwiftLanguageRULE_if_chain)
	var _la int

	localctx = NewIfChainContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(TSwiftLanguageIF_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.expr(0)
	}
	{
		p.SetState(268)
		p.Match(TSwiftLanguageLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&269470688) != 0 {
		{
			p.SetState(269)
			p.Stmt()
		}

		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(275)
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

// IElse_stmtContext is an interface to support dynamic dispatch.
type IElse_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsElse_stmtContext differentiates from other interfaces.
	IsElse_stmtContext()
}

type Else_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_stmtContext() *Else_stmtContext {
	var p = new(Else_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_else_stmt
	return p
}

func InitEmptyElse_stmtContext(p *Else_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_else_stmt
}

func (*Else_stmtContext) IsElse_stmtContext() {}

func NewElse_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_stmtContext {
	var p = new(Else_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_else_stmt

	return p
}

func (s *Else_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_stmtContext) CopyAll(ctx *Else_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Else_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ElseStmtContext struct {
	Else_stmtContext
}

func NewElseStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseStmtContext {
	var p = new(ElseStmtContext)

	InitEmptyElse_stmtContext(&p.Else_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Else_stmtContext))

	return p
}

func (s *ElseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStmtContext) ELSE_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageELSE_KW, 0)
}

func (s *ElseStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLBRACE, 0)
}

func (s *ElseStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRBRACE, 0)
}

func (s *ElseStmtContext) AllStmt() []IStmtContext {
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

func (s *ElseStmtContext) Stmt(i int) IStmtContext {
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

func (s *ElseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitElseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Else_stmt() (localctx IElse_stmtContext) {
	localctx = NewElse_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TSwiftLanguageRULE_else_stmt)
	var _la int

	localctx = NewElseStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Match(TSwiftLanguageELSE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.Match(TSwiftLanguageLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&269470688) != 0 {
		{
			p.SetState(279)
			p.Stmt()
		}

		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(285)
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

// ISwitch_stmtContext is an interface to support dynamic dispatch.
type ISwitch_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitch_stmtContext differentiates from other interfaces.
	IsSwitch_stmtContext()
}

type Switch_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_stmtContext() *Switch_stmtContext {
	var p = new(Switch_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_switch_stmt
	return p
}

func InitEmptySwitch_stmtContext(p *Switch_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_switch_stmt
}

func (*Switch_stmtContext) IsSwitch_stmtContext() {}

func NewSwitch_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_stmtContext {
	var p = new(Switch_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_switch_stmt

	return p
}

func (s *Switch_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_stmtContext) CopyAll(ctx *Switch_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Switch_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchStmtContext struct {
	Switch_stmtContext
}

func NewSwitchStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	InitEmptySwitch_stmtContext(&p.Switch_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Switch_stmtContext))

	return p
}

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtContext) SWITCH_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageSWITCH_KW, 0)
}

func (s *SwitchStmtContext) Expr() IExprContext {
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

func (s *SwitchStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLBRACE, 0)
}

func (s *SwitchStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRBRACE, 0)
}

func (s *SwitchStmtContext) AllSwitch_case() []ISwitch_caseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitch_caseContext); ok {
			len++
		}
	}

	tst := make([]ISwitch_caseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitch_caseContext); ok {
			tst[i] = t.(ISwitch_caseContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStmtContext) Switch_case(i int) ISwitch_caseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_caseContext); ok {
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

	return t.(ISwitch_caseContext)
}

func (s *SwitchStmtContext) Default_case() IDefault_caseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefault_caseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefault_caseContext)
}

func (s *SwitchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitSwitchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Switch_stmt() (localctx ISwitch_stmtContext) {
	localctx = NewSwitch_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, TSwiftLanguageRULE_switch_stmt)
	var _la int

	localctx = NewSwitchStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		p.Match(TSwiftLanguageSWITCH_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.expr(0)
	}
	{
		p.SetState(289)
		p.Match(TSwiftLanguageLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TSwiftLanguageCASE_KW {
		{
			p.SetState(290)
			p.Switch_case()
		}

		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSwiftLanguageDEFAULT_KW {
		{
			p.SetState(296)
			p.Default_case()
		}

	}
	{
		p.SetState(299)
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

// ISwitch_caseContext is an interface to support dynamic dispatch.
type ISwitch_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitch_caseContext differentiates from other interfaces.
	IsSwitch_caseContext()
}

type Switch_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_caseContext() *Switch_caseContext {
	var p = new(Switch_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_switch_case
	return p
}

func InitEmptySwitch_caseContext(p *Switch_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_switch_case
}

func (*Switch_caseContext) IsSwitch_caseContext() {}

func NewSwitch_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_caseContext {
	var p = new(Switch_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_switch_case

	return p
}

func (s *Switch_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_caseContext) CopyAll(ctx *Switch_caseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Switch_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchCaseContext struct {
	Switch_caseContext
}

func NewSwitchCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchCaseContext {
	var p = new(SwitchCaseContext)

	InitEmptySwitch_caseContext(&p.Switch_caseContext)
	p.parser = parser
	p.CopyAll(ctx.(*Switch_caseContext))

	return p
}

func (s *SwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCaseContext) CASE_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageCASE_KW, 0)
}

func (s *SwitchCaseContext) Expr() IExprContext {
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

func (s *SwitchCaseContext) COLON() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageCOLON, 0)
}

func (s *SwitchCaseContext) AllStmt() []IStmtContext {
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

func (s *SwitchCaseContext) Stmt(i int) IStmtContext {
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

func (s *SwitchCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitSwitchCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Switch_case() (localctx ISwitch_caseContext) {
	localctx = NewSwitch_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, TSwiftLanguageRULE_switch_case)
	var _la int

	localctx = NewSwitchCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(TSwiftLanguageCASE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(302)
		p.expr(0)
	}
	{
		p.SetState(303)
		p.Match(TSwiftLanguageCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&269470688) != 0 {
		{
			p.SetState(304)
			p.Stmt()
		}

		p.SetState(309)
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

// IDefault_caseContext is an interface to support dynamic dispatch.
type IDefault_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDefault_caseContext differentiates from other interfaces.
	IsDefault_caseContext()
}

type Default_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefault_caseContext() *Default_caseContext {
	var p = new(Default_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_default_case
	return p
}

func InitEmptyDefault_caseContext(p *Default_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_default_case
}

func (*Default_caseContext) IsDefault_caseContext() {}

func NewDefault_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Default_caseContext {
	var p = new(Default_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_default_case

	return p
}

func (s *Default_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Default_caseContext) CopyAll(ctx *Default_caseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Default_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Default_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DefaultCaseContext struct {
	Default_caseContext
}

func NewDefaultCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultCaseContext {
	var p = new(DefaultCaseContext)

	InitEmptyDefault_caseContext(&p.Default_caseContext)
	p.parser = parser
	p.CopyAll(ctx.(*Default_caseContext))

	return p
}

func (s *DefaultCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultCaseContext) DEFAULT_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageDEFAULT_KW, 0)
}

func (s *DefaultCaseContext) COLON() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageCOLON, 0)
}

func (s *DefaultCaseContext) AllStmt() []IStmtContext {
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

func (s *DefaultCaseContext) Stmt(i int) IStmtContext {
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

func (s *DefaultCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitDefaultCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Default_case() (localctx IDefault_caseContext) {
	localctx = NewDefault_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, TSwiftLanguageRULE_default_case)
	var _la int

	localctx = NewDefaultCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(TSwiftLanguageDEFAULT_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
		p.Match(TSwiftLanguageCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&269470688) != 0 {
		{
			p.SetState(312)
			p.Stmt()
		}

		p.SetState(317)
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

// IWhile_stmtContext is an interface to support dynamic dispatch.
type IWhile_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWhile_stmtContext differentiates from other interfaces.
	IsWhile_stmtContext()
}

type While_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_stmtContext() *While_stmtContext {
	var p = new(While_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_while_stmt
	return p
}

func InitEmptyWhile_stmtContext(p *While_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_while_stmt
}

func (*While_stmtContext) IsWhile_stmtContext() {}

func NewWhile_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_stmtContext {
	var p = new(While_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_while_stmt

	return p
}

func (s *While_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *While_stmtContext) CopyAll(ctx *While_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *While_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type WhileStmtContext struct {
	While_stmtContext
}

func NewWhileStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStmtContext {
	var p = new(WhileStmtContext)

	InitEmptyWhile_stmtContext(&p.While_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*While_stmtContext))

	return p
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) WHILE_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageWHILE_KW, 0)
}

func (s *WhileStmtContext) Expr() IExprContext {
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

func (s *WhileStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLBRACE, 0)
}

func (s *WhileStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRBRACE, 0)
}

func (s *WhileStmtContext) AllStmt() []IStmtContext {
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

func (s *WhileStmtContext) Stmt(i int) IStmtContext {
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

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) While_stmt() (localctx IWhile_stmtContext) {
	localctx = NewWhile_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TSwiftLanguageRULE_while_stmt)
	var _la int

	localctx = NewWhileStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		p.Match(TSwiftLanguageWHILE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(319)
		p.expr(0)
	}
	{
		p.SetState(320)
		p.Match(TSwiftLanguageLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&269470688) != 0 {
		{
			p.SetState(321)
			p.Stmt()
		}

		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(327)
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

// IFor_stmtContext is an interface to support dynamic dispatch.
type IFor_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFor_stmtContext differentiates from other interfaces.
	IsFor_stmtContext()
}

type For_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_stmtContext() *For_stmtContext {
	var p = new(For_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_for_stmt
	return p
}

func InitEmptyFor_stmtContext(p *For_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_for_stmt
}

func (*For_stmtContext) IsFor_stmtContext() {}

func NewFor_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_stmtContext {
	var p = new(For_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_for_stmt

	return p
}

func (s *For_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *For_stmtContext) CopyAll(ctx *For_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *For_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForStmtContext struct {
	For_stmtContext
}

func NewForStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStmtContext {
	var p = new(ForStmtContext)

	InitEmptyFor_stmtContext(&p.For_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_stmtContext))

	return p
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) FOR_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageFOR_KW, 0)
}

func (s *ForStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, 0)
}

func (s *ForStmtContext) IN_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageIN_KW, 0)
}

func (s *ForStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLBRACE, 0)
}

func (s *ForStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRBRACE, 0)
}

func (s *ForStmtContext) Expr() IExprContext {
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

func (s *ForStmtContext) Range_() IRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeContext)
}

func (s *ForStmtContext) AllStmt() []IStmtContext {
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

func (s *ForStmtContext) Stmt(i int) IStmtContext {
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

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) For_stmt() (localctx IFor_stmtContext) {
	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TSwiftLanguageRULE_for_stmt)
	var _la int

	localctx = NewForStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		p.Match(TSwiftLanguageFOR_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(330)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)
		p.Match(TSwiftLanguageIN_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(332)
			p.expr(0)
		}

	case 2:
		{
			p.SetState(333)
			p.Range_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(336)
		p.Match(TSwiftLanguageLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&269470688) != 0 {
		{
			p.SetState(337)
			p.Stmt()
		}

		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(343)
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

// IRangeContext is an interface to support dynamic dispatch.
type IRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRangeContext differentiates from other interfaces.
	IsRangeContext()
}

type RangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeContext() *RangeContext {
	var p = new(RangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_range
	return p
}

func InitEmptyRangeContext(p *RangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_range
}

func (*RangeContext) IsRangeContext() {}

func NewRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeContext {
	var p = new(RangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_range

	return p
}

func (s *RangeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeContext) CopyAll(ctx *RangeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumericRangeContext struct {
	RangeContext
}

func NewNumericRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumericRangeContext {
	var p = new(NumericRangeContext)

	InitEmptyRangeContext(&p.RangeContext)
	p.parser = parser
	p.CopyAll(ctx.(*RangeContext))

	return p
}

func (s *NumericRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericRangeContext) AllExpr() []IExprContext {
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

func (s *NumericRangeContext) Expr(i int) IExprContext {
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

func (s *NumericRangeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TSwiftLanguageDOT)
}

func (s *NumericRangeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageDOT, i)
}

func (s *NumericRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitNumericRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Range_() (localctx IRangeContext) {
	localctx = NewRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, TSwiftLanguageRULE_range)
	localctx = NewNumericRangeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.expr(0)
	}
	{
		p.SetState(346)
		p.Match(TSwiftLanguageDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(347)
		p.Match(TSwiftLanguageDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
		p.Match(TSwiftLanguageDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(349)
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

// IGuard_stmtContext is an interface to support dynamic dispatch.
type IGuard_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsGuard_stmtContext differentiates from other interfaces.
	IsGuard_stmtContext()
}

type Guard_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuard_stmtContext() *Guard_stmtContext {
	var p = new(Guard_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_guard_stmt
	return p
}

func InitEmptyGuard_stmtContext(p *Guard_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_guard_stmt
}

func (*Guard_stmtContext) IsGuard_stmtContext() {}

func NewGuard_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Guard_stmtContext {
	var p = new(Guard_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_guard_stmt

	return p
}

func (s *Guard_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Guard_stmtContext) CopyAll(ctx *Guard_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Guard_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Guard_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GuardStmtContext struct {
	Guard_stmtContext
}

func NewGuardStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GuardStmtContext {
	var p = new(GuardStmtContext)

	InitEmptyGuard_stmtContext(&p.Guard_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Guard_stmtContext))

	return p
}

func (s *GuardStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardStmtContext) GUARD_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageGUARD_KW, 0)
}

func (s *GuardStmtContext) Expr() IExprContext {
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

func (s *GuardStmtContext) ELSE_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageELSE_KW, 0)
}

func (s *GuardStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLBRACE, 0)
}

func (s *GuardStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRBRACE, 0)
}

func (s *GuardStmtContext) AllStmt() []IStmtContext {
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

func (s *GuardStmtContext) Stmt(i int) IStmtContext {
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

func (s *GuardStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitGuardStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Guard_stmt() (localctx IGuard_stmtContext) {
	localctx = NewGuard_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, TSwiftLanguageRULE_guard_stmt)
	var _la int

	localctx = NewGuardStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		p.Match(TSwiftLanguageGUARD_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(352)
		p.expr(0)
	}
	{
		p.SetState(353)
		p.Match(TSwiftLanguageELSE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(354)
		p.Match(TSwiftLanguageLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&269470688) != 0 {
		{
			p.SetState(355)
			p.Stmt()
		}

		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(361)
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

// ITransfer_stmtContext is an interface to support dynamic dispatch.
type ITransfer_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTransfer_stmtContext differentiates from other interfaces.
	IsTransfer_stmtContext()
}

type Transfer_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransfer_stmtContext() *Transfer_stmtContext {
	var p = new(Transfer_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_transfer_stmt
	return p
}

func InitEmptyTransfer_stmtContext(p *Transfer_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_transfer_stmt
}

func (*Transfer_stmtContext) IsTransfer_stmtContext() {}

func NewTransfer_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Transfer_stmtContext {
	var p = new(Transfer_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_transfer_stmt

	return p
}

func (s *Transfer_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Transfer_stmtContext) CopyAll(ctx *Transfer_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Transfer_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Transfer_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ContinueStmtContext struct {
	Transfer_stmtContext
}

func NewContinueStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	InitEmptyTransfer_stmtContext(&p.Transfer_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Transfer_stmtContext))

	return p
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) CONTINUE_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageCONTINUE_KW, 0)
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStmtContext struct {
	Transfer_stmtContext
}

func NewBreakStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStmtContext {
	var p = new(BreakStmtContext)

	InitEmptyTransfer_stmtContext(&p.Transfer_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Transfer_stmtContext))

	return p
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) BREAK_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageBREAK_KW, 0)
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStmtContext struct {
	Transfer_stmtContext
}

func NewReturnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	InitEmptyTransfer_stmtContext(&p.Transfer_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Transfer_stmtContext))

	return p
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) RETURN_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRETURN_KW, 0)
}

func (s *ReturnStmtContext) Expr() IExprContext {
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

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Transfer_stmt() (localctx ITransfer_stmtContext) {
	localctx = NewTransfer_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, TSwiftLanguageRULE_transfer_stmt)
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSwiftLanguageRETURN_KW:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(363)
			p.Match(TSwiftLanguageRETURN_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(365)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(364)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case TSwiftLanguageBREAK_KW:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(367)
			p.Match(TSwiftLanguageBREAK_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TSwiftLanguageCONTINUE_KW:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(368)
			p.Match(TSwiftLanguageCONTINUE_KW)
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

// IFunc_callContext is an interface to support dynamic dispatch.
type IFunc_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_callContext differentiates from other interfaces.
	IsFunc_callContext()
}

type Func_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_callContext() *Func_callContext {
	var p = new(Func_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_func_call
	return p
}

func InitEmptyFunc_callContext(p *Func_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_func_call
}

func (*Func_callContext) IsFunc_callContext() {}

func NewFunc_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_callContext {
	var p = new(Func_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_func_call

	return p
}

func (s *Func_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_callContext) CopyAll(ctx *Func_callContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncCallContext struct {
	Func_callContext
}

func NewFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallContext {
	var p = new(FuncCallContext)

	InitEmptyFunc_callContext(&p.Func_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_callContext))

	return p
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *FuncCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLPAREN, 0)
}

func (s *FuncCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRPAREN, 0)
}

func (s *FuncCallContext) Arg_list() IArg_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArg_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArg_listContext)
}

func (s *FuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Func_call() (localctx IFunc_callContext) {
	localctx = NewFunc_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, TSwiftLanguageRULE_func_call)
	var _la int

	localctx = NewFuncCallContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)
		p.Id_pattern()
	}
	{
		p.SetState(372)
		p.Match(TSwiftLanguageLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&145346642701189120) != 0 {
		{
			p.SetState(373)
			p.Arg_list()
		}

	}
	{
		p.SetState(376)
		p.Match(TSwiftLanguageRPAREN)
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

// IArg_listContext is an interface to support dynamic dispatch.
type IArg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArg_listContext differentiates from other interfaces.
	IsArg_listContext()
}

type Arg_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArg_listContext() *Arg_listContext {
	var p = new(Arg_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_arg_list
	return p
}

func InitEmptyArg_listContext(p *Arg_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_arg_list
}

func (*Arg_listContext) IsArg_listContext() {}

func NewArg_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arg_listContext {
	var p = new(Arg_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_arg_list

	return p
}

func (s *Arg_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Arg_listContext) CopyAll(ctx *Arg_listContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArgListContext struct {
	Arg_listContext
}

func NewArgListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgListContext {
	var p = new(ArgListContext)

	InitEmptyArg_listContext(&p.Arg_listContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arg_listContext))

	return p
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) AllFunc_arg() []IFunc_argContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunc_argContext); ok {
			len++
		}
	}

	tst := make([]IFunc_argContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunc_argContext); ok {
			tst[i] = t.(IFunc_argContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Func_arg(i int) IFunc_argContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_argContext); ok {
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

	return t.(IFunc_argContext)
}

func (s *ArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TSwiftLanguageCOMMA)
}

func (s *ArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageCOMMA, i)
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Arg_list() (localctx IArg_listContext) {
	localctx = NewArg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, TSwiftLanguageRULE_arg_list)
	var _la int

	localctx = NewArgListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Func_arg()
	}
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TSwiftLanguageCOMMA {
		{
			p.SetState(379)
			p.Match(TSwiftLanguageCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.Func_arg()
		}

		p.SetState(385)
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

// IFunc_argContext is an interface to support dynamic dispatch.
type IFunc_argContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_argContext differentiates from other interfaces.
	IsFunc_argContext()
}

type Func_argContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_argContext() *Func_argContext {
	var p = new(Func_argContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_func_arg
	return p
}

func InitEmptyFunc_argContext(p *Func_argContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_func_arg
}

func (*Func_argContext) IsFunc_argContext() {}

func NewFunc_argContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_argContext {
	var p = new(Func_argContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_func_arg

	return p
}

func (s *Func_argContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_argContext) CopyAll(ctx *Func_argContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_argContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_argContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncArgContext struct {
	Func_argContext
}

func NewFuncArgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncArgContext {
	var p = new(FuncArgContext)

	InitEmptyFunc_argContext(&p.Func_argContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_argContext))

	return p
}

func (s *FuncArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *FuncArgContext) Expr() IExprContext {
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

func (s *FuncArgContext) ID() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, 0)
}

func (s *FuncArgContext) COLON() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageCOLON, 0)
}

func (s *FuncArgContext) ANPERSAND() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageANPERSAND, 0)
}

func (s *FuncArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitFuncArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Func_arg() (localctx IFunc_argContext) {
	localctx = NewFunc_argContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, TSwiftLanguageRULE_func_arg)
	var _la int

	localctx = NewFuncArgContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(388)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(386)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(387)
			p.Match(TSwiftLanguageCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSwiftLanguageANPERSAND {
		{
			p.SetState(390)
			p.Match(TSwiftLanguageANPERSAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(393)
			p.Id_pattern()
		}

	case 2:
		{
			p.SetState(394)
			p.expr(0)
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

// IFunc_dclContext is an interface to support dynamic dispatch.
type IFunc_dclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_dclContext differentiates from other interfaces.
	IsFunc_dclContext()
}

type Func_dclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_dclContext() *Func_dclContext {
	var p = new(Func_dclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_func_dcl
	return p
}

func InitEmptyFunc_dclContext(p *Func_dclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_func_dcl
}

func (*Func_dclContext) IsFunc_dclContext() {}

func NewFunc_dclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_dclContext {
	var p = new(Func_dclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_func_dcl

	return p
}

func (s *Func_dclContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_dclContext) CopyAll(ctx *Func_dclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_dclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_dclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncDeclContext struct {
	Func_dclContext
}

func NewFuncDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncDeclContext {
	var p = new(FuncDeclContext)

	InitEmptyFunc_dclContext(&p.Func_dclContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_dclContext))

	return p
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) FUNC_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageFUNC_KW, 0)
}

func (s *FuncDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, 0)
}

func (s *FuncDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLPAREN, 0)
}

func (s *FuncDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRPAREN, 0)
}

func (s *FuncDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLBRACE, 0)
}

func (s *FuncDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRBRACE, 0)
}

func (s *FuncDeclContext) Param_list() IParam_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_listContext)
}

func (s *FuncDeclContext) ARROW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageARROW, 0)
}

func (s *FuncDeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FuncDeclContext) AllStmt() []IStmtContext {
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

func (s *FuncDeclContext) Stmt(i int) IStmtContext {
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

func (s *FuncDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitFuncDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Func_dcl() (localctx IFunc_dclContext) {
	localctx = NewFunc_dclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, TSwiftLanguageRULE_func_dcl)
	var _la int

	localctx = NewFuncDeclContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		p.Match(TSwiftLanguageFUNC_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(399)
		p.Match(TSwiftLanguageLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSwiftLanguageID {
		{
			p.SetState(400)
			p.Param_list()
		}

	}
	{
		p.SetState(403)
		p.Match(TSwiftLanguageRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSwiftLanguageARROW {
		{
			p.SetState(404)
			p.Match(TSwiftLanguageARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.Type_()
		}

	}
	{
		p.SetState(408)
		p.Match(TSwiftLanguageLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&269470688) != 0 {
		{
			p.SetState(409)
			p.Stmt()
		}

		p.SetState(414)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(415)
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

// IParam_listContext is an interface to support dynamic dispatch.
type IParam_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParam_listContext differentiates from other interfaces.
	IsParam_listContext()
}

type Param_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_listContext() *Param_listContext {
	var p = new(Param_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_param_list
	return p
}

func InitEmptyParam_listContext(p *Param_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_param_list
}

func (*Param_listContext) IsParam_listContext() {}

func NewParam_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_listContext {
	var p = new(Param_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_param_list

	return p
}

func (s *Param_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_listContext) CopyAll(ctx *Param_listContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Param_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParamListContext struct {
	Param_listContext
}

func NewParamListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamListContext {
	var p = new(ParamListContext)

	InitEmptyParam_listContext(&p.Param_listContext)
	p.parser = parser
	p.CopyAll(ctx.(*Param_listContext))

	return p
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) AllFunc_param() []IFunc_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunc_paramContext); ok {
			len++
		}
	}

	tst := make([]IFunc_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunc_paramContext); ok {
			tst[i] = t.(IFunc_paramContext)
			i++
		}
	}

	return tst
}

func (s *ParamListContext) Func_param(i int) IFunc_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_paramContext); ok {
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

	return t.(IFunc_paramContext)
}

func (s *ParamListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TSwiftLanguageCOMMA)
}

func (s *ParamListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageCOMMA, i)
}

func (s *ParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitParamList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Param_list() (localctx IParam_listContext) {
	localctx = NewParam_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, TSwiftLanguageRULE_param_list)
	var _la int

	localctx = NewParamListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(417)
		p.Func_param()
	}
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TSwiftLanguageCOMMA {
		{
			p.SetState(418)
			p.Match(TSwiftLanguageCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.Func_param()
		}

		p.SetState(424)
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

// IFunc_paramContext is an interface to support dynamic dispatch.
type IFunc_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_paramContext differentiates from other interfaces.
	IsFunc_paramContext()
}

type Func_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_paramContext() *Func_paramContext {
	var p = new(Func_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_func_param
	return p
}

func InitEmptyFunc_paramContext(p *Func_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_func_param
}

func (*Func_paramContext) IsFunc_paramContext() {}

func NewFunc_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_paramContext {
	var p = new(Func_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_func_param

	return p
}

func (s *Func_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_paramContext) CopyAll(ctx *Func_paramContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncParamContext struct {
	Func_paramContext
}

func NewFuncParamContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncParamContext {
	var p = new(FuncParamContext)

	InitEmptyFunc_paramContext(&p.Func_paramContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_paramContext))

	return p
}

func (s *FuncParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TSwiftLanguageID)
}

func (s *FuncParamContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, i)
}

func (s *FuncParamContext) COLON() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageCOLON, 0)
}

func (s *FuncParamContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FuncParamContext) INOUT_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageINOUT_KW, 0)
}

func (s *FuncParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitFuncParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Func_param() (localctx IFunc_paramContext) {
	localctx = NewFunc_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, TSwiftLanguageRULE_func_param)
	var _la int

	localctx = NewFuncParamContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(426)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(425)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(428)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(429)
		p.Match(TSwiftLanguageCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSwiftLanguageINOUT_KW {
		{
			p.SetState(430)
			p.Match(TSwiftLanguageINOUT_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(433)
		p.Type_()
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

// IStrct_dclContext is an interface to support dynamic dispatch.
type IStrct_dclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStrct_dclContext differentiates from other interfaces.
	IsStrct_dclContext()
}

type Strct_dclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrct_dclContext() *Strct_dclContext {
	var p = new(Strct_dclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_strct_dcl
	return p
}

func InitEmptyStrct_dclContext(p *Strct_dclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_strct_dcl
}

func (*Strct_dclContext) IsStrct_dclContext() {}

func NewStrct_dclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Strct_dclContext {
	var p = new(Strct_dclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_strct_dcl

	return p
}

func (s *Strct_dclContext) GetParser() antlr.Parser { return s.parser }

func (s *Strct_dclContext) CopyAll(ctx *Strct_dclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Strct_dclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Strct_dclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructDeclContext struct {
	Strct_dclContext
}

func NewStructDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDeclContext {
	var p = new(StructDeclContext)

	InitEmptyStrct_dclContext(&p.Strct_dclContext)
	p.parser = parser
	p.CopyAll(ctx.(*Strct_dclContext))

	return p
}

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclContext) STRUCT_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageSTRUCT_KW, 0)
}

func (s *StructDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, 0)
}

func (s *StructDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLBRACE, 0)
}

func (s *StructDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRBRACE, 0)
}

func (s *StructDeclContext) AllStruct_prop() []IStruct_propContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStruct_propContext); ok {
			len++
		}
	}

	tst := make([]IStruct_propContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStruct_propContext); ok {
			tst[i] = t.(IStruct_propContext)
			i++
		}
	}

	return tst
}

func (s *StructDeclContext) Struct_prop(i int) IStruct_propContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_propContext); ok {
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

	return t.(IStruct_propContext)
}

func (s *StructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitStructDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Strct_dcl() (localctx IStrct_dclContext) {
	localctx = NewStrct_dclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, TSwiftLanguageRULE_strct_dcl)
	var _la int

	localctx = NewStructDeclContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(435)
		p.Match(TSwiftLanguageSTRUCT_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(436)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(437)
		p.Match(TSwiftLanguageLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2097376) != 0 {
		{
			p.SetState(438)
			p.Struct_prop()
		}

		p.SetState(443)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(444)
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

// IStruct_propContext is an interface to support dynamic dispatch.
type IStruct_propContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_propContext differentiates from other interfaces.
	IsStruct_propContext()
}

type Struct_propContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_propContext() *Struct_propContext {
	var p = new(Struct_propContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_struct_prop
	return p
}

func InitEmptyStruct_propContext(p *Struct_propContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_struct_prop
}

func (*Struct_propContext) IsStruct_propContext() {}

func NewStruct_propContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_propContext {
	var p = new(Struct_propContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_struct_prop

	return p
}

func (s *Struct_propContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_propContext) CopyAll(ctx *Struct_propContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_propContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_propContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructAttrContext struct {
	Struct_propContext
}

func NewStructAttrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructAttrContext {
	var p = new(StructAttrContext)

	InitEmptyStruct_propContext(&p.Struct_propContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_propContext))

	return p
}

func (s *StructAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAttrContext) Var_type() IVar_typeContext {
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

func (s *StructAttrContext) ID() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, 0)
}

func (s *StructAttrContext) COLON() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageCOLON, 0)
}

func (s *StructAttrContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructAttrContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageEQUALS, 0)
}

func (s *StructAttrContext) Expr() IExprContext {
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

func (s *StructAttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitStructAttr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructFuncContext struct {
	Struct_propContext
}

func NewStructFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructFuncContext {
	var p = new(StructFuncContext)

	InitEmptyStruct_propContext(&p.Struct_propContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_propContext))

	return p
}

func (s *StructFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFuncContext) Func_dcl() IFunc_dclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_dclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_dclContext)
}

func (s *StructFuncContext) MUTATING_KW() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageMUTATING_KW, 0)
}

func (s *StructFuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitStructFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Struct_prop() (localctx IStruct_propContext) {
	localctx = NewStruct_propContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, TSwiftLanguageRULE_struct_prop)
	var _la int

	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSwiftLanguageLET_KW, TSwiftLanguageVAR_KW:
		localctx = NewStructAttrContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(446)
			p.Var_type()
		}
		{
			p.SetState(447)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(450)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TSwiftLanguageCOLON {
			{
				p.SetState(448)
				p.Match(TSwiftLanguageCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(449)
				p.Type_()
			}

		}
		p.SetState(454)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TSwiftLanguageEQUALS {
			{
				p.SetState(452)
				p.Match(TSwiftLanguageEQUALS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(453)
				p.expr(0)
			}

		}

	case TSwiftLanguageFUNC_KW, TSwiftLanguageMUTATING_KW:
		localctx = NewStructFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(457)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TSwiftLanguageMUTATING_KW {
			{
				p.SetState(456)
				p.Match(TSwiftLanguageMUTATING_KW)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(459)
			p.Func_dcl()
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

// IStruct_vectorContext is an interface to support dynamic dispatch.
type IStruct_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_vectorContext differentiates from other interfaces.
	IsStruct_vectorContext()
}

type Struct_vectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_vectorContext() *Struct_vectorContext {
	var p = new(Struct_vectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_struct_vector
	return p
}

func InitEmptyStruct_vectorContext(p *Struct_vectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSwiftLanguageRULE_struct_vector
}

func (*Struct_vectorContext) IsStruct_vectorContext() {}

func NewStruct_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_vectorContext {
	var p = new(Struct_vectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSwiftLanguageRULE_struct_vector

	return p
}

func (s *Struct_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_vectorContext) CopyAll(ctx *Struct_vectorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructVectorContext struct {
	Struct_vectorContext
}

func NewStructVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructVectorContext {
	var p = new(StructVectorContext)

	InitEmptyStruct_vectorContext(&p.Struct_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_vectorContext))

	return p
}

func (s *StructVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructVectorContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLBRACK, 0)
}

func (s *StructVectorContext) ID() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageID, 0)
}

func (s *StructVectorContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRBRACK, 0)
}

func (s *StructVectorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageLPAREN, 0)
}

func (s *StructVectorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TSwiftLanguageRPAREN, 0)
}

func (s *StructVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSwiftLanguageVisitor:
		return t.VisitStructVector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSwiftLanguage) Struct_vector() (localctx IStruct_vectorContext) {
	localctx = NewStruct_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, TSwiftLanguageRULE_struct_vector)
	localctx = NewStructVectorContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(462)
		p.Match(TSwiftLanguageLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(463)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(464)
		p.Match(TSwiftLanguageRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(465)
		p.Match(TSwiftLanguageLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(466)
		p.Match(TSwiftLanguageRPAREN)
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
	case 15:
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
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
