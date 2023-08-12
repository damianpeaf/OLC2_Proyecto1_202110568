package dot

import (
	"log"
	"main/compiler"

	"github.com/antlr4-go/antlr/v4"
)

type DotVisitor struct {
	compiler.BaseTSwiftLanguageVisitor
	output string
}

func (v *DotVisitor) Visit(tree antlr.ParseTree) interface{} {

	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}

}

func (v *DotVisitor) VisitProgram(ctx *compiler.ProgramContext) interface{} {

	v.output += "digraph G {\n"
	v.output += "node [shape=record];\n"

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitDelimiter(ctx *compiler.DelimiterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitStmt(ctx *compiler.StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitTypeValueDecl(ctx *compiler.TypeValueDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitValueDecl(ctx *compiler.ValueDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitTypeDecl(ctx *compiler.TypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitVar_type(ctx *compiler.Var_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitPrimitive_type(ctx *compiler.Primitive_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitDirectAssign(ctx *compiler.DirectAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitArithmeticAssign(ctx *compiler.ArithmeticAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitIdPattern(ctx *compiler.IdPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitIntLiteral(ctx *compiler.IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitFloatLiteral(ctx *compiler.FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitStringLiteral(ctx *compiler.StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitBoolLiteral(ctx *compiler.BoolLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitNilLiteral(ctx *compiler.NilLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitLiteralExp(ctx *compiler.LiteralExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitIdExp(ctx *compiler.IdExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitParenExp(ctx *compiler.ParenExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitFuncCallExp(ctx *compiler.FuncCallExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitUnaryExp(ctx *compiler.UnaryExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitBinaryExp(ctx *compiler.BinaryExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitIfStmt(ctx *compiler.IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitIfChain(ctx *compiler.IfChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitElseStmt(ctx *compiler.ElseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitSwitchStmt(ctx *compiler.SwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitSwitchCase(ctx *compiler.SwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitDefaultCase(ctx *compiler.DefaultCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitWhileStmt(ctx *compiler.WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitForStmt(ctx *compiler.ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitNumericRange(ctx *compiler.NumericRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitGuardStmt(ctx *compiler.GuardStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitReturnStmt(ctx *compiler.ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitBreakStmt(ctx *compiler.BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitContinueStmt(ctx *compiler.ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitFuncCall(ctx *compiler.FuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitArgList(ctx *compiler.ArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitFuncArg(ctx *compiler.FuncArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitFuncDecl(ctx *compiler.FuncDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitParamList(ctx *compiler.ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *DotVisitor) VisitFuncParam(ctx *compiler.FuncParamContext) interface{} {
	return v.VisitChildren(ctx)
}
