// Code generated from compiler\TSwiftLanguage.g4 by ANTLR 4.13.0. DO NOT EDIT.

package compiler // TSwiftLanguage
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by TSwiftLanguage.
type TSwiftLanguageVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TSwiftLanguage#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#delimiter.
	VisitDelimiter(ctx *DelimiterContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#TypeValueDecl.
	VisitTypeValueDecl(ctx *TypeValueDeclContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#ValueDecl.
	VisitValueDecl(ctx *ValueDeclContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#TypeDecl.
	VisitTypeDecl(ctx *TypeDeclContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#var_type.
	VisitVar_type(ctx *Var_typeContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#primitive_type.
	VisitPrimitive_type(ctx *Primitive_typeContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#DirectAssign.
	VisitDirectAssign(ctx *DirectAssignContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#ArithmeticAssign.
	VisitArithmeticAssign(ctx *ArithmeticAssignContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#IdPattern.
	VisitIdPattern(ctx *IdPatternContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#IntLiteral.
	VisitIntLiteral(ctx *IntLiteralContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#FloatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#StringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#BoolLiteral.
	VisitBoolLiteral(ctx *BoolLiteralContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#NilLiteral.
	VisitNilLiteral(ctx *NilLiteralContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#LiteralExp.
	VisitLiteralExp(ctx *LiteralExpContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#IdExp.
	VisitIdExp(ctx *IdExpContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#ParenExp.
	VisitParenExp(ctx *ParenExpContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#UnaryExp.
	VisitUnaryExp(ctx *UnaryExpContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#BinaryExp.
	VisitBinaryExp(ctx *BinaryExpContext) interface{}

	// Visit a parse tree produced by TSwiftLanguage#if_stmt.
	VisitIf_stmt(ctx *If_stmtContext) interface{}
}
