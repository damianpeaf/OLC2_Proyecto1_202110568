parser grammar TSwiftLanguage;

options {
	tokenVocab = TSwiftLexer;
	// language = Swift; superClass = SwiftParserBaseListener;
}

program: (stmt)* EOF?;

delimiter: SEMICOLON | EOF;

stmt:
	decl_stmt delimiter
	| assign_stmt delimiter
	| if_stmt
	| switch_stmt;

decl_stmt:
	var_type ID COLON primitive_type EQUALS expr		# TypeValueDecl
	| var_type ID EQUALS expr							# ValueDecl
	| var_type ID COLON primitive_type INTERROGATION	# TypeDecl;

var_type: VAR_KW | LET_KW;

primitive_type:
	INTEGER_TYPE
	| FLOAT_TYPE
	| STRING_TYPE
	| BOOL_TYPE;

assign_stmt:
	id_pattern EQUALS expr								# DirectAssign
	| id_pattern op = (PLUS_EQUALS | MINUS_EQUALS) expr	# ArithmeticAssign;

id_pattern: ID (DOT ID)* # IdPattern;

literal:
	INTEGER_LITERAL		# IntLiteral
	| FLOAT_LITERAL		# FloatLiteral
	| STRING_LITERAL	# StringLiteral
	| BOOL_LITERAL		# BoolLiteral
	| NIL_LITERAL		# NilLiteral;

expr:
	left = expr op = (PLUS | MINUS) right = expr					# BinaryExp // a + b | a - b
	| left = expr op = (MULT | DIV) right = expr					# BinaryExp // a * b | a / b
	| left = expr op = (EQUALS_EQUALS | NOT_EQUALS) right = expr	# BinaryExp // a == b | a != b
	| left = expr op = (
		LESS_THAN
		| LESS_THAN_OR_EQUAL
		| GREATER_THAN
		| GREATER_THAN_OR_EQUAL
	) right = expr						# BinaryExp // a < b | a <= b | a > b | a >= b
	| left = expr op = AND right = expr	# BinaryExp // a && b
	| left = expr op = OR right = expr	# BinaryExp // a || b
	| op = (NOT | MINUS) expr			# UnaryExp // !a | -a
	| LPAREN expr RPAREN				# ParenExp // (a)
	| ID								# IdExp // a
	| literal							# LiteralExp;
// StructMethodCallExp, StructPropertyCallExp, FunctionCallExp, vector, matrix;  (++, --)?

if_stmt: if_chain (ELSE_KW if_chain)* else_stmt? # IfStmt;

if_chain: IF_KW expr LBRACE stmt* RBRACE # IfChain;
else_stmt: ELSE_KW LBRACE stmt* RBRACE # ElseStmt;

switch_stmt:
	SWITCH_KW expr LBRACE switch_case* default_case? RBRACE # SwitchStmt;

switch_case: CASE_KW expr COLON stmt* # SwitchCase;

default_case: DEFAULT_KW COLON stmt* # DefaultCase;