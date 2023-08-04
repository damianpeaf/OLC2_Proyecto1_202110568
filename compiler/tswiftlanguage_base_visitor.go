// Code generated from compiler\TSwiftLanguage.g4 by ANTLR 4.13.0. DO NOT EDIT.

package compiler // TSwiftLanguage
import "github.com/antlr4-go/antlr/v4"

type BaseTSwiftLanguageVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTSwiftLanguageVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitDelimiter(ctx *DelimiterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitTypeValueDecl(ctx *TypeValueDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitValueDecl(ctx *ValueDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitTypeDecl(ctx *TypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitVar_type(ctx *Var_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitPrimitive_type(ctx *Primitive_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitDirectAssign(ctx *DirectAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitArithmeticAssign(ctx *ArithmeticAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitIdPattern(ctx *IdPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitBoolLiteral(ctx *BoolLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitNilLiteral(ctx *NilLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitLiteralExp(ctx *LiteralExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitIdExp(ctx *IdExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitParenExp(ctx *ParenExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitUnaryExp(ctx *UnaryExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitBinaryExp(ctx *BinaryExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitIfChain(ctx *IfChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitElseStmt(ctx *ElseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitSwitchStmt(ctx *SwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitSwitchCase(ctx *SwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSwiftLanguageVisitor) VisitDefaultCase(ctx *DefaultCaseContext) interface{} {
	return v.VisitChildren(ctx)
}
