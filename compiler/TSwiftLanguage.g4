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

literal:
	INTEGER_LITERAL		# IntLiteral
	| FLOAT_LITERAL		# FloatLiteral
	| STRING_LITERAL	# StringLiteral
	| BOOL_LITERAL		# BoolLiteral;

expr: literal;

if_stmt: IF_KW expr LBRACE stmt* RBRACE;

