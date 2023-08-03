package visitor

import (
	"fmt"
	"log"
	"main/compiler"
	"main/repl"
	"main/value"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type ReplVisitor struct {
	compiler.BaseTSwiftLanguageVisitor
	ReplContext *repl.ReplContext
}

func NewVisitor() *ReplVisitor {
	return &ReplVisitor{
		ReplContext: repl.NewReplContext(),
	}
}

func (v *ReplVisitor) Visit(tree antlr.ParseTree) interface{} {

	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}

}

func (v *ReplVisitor) VisitProgram(ctx *compiler.ProgramContext) interface{} {

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	return nil
}

func (v *ReplVisitor) VisitStmt(ctx *compiler.StmtContext) interface{} {

	if ctx.Decl_stmt() != nil {
		v.Visit(ctx.Decl_stmt())
	} else if ctx.Assign_stmt() != nil {
		v.Visit(ctx.Assign_stmt())
	}

	return nil
}

func isDeclConst(lexval string) bool {
	return lexval == "let"
}

func (v *ReplVisitor) VisitTypeValueDecl(ctx *compiler.TypeValueDeclContext) interface{} {

	isConst := isDeclConst(ctx.Var_type().GetText())
	varName := ctx.ID().GetText()
	varType := ctx.Primitive_type().GetText()
	varValue := v.Visit(ctx.Expr()).(value.IVOR)

	variable := v.ReplContext.ScopeTrace.CurrentScope.AddVariable(varName, varType, varValue, isConst)

	ok, msg := variable.Validate()

	// Add variable to scope
	if !ok {
		log.Fatal(msg)
	}

	return v.VisitChildren(ctx)
}

func (v *ReplVisitor) VisitValueDecl(ctx *compiler.ValueDeclContext) interface{} {

	isConst := isDeclConst(ctx.Var_type().GetText())
	varName := ctx.ID().GetText()
	varValue := v.Visit(ctx.Expr()).(value.IVOR)
	varType := varValue.Type()

	variable := v.ReplContext.ScopeTrace.CurrentScope.AddVariable(varName, varType, varValue, isConst)

	ok, msg := variable.Validate()

	// Add variable to scope
	if !ok {
		log.Fatal(msg)
	}

	return v.VisitChildren(ctx)

}

func (v *ReplVisitor) VisitTypeDecl(ctx *compiler.TypeDeclContext) interface{} {

	isConst := isDeclConst(ctx.Var_type().GetText())
	varName := ctx.ID().GetText()
	varType := ctx.Primitive_type().GetText()

	variable := v.ReplContext.ScopeTrace.CurrentScope.AddVariable(varName, varType, value.DefaultNilValue, isConst)

	ok, msg := variable.Validate()

	// Add variable to scope
	if !ok {
		log.Fatal(msg)
	}

	return v.VisitChildren(ctx)
}

func (v *ReplVisitor) VisitAssign(ctx *compiler.AssignContext) interface{} {

	varName := ctx.ID().GetText()
	varValue := v.Visit(ctx.Expr()).(value.IVOR)

	variable := v.ReplContext.ScopeTrace.CurrentScope.GetVariable(varName)

	if variable == nil {
		log.Fatal("Variable not found")
	}

	variable.Value = varValue

	ok, msg := variable.Validate()

	if !ok {
		log.Fatal(msg)
	}

	return v.VisitChildren(ctx)

}

func (v *ReplVisitor) VisitIntLiteral(ctx *compiler.IntLiteralContext) interface{} {

	intVal, _ := strconv.Atoi(ctx.GetText())

	a := value.IntValue{
		InternalValue: intVal,
	}

	fmt.Println(a.Value())
	fmt.Println(a.Type())

	return a

}

func (v *ReplVisitor) VisitFloatLiteral(ctx *compiler.FloatLiteralContext) interface{} {

	floatVal, _ := strconv.ParseFloat(ctx.GetText(), 64)

	return value.FloatValue{
		InternalValue: floatVal,
	}

}

func (v *ReplVisitor) VisitStringLiteral(ctx *compiler.StringLiteralContext) interface{} {

	stringVal := ctx.GetText()

	return value.StringValue{
		InternalValue: stringVal,
	}

}

func (v *ReplVisitor) VisitBoolLiteral(ctx *compiler.BoolLiteralContext) interface{} {

	boolVal, _ := strconv.ParseBool(ctx.GetText())

	return value.BoolValue{
		InternalValue: boolVal,
	}

}

func (v *ReplVisitor) VisitExpr(ctx *compiler.ExprContext) interface{} {

	if ctx.Literal() != nil {
		return v.Visit(ctx.Literal())
	}

	return nil

}

func (v *ReplVisitor) VisitIf_stmt(ctx *compiler.If_stmtContext) interface{} {
	return nil
}
