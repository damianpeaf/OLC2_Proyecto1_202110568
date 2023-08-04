package visitor

import (
	"log"
	"main/compiler"
	"main/repl"
	"main/value"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type ReplVisitor struct {
	compiler.BaseTSwiftLanguageVisitor
	ScopeTrace *repl.ScopeTrace
}

func NewVisitor() *ReplVisitor {
	return &ReplVisitor{
		ScopeTrace: repl.NewScopeTrace(),
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

	variable := v.ScopeTrace.AddVariable(varName, varType, varValue, isConst)

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

	variable := v.ScopeTrace.AddVariable(varName, varType, varValue, isConst)

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

	variable := v.ScopeTrace.AddVariable(varName, varType, value.DefaultNilValue, isConst)

	ok, msg := variable.Validate()

	// Add variable to scope
	if !ok {
		log.Fatal(msg)
	}

	return v.VisitChildren(ctx)
}

func (v *ReplVisitor) VisitDirectAssign(ctx *compiler.DirectAssignContext) interface{} {

	varName := v.Visit(ctx.Id_pattern()).(string)
	varValue := v.Visit(ctx.Expr()).(value.IVOR)

	variable := v.ScopeTrace.GetVariable(varName)

	if variable == nil {
		log.Fatal("Variable not found")
	}

	// TODO: asign method
	variable.Value = varValue

	return v.VisitChildren(ctx)

}

func (v *ReplVisitor) VisitArithmeticAssign(ctx *compiler.ArithmeticAssignContext) interface{} {
	varName := v.Visit(ctx.Id_pattern()).(string)

	variable := v.ScopeTrace.GetVariable(varName)

	if variable == nil {
		log.Fatal("Variable not found")
	}

	leftValue := variable.Value
	rightValue := v.Visit(ctx.Expr()).(value.IVOR)

	op := string(ctx.GetOp().GetText()[0])

	strat, ok := repl.BinaryStrats[op]

	if !ok {
		log.Fatal("Binary operator not found")
	}

	ok, msg, result := strat.Validate(leftValue, rightValue)

	if !ok {
		log.Fatal(msg)
	}

	// TODO: asign method
	variable.Value = result

	return nil
}

func (v *ReplVisitor) VisitIdPattern(ctx *compiler.IdPatternContext) interface{} {
	return ctx.GetText()
}

func (v *ReplVisitor) VisitIntLiteral(ctx *compiler.IntLiteralContext) interface{} {

	intVal, _ := strconv.Atoi(ctx.GetText())

	return value.IntValue{
		InternalValue: intVal,
	}

}

func (v *ReplVisitor) VisitFloatLiteral(ctx *compiler.FloatLiteralContext) interface{} {

	floatVal, _ := strconv.ParseFloat(ctx.GetText(), 64)

	return value.FloatValue{
		InternalValue: floatVal,
	}

}

func (v *ReplVisitor) VisitStringLiteral(ctx *compiler.StringLiteralContext) interface{} {

	// remove quotes
	// Todo: scape sequences
	stringVal := ctx.GetText()[1 : len(ctx.GetText())-1]

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

func (v *ReplVisitor) VisitNilLiteral(ctx *compiler.NilLiteralContext) interface{} {
	return value.DefaultNilValue
}

func (v *ReplVisitor) VisitLiteralExp(ctx *compiler.LiteralExpContext) interface{} {
	return v.Visit(ctx.Literal())
}

func (v *ReplVisitor) VisitIdExp(ctx *compiler.IdExpContext) interface{} {
	varName := ctx.ID().GetText()

	// TODO: check if variable exists
	variable := v.ScopeTrace.GetVariable(varName)

	return variable.Value
}

func (v *ReplVisitor) VisitParenExp(ctx *compiler.ParenExpContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *ReplVisitor) VisitUnaryExp(ctx *compiler.UnaryExpContext) interface{} {

	exp := v.Visit(ctx.Expr()).(value.IVOR)

	strat, ok := repl.UnaryStrats[ctx.GetOp().GetText()]

	if !ok {
		log.Fatal("Unary operator not found")
	}

	ok, msg, result := strat.Validate(exp)

	if !ok {
		log.Fatal(msg)
	}

	return result

}

func (v *ReplVisitor) VisitBinaryExp(ctx *compiler.BinaryExpContext) interface{} {

	left := v.Visit(ctx.GetLeft()).(value.IVOR)
	right := v.Visit(ctx.GetRight()).(value.IVOR)

	strat, ok := repl.BinaryStrats[ctx.GetOp().GetText()]

	if !ok {
		log.Fatal("Binary operator not found")
	}

	ok, msg, result := strat.Validate(left, right)

	if !ok {
		log.Fatal(msg)
	}

	return result
}

func (v *ReplVisitor) VisitIf_stmt(ctx *compiler.If_stmtContext) interface{} {
	return nil
}
