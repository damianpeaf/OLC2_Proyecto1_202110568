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
		"'continue'", "'return'", "'guard'", "'inout'", "'in'", "", "", "",
		"", "'nil'", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'='", "'+='", "'-='",
		"'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'&&'", "'||'", "'!'",
		"'('", "')'", "'{'", "'}'", "'['", "']'", "','", "'.'", "':'", "'->'",
		"'?'", "'&'",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "COMMENT", "MULTILINE_COMMENT", "SEMICOLON", "LET_KW", "VAR_KW",
		"FUNC_KW", "STRUCT_KW", "IF_KW", "ELSE_KW", "SWITCH_KW", "CASE_KW",
		"DEFAULT_KW", "FOR_KW", "WHILE_KW", "BREAK_KW", "CONTINUE_KW", "RETURN_KW",
		"GUARD_KW", "INOUT_KW", "IN_KW", "INTEGER_LITERAL", "FLOAT_LITERAL",
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
		"func_arg", "func_dcl", "param_list", "func_param",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 56, 428, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 5, 0, 68, 8, 0, 10, 0, 12, 0, 71, 9, 0, 1, 0, 3,
		0, 74, 8, 0, 1, 1, 3, 1, 77, 8, 1, 1, 1, 3, 1, 80, 8, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 100, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 3, 3, 120, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 126, 8, 4, 10, 4, 12,
		4, 129, 9, 4, 3, 4, 131, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		4, 5, 140, 8, 5, 11, 5, 12, 5, 141, 1, 6, 1, 6, 3, 6, 146, 8, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 3, 8, 163, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 3, 10, 175, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 3, 12, 193, 8, 12, 1, 13, 1, 13, 1, 13, 5, 13, 198, 8, 13, 10, 13,
		12, 13, 201, 9, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 208, 8, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 3, 15, 223, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 5, 15, 243, 8, 15, 10, 15, 12, 15, 246, 9, 15, 1, 16,
		1, 16, 1, 16, 5, 16, 251, 8, 16, 10, 16, 12, 16, 254, 9, 16, 1, 16, 3,
		16, 257, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 263, 8, 17, 10, 17,
		12, 17, 266, 9, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 5, 18, 273, 8, 18,
		10, 18, 12, 18, 276, 9, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5,
		19, 284, 8, 19, 10, 19, 12, 19, 287, 9, 19, 1, 19, 3, 19, 290, 8, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 298, 8, 20, 10, 20, 12, 20,
		301, 9, 20, 1, 21, 1, 21, 1, 21, 5, 21, 306, 8, 21, 10, 21, 12, 21, 309,
		9, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 315, 8, 22, 10, 22, 12, 22, 318,
		9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 327, 8,
		23, 1, 23, 1, 23, 5, 23, 331, 8, 23, 10, 23, 12, 23, 334, 9, 23, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 5, 25, 349, 8, 25, 10, 25, 12, 25, 352, 9, 25, 1, 25, 1, 25,
		1, 26, 1, 26, 3, 26, 358, 8, 26, 1, 26, 1, 26, 3, 26, 362, 8, 26, 1, 27,
		1, 27, 1, 27, 3, 27, 367, 8, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 5,
		28, 374, 8, 28, 10, 28, 12, 28, 377, 9, 28, 1, 29, 1, 29, 3, 29, 381, 8,
		29, 1, 29, 3, 29, 384, 8, 29, 1, 29, 1, 29, 3, 29, 388, 8, 29, 1, 30, 1,
		30, 1, 30, 1, 30, 3, 30, 394, 8, 30, 1, 30, 1, 30, 1, 30, 3, 30, 399, 8,
		30, 1, 30, 1, 30, 5, 30, 403, 8, 30, 10, 30, 12, 30, 406, 9, 30, 1, 30,
		1, 30, 1, 31, 1, 31, 1, 31, 5, 31, 413, 8, 31, 10, 31, 12, 31, 416, 9,
		31, 1, 32, 3, 32, 419, 8, 32, 1, 32, 1, 32, 1, 32, 3, 32, 424, 8, 32, 1,
		32, 1, 32, 1, 32, 0, 1, 30, 33, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
		58, 60, 62, 64, 0, 8, 1, 0, 5, 6, 1, 0, 34, 35, 1, 0, 33, 35, 2, 0, 29,
		29, 44, 44, 1, 0, 30, 32, 1, 0, 28, 29, 1, 0, 38, 41, 1, 0, 36, 37, 462,
		0, 69, 1, 0, 0, 0, 2, 79, 1, 0, 0, 0, 4, 99, 1, 0, 0, 0, 6, 119, 1, 0,
		0, 0, 8, 121, 1, 0, 0, 0, 10, 134, 1, 0, 0, 0, 12, 145, 1, 0, 0, 0, 14,
		157, 1, 0, 0, 0, 16, 162, 1, 0, 0, 0, 18, 164, 1, 0, 0, 0, 20, 174, 1,
		0, 0, 0, 22, 176, 1, 0, 0, 0, 24, 192, 1, 0, 0, 0, 26, 194, 1, 0, 0, 0,
		28, 207, 1, 0, 0, 0, 30, 222, 1, 0, 0, 0, 32, 247, 1, 0, 0, 0, 34, 258,
		1, 0, 0, 0, 36, 269, 1, 0, 0, 0, 38, 279, 1, 0, 0, 0, 40, 293, 1, 0, 0,
		0, 42, 302, 1, 0, 0, 0, 44, 310, 1, 0, 0, 0, 46, 321, 1, 0, 0, 0, 48, 337,
		1, 0, 0, 0, 50, 343, 1, 0, 0, 0, 52, 361, 1, 0, 0, 0, 54, 363, 1, 0, 0,
		0, 56, 370, 1, 0, 0, 0, 58, 380, 1, 0, 0, 0, 60, 389, 1, 0, 0, 0, 62, 409,
		1, 0, 0, 0, 64, 418, 1, 0, 0, 0, 66, 68, 3, 4, 2, 0, 67, 66, 1, 0, 0, 0,
		68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 73, 1,
		0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 74, 5, 0, 0, 1, 73, 72, 1, 0, 0, 0, 73,
		74, 1, 0, 0, 0, 74, 1, 1, 0, 0, 0, 75, 77, 5, 4, 0, 0, 76, 75, 1, 0, 0,
		0, 76, 77, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 80, 5, 0, 0, 1, 79, 76,
		1, 0, 0, 0, 79, 78, 1, 0, 0, 0, 80, 3, 1, 0, 0, 0, 81, 82, 3, 6, 3, 0,
		82, 83, 3, 2, 1, 0, 83, 100, 1, 0, 0, 0, 84, 85, 3, 24, 12, 0, 85, 86,
		3, 2, 1, 0, 86, 100, 1, 0, 0, 0, 87, 88, 3, 52, 26, 0, 88, 89, 3, 2, 1,
		0, 89, 100, 1, 0, 0, 0, 90, 100, 3, 32, 16, 0, 91, 100, 3, 38, 19, 0, 92,
		100, 3, 44, 22, 0, 93, 100, 3, 46, 23, 0, 94, 100, 3, 50, 25, 0, 95, 96,
		3, 54, 27, 0, 96, 97, 3, 2, 1, 0, 97, 100, 1, 0, 0, 0, 98, 100, 3, 60,
		30, 0, 99, 81, 1, 0, 0, 0, 99, 84, 1, 0, 0, 0, 99, 87, 1, 0, 0, 0, 99,
		90, 1, 0, 0, 0, 99, 91, 1, 0, 0, 0, 99, 92, 1, 0, 0, 0, 99, 93, 1, 0, 0,
		0, 99, 94, 1, 0, 0, 0, 99, 95, 1, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100, 5,
		1, 0, 0, 0, 101, 102, 3, 14, 7, 0, 102, 103, 5, 27, 0, 0, 103, 104, 5,
		53, 0, 0, 104, 105, 3, 16, 8, 0, 105, 106, 5, 33, 0, 0, 106, 107, 3, 30,
		15, 0, 107, 120, 1, 0, 0, 0, 108, 109, 3, 14, 7, 0, 109, 110, 5, 27, 0,
		0, 110, 111, 5, 33, 0, 0, 111, 112, 3, 30, 15, 0, 112, 120, 1, 0, 0, 0,
		113, 114, 3, 14, 7, 0, 114, 115, 5, 27, 0, 0, 115, 116, 5, 53, 0, 0, 116,
		117, 3, 16, 8, 0, 117, 118, 5, 55, 0, 0, 118, 120, 1, 0, 0, 0, 119, 101,
		1, 0, 0, 0, 119, 108, 1, 0, 0, 0, 119, 113, 1, 0, 0, 0, 120, 7, 1, 0, 0,
		0, 121, 130, 5, 49, 0, 0, 122, 127, 3, 30, 15, 0, 123, 124, 5, 51, 0, 0,
		124, 126, 3, 30, 15, 0, 125, 123, 1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127,
		125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 127,
		1, 0, 0, 0, 130, 122, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 132, 1, 0,
		0, 0, 132, 133, 5, 50, 0, 0, 133, 9, 1, 0, 0, 0, 134, 139, 3, 26, 13, 0,
		135, 136, 5, 49, 0, 0, 136, 137, 3, 30, 15, 0, 137, 138, 5, 50, 0, 0, 138,
		140, 1, 0, 0, 0, 139, 135, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 139,
		1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 11, 1, 0, 0, 0, 143, 146, 3, 18,
		9, 0, 144, 146, 3, 20, 10, 0, 145, 143, 1, 0, 0, 0, 145, 144, 1, 0, 0,
		0, 146, 147, 1, 0, 0, 0, 147, 148, 5, 45, 0, 0, 148, 149, 5, 27, 0, 0,
		149, 150, 5, 53, 0, 0, 150, 151, 3, 30, 15, 0, 151, 152, 5, 51, 0, 0, 152,
		153, 5, 27, 0, 0, 153, 154, 5, 53, 0, 0, 154, 155, 3, 30, 15, 0, 155, 156,
		5, 46, 0, 0, 156, 13, 1, 0, 0, 0, 157, 158, 7, 0, 0, 0, 158, 15, 1, 0,
		0, 0, 159, 163, 5, 27, 0, 0, 160, 163, 3, 18, 9, 0, 161, 163, 3, 20, 10,
		0, 162, 159, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 161, 1, 0, 0, 0, 163,
		17, 1, 0, 0, 0, 164, 165, 5, 49, 0, 0, 165, 166, 5, 27, 0, 0, 166, 167,
		5, 50, 0, 0, 167, 19, 1, 0, 0, 0, 168, 175, 3, 22, 11, 0, 169, 170, 5,
		49, 0, 0, 170, 171, 5, 49, 0, 0, 171, 172, 5, 27, 0, 0, 172, 173, 5, 50,
		0, 0, 173, 175, 5, 50, 0, 0, 174, 168, 1, 0, 0, 0, 174, 169, 1, 0, 0, 0,
		175, 21, 1, 0, 0, 0, 176, 177, 5, 49, 0, 0, 177, 178, 3, 20, 10, 0, 178,
		179, 5, 50, 0, 0, 179, 23, 1, 0, 0, 0, 180, 181, 3, 26, 13, 0, 181, 182,
		5, 33, 0, 0, 182, 183, 3, 30, 15, 0, 183, 193, 1, 0, 0, 0, 184, 185, 3,
		26, 13, 0, 185, 186, 7, 1, 0, 0, 186, 187, 3, 30, 15, 0, 187, 193, 1, 0,
		0, 0, 188, 189, 3, 10, 5, 0, 189, 190, 7, 2, 0, 0, 190, 191, 3, 30, 15,
		0, 191, 193, 1, 0, 0, 0, 192, 180, 1, 0, 0, 0, 192, 184, 1, 0, 0, 0, 192,
		188, 1, 0, 0, 0, 193, 25, 1, 0, 0, 0, 194, 199, 5, 27, 0, 0, 195, 196,
		5, 52, 0, 0, 196, 198, 5, 27, 0, 0, 197, 195, 1, 0, 0, 0, 198, 201, 1,
		0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 27, 1, 0, 0,
		0, 201, 199, 1, 0, 0, 0, 202, 208, 5, 22, 0, 0, 203, 208, 5, 23, 0, 0,
		204, 208, 5, 24, 0, 0, 205, 208, 5, 25, 0, 0, 206, 208, 5, 26, 0, 0, 207,
		202, 1, 0, 0, 0, 207, 203, 1, 0, 0, 0, 207, 204, 1, 0, 0, 0, 207, 205,
		1, 0, 0, 0, 207, 206, 1, 0, 0, 0, 208, 29, 1, 0, 0, 0, 209, 210, 6, 15,
		-1, 0, 210, 211, 5, 45, 0, 0, 211, 212, 3, 30, 15, 0, 212, 213, 5, 46,
		0, 0, 213, 223, 1, 0, 0, 0, 214, 223, 3, 54, 27, 0, 215, 223, 3, 26, 13,
		0, 216, 223, 3, 10, 5, 0, 217, 223, 3, 28, 14, 0, 218, 223, 3, 8, 4, 0,
		219, 223, 3, 12, 6, 0, 220, 221, 7, 3, 0, 0, 221, 223, 3, 30, 15, 7, 222,
		209, 1, 0, 0, 0, 222, 214, 1, 0, 0, 0, 222, 215, 1, 0, 0, 0, 222, 216,
		1, 0, 0, 0, 222, 217, 1, 0, 0, 0, 222, 218, 1, 0, 0, 0, 222, 219, 1, 0,
		0, 0, 222, 220, 1, 0, 0, 0, 223, 244, 1, 0, 0, 0, 224, 225, 10, 6, 0, 0,
		225, 226, 7, 4, 0, 0, 226, 243, 3, 30, 15, 7, 227, 228, 10, 5, 0, 0, 228,
		229, 7, 5, 0, 0, 229, 243, 3, 30, 15, 6, 230, 231, 10, 4, 0, 0, 231, 232,
		7, 6, 0, 0, 232, 243, 3, 30, 15, 5, 233, 234, 10, 3, 0, 0, 234, 235, 7,
		7, 0, 0, 235, 243, 3, 30, 15, 4, 236, 237, 10, 2, 0, 0, 237, 238, 5, 42,
		0, 0, 238, 243, 3, 30, 15, 3, 239, 240, 10, 1, 0, 0, 240, 241, 5, 43, 0,
		0, 241, 243, 3, 30, 15, 2, 242, 224, 1, 0, 0, 0, 242, 227, 1, 0, 0, 0,
		242, 230, 1, 0, 0, 0, 242, 233, 1, 0, 0, 0, 242, 236, 1, 0, 0, 0, 242,
		239, 1, 0, 0, 0, 243, 246, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 244, 245,
		1, 0, 0, 0, 245, 31, 1, 0, 0, 0, 246, 244, 1, 0, 0, 0, 247, 252, 3, 34,
		17, 0, 248, 249, 5, 10, 0, 0, 249, 251, 3, 34, 17, 0, 250, 248, 1, 0, 0,
		0, 251, 254, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253,
		256, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 255, 257, 3, 36, 18, 0, 256, 255,
		1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 33, 1, 0, 0, 0, 258, 259, 5, 9,
		0, 0, 259, 260, 3, 30, 15, 0, 260, 264, 5, 47, 0, 0, 261, 263, 3, 4, 2,
		0, 262, 261, 1, 0, 0, 0, 263, 266, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 264,
		265, 1, 0, 0, 0, 265, 267, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 267, 268,
		5, 48, 0, 0, 268, 35, 1, 0, 0, 0, 269, 270, 5, 10, 0, 0, 270, 274, 5, 47,
		0, 0, 271, 273, 3, 4, 2, 0, 272, 271, 1, 0, 0, 0, 273, 276, 1, 0, 0, 0,
		274, 272, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 277, 1, 0, 0, 0, 276,
		274, 1, 0, 0, 0, 277, 278, 5, 48, 0, 0, 278, 37, 1, 0, 0, 0, 279, 280,
		5, 11, 0, 0, 280, 281, 3, 30, 15, 0, 281, 285, 5, 47, 0, 0, 282, 284, 3,
		40, 20, 0, 283, 282, 1, 0, 0, 0, 284, 287, 1, 0, 0, 0, 285, 283, 1, 0,
		0, 0, 285, 286, 1, 0, 0, 0, 286, 289, 1, 0, 0, 0, 287, 285, 1, 0, 0, 0,
		288, 290, 3, 42, 21, 0, 289, 288, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290,
		291, 1, 0, 0, 0, 291, 292, 5, 48, 0, 0, 292, 39, 1, 0, 0, 0, 293, 294,
		5, 12, 0, 0, 294, 295, 3, 30, 15, 0, 295, 299, 5, 53, 0, 0, 296, 298, 3,
		4, 2, 0, 297, 296, 1, 0, 0, 0, 298, 301, 1, 0, 0, 0, 299, 297, 1, 0, 0,
		0, 299, 300, 1, 0, 0, 0, 300, 41, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 302,
		303, 5, 13, 0, 0, 303, 307, 5, 53, 0, 0, 304, 306, 3, 4, 2, 0, 305, 304,
		1, 0, 0, 0, 306, 309, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 307, 308, 1, 0,
		0, 0, 308, 43, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 310, 311, 5, 15, 0, 0,
		311, 312, 3, 30, 15, 0, 312, 316, 5, 47, 0, 0, 313, 315, 3, 4, 2, 0, 314,
		313, 1, 0, 0, 0, 315, 318, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 316, 317,
		1, 0, 0, 0, 317, 319, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 319, 320, 5, 48,
		0, 0, 320, 45, 1, 0, 0, 0, 321, 322, 5, 14, 0, 0, 322, 323, 5, 27, 0, 0,
		323, 326, 5, 21, 0, 0, 324, 327, 3, 30, 15, 0, 325, 327, 3, 48, 24, 0,
		326, 324, 1, 0, 0, 0, 326, 325, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328,
		332, 5, 47, 0, 0, 329, 331, 3, 4, 2, 0, 330, 329, 1, 0, 0, 0, 331, 334,
		1, 0, 0, 0, 332, 330, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 335, 1, 0,
		0, 0, 334, 332, 1, 0, 0, 0, 335, 336, 5, 48, 0, 0, 336, 47, 1, 0, 0, 0,
		337, 338, 3, 30, 15, 0, 338, 339, 5, 52, 0, 0, 339, 340, 5, 52, 0, 0, 340,
		341, 5, 52, 0, 0, 341, 342, 3, 30, 15, 0, 342, 49, 1, 0, 0, 0, 343, 344,
		5, 19, 0, 0, 344, 345, 3, 30, 15, 0, 345, 346, 5, 10, 0, 0, 346, 350, 5,
		47, 0, 0, 347, 349, 3, 4, 2, 0, 348, 347, 1, 0, 0, 0, 349, 352, 1, 0, 0,
		0, 350, 348, 1, 0, 0, 0, 350, 351, 1, 0, 0, 0, 351, 353, 1, 0, 0, 0, 352,
		350, 1, 0, 0, 0, 353, 354, 5, 48, 0, 0, 354, 51, 1, 0, 0, 0, 355, 357,
		5, 18, 0, 0, 356, 358, 3, 30, 15, 0, 357, 356, 1, 0, 0, 0, 357, 358, 1,
		0, 0, 0, 358, 362, 1, 0, 0, 0, 359, 362, 5, 16, 0, 0, 360, 362, 5, 17,
		0, 0, 361, 355, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 361, 360, 1, 0, 0, 0,
		362, 53, 1, 0, 0, 0, 363, 364, 3, 26, 13, 0, 364, 366, 5, 45, 0, 0, 365,
		367, 3, 56, 28, 0, 366, 365, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 368,
		1, 0, 0, 0, 368, 369, 5, 46, 0, 0, 369, 55, 1, 0, 0, 0, 370, 375, 3, 58,
		29, 0, 371, 372, 5, 51, 0, 0, 372, 374, 3, 58, 29, 0, 373, 371, 1, 0, 0,
		0, 374, 377, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 376,
		57, 1, 0, 0, 0, 377, 375, 1, 0, 0, 0, 378, 379, 5, 27, 0, 0, 379, 381,
		5, 53, 0, 0, 380, 378, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381, 383, 1, 0,
		0, 0, 382, 384, 5, 56, 0, 0, 383, 382, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0,
		384, 387, 1, 0, 0, 0, 385, 388, 3, 26, 13, 0, 386, 388, 3, 30, 15, 0, 387,
		385, 1, 0, 0, 0, 387, 386, 1, 0, 0, 0, 388, 59, 1, 0, 0, 0, 389, 390, 5,
		7, 0, 0, 390, 391, 5, 27, 0, 0, 391, 393, 5, 45, 0, 0, 392, 394, 3, 62,
		31, 0, 393, 392, 1, 0, 0, 0, 393, 394, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0,
		395, 398, 5, 46, 0, 0, 396, 397, 5, 54, 0, 0, 397, 399, 3, 16, 8, 0, 398,
		396, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 404,
		5, 47, 0, 0, 401, 403, 3, 4, 2, 0, 402, 401, 1, 0, 0, 0, 403, 406, 1, 0,
		0, 0, 404, 402, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405, 407, 1, 0, 0, 0,
		406, 404, 1, 0, 0, 0, 407, 408, 5, 48, 0, 0, 408, 61, 1, 0, 0, 0, 409,
		414, 3, 64, 32, 0, 410, 411, 5, 51, 0, 0, 411, 413, 3, 64, 32, 0, 412,
		410, 1, 0, 0, 0, 413, 416, 1, 0, 0, 0, 414, 412, 1, 0, 0, 0, 414, 415,
		1, 0, 0, 0, 415, 63, 1, 0, 0, 0, 416, 414, 1, 0, 0, 0, 417, 419, 5, 27,
		0, 0, 418, 417, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0,
		420, 421, 5, 27, 0, 0, 421, 423, 5, 53, 0, 0, 422, 424, 5, 20, 0, 0, 423,
		422, 1, 0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 425, 1, 0, 0, 0, 425, 426,
		3, 16, 8, 0, 426, 65, 1, 0, 0, 0, 43, 69, 73, 76, 79, 99, 119, 127, 130,
		141, 145, 162, 174, 192, 199, 207, 222, 242, 244, 252, 256, 264, 274, 285,
		289, 299, 307, 316, 326, 332, 350, 357, 361, 366, 375, 380, 383, 387, 393,
		398, 404, 414, 418, 423,
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
	TSwiftLanguageIN_KW                 = 21
	TSwiftLanguageINTEGER_LITERAL       = 22
	TSwiftLanguageFLOAT_LITERAL         = 23
	TSwiftLanguageSTRING_LITERAL        = 24
	TSwiftLanguageBOOL_LITERAL          = 25
	TSwiftLanguageNIL_LITERAL           = 26
	TSwiftLanguageID                    = 27
	TSwiftLanguagePLUS                  = 28
	TSwiftLanguageMINUS                 = 29
	TSwiftLanguageMULT                  = 30
	TSwiftLanguageDIV                   = 31
	TSwiftLanguageMOD                   = 32
	TSwiftLanguageEQUALS                = 33
	TSwiftLanguagePLUS_EQUALS           = 34
	TSwiftLanguageMINUS_EQUALS          = 35
	TSwiftLanguageEQUALS_EQUALS         = 36
	TSwiftLanguageNOT_EQUALS            = 37
	TSwiftLanguageLESS_THAN             = 38
	TSwiftLanguageLESS_THAN_OR_EQUAL    = 39
	TSwiftLanguageGREATER_THAN          = 40
	TSwiftLanguageGREATER_THAN_OR_EQUAL = 41
	TSwiftLanguageAND                   = 42
	TSwiftLanguageOR                    = 43
	TSwiftLanguageNOT                   = 44
	TSwiftLanguageLPAREN                = 45
	TSwiftLanguageRPAREN                = 46
	TSwiftLanguageLBRACE                = 47
	TSwiftLanguageRBRACE                = 48
	TSwiftLanguageLBRACK                = 49
	TSwiftLanguageRBRACK                = 50
	TSwiftLanguageCOMMA                 = 51
	TSwiftLanguageDOT                   = 52
	TSwiftLanguageCOLON                 = 53
	TSwiftLanguageARROW                 = 54
	TSwiftLanguageINTERROGATION         = 55
	TSwiftLanguageANPERSAND             = 56
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
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135252704) != 0 {
		{
			p.SetState(66)
			p.Stmt()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(72)
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

	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TSwiftLanguageSEMICOLON {
			{
				p.SetState(75)
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
			p.SetState(78)
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
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(81)
			p.Decl_stmt()
		}
		{
			p.SetState(82)
			p.Delimiter()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.Assign_stmt()
		}
		{
			p.SetState(85)
			p.Delimiter()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(87)
			p.Transfer_stmt()
		}
		{
			p.SetState(88)
			p.Delimiter()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(90)
			p.If_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(91)
			p.Switch_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(92)
			p.While_stmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(93)
			p.For_stmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(94)
			p.Guard_stmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(95)
			p.Func_call()
		}
		{
			p.SetState(96)
			p.Delimiter()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(98)
			p.Func_dcl()
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
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeValueDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.Var_type()
		}
		{
			p.SetState(102)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(TSwiftLanguageCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Type_()
		}
		{
			p.SetState(105)
			p.Match(TSwiftLanguageEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.expr(0)
		}

	case 2:
		localctx = NewValueDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
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
			p.Match(TSwiftLanguageEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.expr(0)
		}

	case 3:
		localctx = NewTypeDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(113)
			p.Var_type()
		}
		{
			p.SetState(114)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Match(TSwiftLanguageCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Type_()
		}
		{
			p.SetState(117)
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
		p.SetState(121)
		p.Match(TSwiftLanguageLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&615727312666624) != 0 {
		{
			p.SetState(122)
			p.expr(0)
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TSwiftLanguageCOMMA {
			{
				p.SetState(123)
				p.Match(TSwiftLanguageCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(124)
				p.expr(0)
			}

			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(132)
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
		p.SetState(134)
		p.Id_pattern()
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(135)
				p.Match(TSwiftLanguageLBRACK)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(136)
				p.expr(0)
			}
			{
				p.SetState(137)
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

		p.SetState(141)
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
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(143)
			p.Vector_type()
		}

	case 2:
		{
			p.SetState(144)
			p.Matrix_type()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(147)
		p.Match(TSwiftLanguageLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(TSwiftLanguageCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.expr(0)
	}
	{
		p.SetState(151)
		p.Match(TSwiftLanguageCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(TSwiftLanguageCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.expr(0)
	}
	{
		p.SetState(155)
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
		p.SetState(157)
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
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(160)
			p.Vector_type()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(161)
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
		p.SetState(164)
		p.Match(TSwiftLanguageLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
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
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)
			p.Aux_matrix_type()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(169)
			p.Match(TSwiftLanguageLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.Match(TSwiftLanguageLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Match(TSwiftLanguageRBRACK)
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
		p.SetState(176)
		p.Match(TSwiftLanguageLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Matrix_type()
	}
	{
		p.SetState(178)
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

	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDirectAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Id_pattern()
		}
		{
			p.SetState(181)
			p.Match(TSwiftLanguageEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.expr(0)
		}

	case 2:
		localctx = NewArithmeticAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.Id_pattern()
		}
		{
			p.SetState(185)

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
			p.SetState(186)
			p.expr(0)
		}

	case 3:
		localctx = NewVectorAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(188)
			p.Vector_item()
		}
		{
			p.SetState(189)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*VectorAssignContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&60129542144) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*VectorAssignContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(190)
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
		p.SetState(194)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(199)
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
				p.SetState(195)
				p.Match(TSwiftLanguageDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(196)
				p.Match(TSwiftLanguageID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(201)
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
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSwiftLanguageINTEGER_LITERAL:
		localctx = NewIntLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
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
			p.SetState(203)
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
			p.SetState(204)
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
			p.SetState(205)
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
			p.SetState(206)
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
	p.SetState(222)
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
			p.SetState(210)
			p.Match(TSwiftLanguageLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.expr(0)
		}
		{
			p.SetState(212)
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
			p.SetState(214)
			p.Func_call()
		}

	case 3:
		localctx = NewIdExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(215)
			p.Id_pattern()
		}

	case 4:
		localctx = NewVectorItemExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(216)
			p.Vector_item()
		}

	case 5:
		localctx = NewLiteralExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(217)
			p.Literal()
		}

	case 6:
		localctx = NewVectorExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(218)
			p.Vector_expr()
		}

	case 7:
		localctx = NewRepeatingExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(219)
			p.Repeating()
		}

	case 8:
		localctx = NewUnaryExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(220)

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
			p.SetState(221)
			p.expr(7)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(244)
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
			p.SetState(242)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TSwiftLanguageRULE_expr)
				p.SetState(224)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(225)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7516192768) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(226)

					var _x = p.expr(7)

					localctx.(*BinaryExpContext).right = _x
				}

			case 2:
				localctx = NewBinaryExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TSwiftLanguageRULE_expr)
				p.SetState(227)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(228)

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
					p.SetState(229)

					var _x = p.expr(6)

					localctx.(*BinaryExpContext).right = _x
				}

			case 3:
				localctx = NewBinaryExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TSwiftLanguageRULE_expr)
				p.SetState(230)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(231)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4123168604160) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(232)

					var _x = p.expr(5)

					localctx.(*BinaryExpContext).right = _x
				}

			case 4:
				localctx = NewBinaryExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TSwiftLanguageRULE_expr)
				p.SetState(233)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(234)

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
					p.SetState(235)

					var _x = p.expr(4)

					localctx.(*BinaryExpContext).right = _x
				}

			case 5:
				localctx = NewBinaryExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TSwiftLanguageRULE_expr)
				p.SetState(236)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(237)

					var _m = p.Match(TSwiftLanguageAND)

					localctx.(*BinaryExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(238)

					var _x = p.expr(3)

					localctx.(*BinaryExpContext).right = _x
				}

			case 6:
				localctx = NewBinaryExpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TSwiftLanguageRULE_expr)
				p.SetState(239)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(240)

					var _m = p.Match(TSwiftLanguageOR)

					localctx.(*BinaryExpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(241)

					var _x = p.expr(2)

					localctx.(*BinaryExpContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(246)
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
		p.SetState(247)
		p.If_chain()
	}
	p.SetState(252)
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
				p.SetState(248)
				p.Match(TSwiftLanguageELSE_KW)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(249)
				p.If_chain()
			}

		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSwiftLanguageELSE_KW {
		{
			p.SetState(255)
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
		p.SetState(258)
		p.Match(TSwiftLanguageIF_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(259)
		p.expr(0)
	}
	{
		p.SetState(260)
		p.Match(TSwiftLanguageLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135252704) != 0 {
		{
			p.SetState(261)
			p.Stmt()
		}

		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(267)
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
		p.SetState(269)
		p.Match(TSwiftLanguageELSE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(270)
		p.Match(TSwiftLanguageLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135252704) != 0 {
		{
			p.SetState(271)
			p.Stmt()
		}

		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(277)
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
		p.SetState(279)
		p.Match(TSwiftLanguageSWITCH_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)
		p.expr(0)
	}
	{
		p.SetState(281)
		p.Match(TSwiftLanguageLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TSwiftLanguageCASE_KW {
		{
			p.SetState(282)
			p.Switch_case()
		}

		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSwiftLanguageDEFAULT_KW {
		{
			p.SetState(288)
			p.Default_case()
		}

	}
	{
		p.SetState(291)
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
		p.SetState(293)
		p.Match(TSwiftLanguageCASE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(294)
		p.expr(0)
	}
	{
		p.SetState(295)
		p.Match(TSwiftLanguageCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135252704) != 0 {
		{
			p.SetState(296)
			p.Stmt()
		}

		p.SetState(301)
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
		p.SetState(302)
		p.Match(TSwiftLanguageDEFAULT_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135252704) != 0 {
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
		p.SetState(310)
		p.Match(TSwiftLanguageWHILE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
		p.expr(0)
	}
	{
		p.SetState(312)
		p.Match(TSwiftLanguageLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135252704) != 0 {
		{
			p.SetState(313)
			p.Stmt()
		}

		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(319)
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
		p.SetState(321)
		p.Match(TSwiftLanguageFOR_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.Match(TSwiftLanguageIN_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(324)
			p.expr(0)
		}

	case 2:
		{
			p.SetState(325)
			p.Range_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(328)
		p.Match(TSwiftLanguageLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135252704) != 0 {
		{
			p.SetState(329)
			p.Stmt()
		}

		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(335)
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
		p.SetState(337)
		p.expr(0)
	}
	{
		p.SetState(338)
		p.Match(TSwiftLanguageDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(339)
		p.Match(TSwiftLanguageDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(340)
		p.Match(TSwiftLanguageDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(341)
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
		p.SetState(343)
		p.Match(TSwiftLanguageGUARD_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)
		p.expr(0)
	}
	{
		p.SetState(345)
		p.Match(TSwiftLanguageELSE_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(346)
		p.Match(TSwiftLanguageLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135252704) != 0 {
		{
			p.SetState(347)
			p.Stmt()
		}

		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(353)
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
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSwiftLanguageRETURN_KW:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(355)
			p.Match(TSwiftLanguageRETURN_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(357)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(356)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case TSwiftLanguageBREAK_KW:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(359)
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
			p.SetState(360)
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
		p.SetState(363)
		p.Id_pattern()
	}
	{
		p.SetState(364)
		p.Match(TSwiftLanguageLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72673321350594560) != 0 {
		{
			p.SetState(365)
			p.Arg_list()
		}

	}
	{
		p.SetState(368)
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
		p.SetState(370)
		p.Func_arg()
	}
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TSwiftLanguageCOMMA {
		{
			p.SetState(371)
			p.Match(TSwiftLanguageCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)
			p.Func_arg()
		}

		p.SetState(377)
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
	p.SetState(380)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(378)
			p.Match(TSwiftLanguageID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(379)
			p.Match(TSwiftLanguageCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSwiftLanguageANPERSAND {
		{
			p.SetState(382)
			p.Match(TSwiftLanguageANPERSAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(385)
			p.Id_pattern()
		}

	case 2:
		{
			p.SetState(386)
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
		p.SetState(389)
		p.Match(TSwiftLanguageFUNC_KW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(390)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(391)
		p.Match(TSwiftLanguageLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSwiftLanguageID {
		{
			p.SetState(392)
			p.Param_list()
		}

	}
	{
		p.SetState(395)
		p.Match(TSwiftLanguageRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSwiftLanguageARROW {
		{
			p.SetState(396)
			p.Match(TSwiftLanguageARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)
			p.Type_()
		}

	}
	{
		p.SetState(400)
		p.Match(TSwiftLanguageLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135252704) != 0 {
		{
			p.SetState(401)
			p.Stmt()
		}

		p.SetState(406)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(407)
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
		p.SetState(409)
		p.Func_param()
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TSwiftLanguageCOMMA {
		{
			p.SetState(410)
			p.Match(TSwiftLanguageCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(411)
			p.Func_param()
		}

		p.SetState(416)
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
	p.SetState(418)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(417)
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
		p.SetState(420)
		p.Match(TSwiftLanguageID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(421)
		p.Match(TSwiftLanguageCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSwiftLanguageINOUT_KW {
		{
			p.SetState(422)
			p.Match(TSwiftLanguageINOUT_KW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(425)
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
