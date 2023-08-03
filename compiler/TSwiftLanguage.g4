parser grammar TSwiftLanguage;

options {
	tokenVocab = TSwiftLexer;
	// language = Swift; superClass = SwiftParserBaseListener;
}

program: (stmt delimiter)*;

delimiter: SEMICOLON | NEWLINE+ | EOF;

stmt: decl_stmt | assign_stmt;

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

assign_stmt: ID EQUALS expr # Assign;
// struct assign a.b.c = 1; +=, -=

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
	) right = expr				# BinaryExp // a < b | a <= b | a > b | a >= b
	| op = (NOT | MINUS) expr	# UnaryExp // !a | -a
	| LPAREN expr RPAREN		# ParenExp // (a)
	| ID						# IdExp // a
	| literal					# LiteralExp;
// StructMethodCallExp, StructPropertyCallExp, FunctionCallExp, vector, matrix;  (++, --)?

if_stmt: IF_KW expr LBRACE stmt* RBRACE;

