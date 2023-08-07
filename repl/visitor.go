package repl

import (
	"log"
	"main/compiler"
	"main/value"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type ReplVisitor struct {
	compiler.BaseTSwiftLanguageVisitor
	ScopeTrace *ScopeTrace
	CallStack  *CallStack
	Console    *Console
	ErrorTable *ErrorTable
}

func NewVisitor() *ReplVisitor {
	return &ReplVisitor{
		ScopeTrace: NewScopeTrace(),
		CallStack:  NewCallStack(),
		Console:    NewConsole(),
		ErrorTable: NewErrorTable(),
	}
}

func (v *ReplVisitor) GetReplContext() *ReplContext {
	return &ReplContext{
		Console:    v.Console,
		ScopeTrace: v.ScopeTrace,
		CallStack:  v.CallStack,
		ErrorTable: v.ErrorTable,
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
	} else if ctx.If_stmt() != nil {
		v.Visit(ctx.If_stmt())
	} else if ctx.Switch_stmt() != nil {
		v.Visit(ctx.Switch_stmt())
	} else if ctx.While_stmt() != nil {
		v.Visit(ctx.While_stmt())
	} else if ctx.For_stmt() != nil {
		v.Visit(ctx.For_stmt())
	} else if ctx.Guard_stmt() != nil {
		v.Visit(ctx.Guard_stmt())
	} else if ctx.Transfer_stmt() != nil {
		v.Visit(ctx.Transfer_stmt())
	} else if ctx.Func_call() != nil {
		v.Visit(ctx.Func_call())
	} else if ctx.Func_dcl() != nil {
		v.Visit(ctx.Func_dcl())
	} else {
		log.Fatal("Statement not found")
	}

	return nil
}

func isDeclConst(lexval string) bool {
	return lexval == "let"
}

func (v *ReplVisitor) VisitTypeValueDecl(ctx *compiler.TypeValueDeclContext) interface{} {

	isConst := isDeclConst(ctx.Var_type().GetText())
	varName := ctx.ID().GetText()
	// todo: Change primitive type to a more general type
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
	// todo: Change primitive type to a more general type
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

	strat, ok := BinaryStrats[op]

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

	return &value.IntValue{
		InternalValue: intVal,
	}

}

func (v *ReplVisitor) VisitFloatLiteral(ctx *compiler.FloatLiteralContext) interface{} {

	floatVal, _ := strconv.ParseFloat(ctx.GetText(), 64)

	return &value.FloatValue{
		InternalValue: floatVal,
	}

}

func (v *ReplVisitor) VisitStringLiteral(ctx *compiler.StringLiteralContext) interface{} {

	// remove quotes
	// Todo: scape sequences
	stringVal := ctx.GetText()[1 : len(ctx.GetText())-1]

	return &value.StringValue{
		InternalValue: stringVal,
	}

}

func (v *ReplVisitor) VisitBoolLiteral(ctx *compiler.BoolLiteralContext) interface{} {

	boolVal, _ := strconv.ParseBool(ctx.GetText())

	return &value.BoolValue{
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

	// ? pointer
	return variable.Value
}

func (v *ReplVisitor) VisitParenExp(ctx *compiler.ParenExpContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *ReplVisitor) VisitFuncCallExp(ctx *compiler.FuncCallExpContext) interface{} {
	return v.Visit(ctx.Func_call())
}

func (v *ReplVisitor) VisitUnaryExp(ctx *compiler.UnaryExpContext) interface{} {

	exp := v.Visit(ctx.Expr()).(value.IVOR)

	strat, ok := UnaryStrats[ctx.GetOp().GetText()]

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

	strat, ok := BinaryStrats[ctx.GetOp().GetText()]

	if !ok {
		log.Fatal("Binary operator not found")
	}

	ok, msg, result := strat.Validate(left, right)

	if !ok {
		log.Fatal(msg)
	}

	return result
}

func (v *ReplVisitor) VisitIfStmt(ctx *compiler.IfStmtContext) interface{} {

	runChain := true

	for _, ifStmt := range ctx.AllIf_chain() {

		runChain = !v.Visit(ifStmt).(bool)

		if !runChain {
			break
		}
	}

	if runChain && ctx.Else_stmt() != nil {
		v.Visit(ctx.Else_stmt())
	}

	return nil
}

func (v *ReplVisitor) VisitIfChain(ctx *compiler.IfChainContext) interface{} {

	condition := v.Visit(ctx.Expr()).(value.IVOR)

	if condition.Type() != value.IVOR_BOOL {
		log.Fatal("Condition must be a boolean")
	}

	if condition.(*value.BoolValue).InternalValue {

		// Push scope
		v.ScopeTrace.PushScope("if")

		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}

		// Pop scope
		v.ScopeTrace.PopScope()

		return true
	}

	return false
}

func (v *ReplVisitor) VisitElseStmt(ctx *compiler.ElseStmtContext) interface{} {

	// Push scope
	v.ScopeTrace.PushScope("else")

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	// Pop scope
	v.ScopeTrace.PopScope()

	return nil
}

func (v *ReplVisitor) VisitSwitchStmt(ctx *compiler.SwitchStmtContext) interface{} {

	mainValue := v.Visit(ctx.Expr()).(value.IVOR)

	v.ScopeTrace.PushScope("switch")

	// Push break switchItem to call stack [breakable]
	switchItem := &CallStackItem{
		ReturnValue: value.DefaultNilValue,
		Type: []string{
			BreakItem,
		},
	}

	v.CallStack.Push(switchItem)

	// handle break statements from call stack
	defer func() {

		v.ScopeTrace.PopScope()       // pop switch scope
		v.CallStack.Clean(switchItem) // clean item if it's still in call stack

		if item, ok := recover().(*CallStackItem); item != nil && ok {

			// Not a switch item, propagate panic
			if item != switchItem {
				panic(item)
			}
		}
	}()

	// evaluate cases
	for _, switchCase := range ctx.AllSwitch_case() {

		caseValue := v.GetCaseValue(switchCase)

		// ? use binary strat
		if caseValue.Type() != mainValue.Type() {
			// warning
			log.Fatal("Case value must be same type as switch value")
		}

		if caseValue.Value() == mainValue.Value() {
			v.Visit(switchCase)
		}

	}

	// evaluate default
	// ! fix: default case must be evaluated after all cases only if no case was evaluated
	if ctx.Default_case() != nil {
		v.Visit(ctx.Default_case())
	}

	return nil
}

func (v *ReplVisitor) GetCaseValue(tree antlr.ParseTree) value.IVOR {

	switch val := tree.(type) {
	case *compiler.SwitchCaseContext:
		return v.Visit(val.Expr()).(value.IVOR)
	default:
		return nil
	}

}

func (v *ReplVisitor) VisitSwitchCase(ctx *compiler.SwitchCaseContext) interface{} {

	// * all cases inside switch case will share the same scope

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}
	return nil
}

func (v *ReplVisitor) VisitDefaultCase(ctx *compiler.DefaultCaseContext) interface{} {
	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}
	return nil
}

func (v *ReplVisitor) VisitWhileStmt(ctx *compiler.WhileStmtContext) interface{} {

	condition := v.Visit(ctx.Expr()).(value.IVOR)
	// Push scope
	whileScope := v.ScopeTrace.PushScope("while")

	// Push whileItem to call stack [breakable, continuable]
	whileItem := &CallStackItem{
		ReturnValue: value.DefaultNilValue,
		Type: []string{
			BreakItem,
			ContinueItem,
		},
	}

	v.CallStack.Push(whileItem)

	v.VisitInnerWhile(ctx, condition, whileScope, whileItem)

	v.ScopeTrace.PopScope()      // pop while scope
	v.CallStack.Clean(whileItem) // clean item if it's still in call stack

	return nil
}

func (v *ReplVisitor) VisitInnerWhile(ctx *compiler.WhileStmtContext, condition value.IVOR, whileScope *BaseScope, whileItem *CallStackItem) {

	// ? use binary strat
	if condition.Type() != value.IVOR_BOOL {
		log.Fatal("Condition must be a boolean")
		return
	}

	// reset scope
	whileScope.Reset()

	// handle break and continue statements from call stack
	defer func() {

		if item, ok := recover().(*CallStackItem); item != nil && ok {

			// Not a while item, propagate panic
			if item != whileItem {
				panic(item)
			}

			// Continue
			if item.IsAction(ContinueItem) {
				item.ResetAction()                                       // reset action, can be used again
				v.VisitInnerWhile(ctx, condition, whileScope, whileItem) // continue

			} else if item.IsAction(BreakItem) {
				// Break
				return
			}
		}
	}()

	for condition.(*value.BoolValue).InternalValue {

		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}

		condition = v.Visit(ctx.Expr()).(value.IVOR)

		if condition.Type() != value.IVOR_BOOL {
			log.Fatal("Condition must be a boolean")
		}

		// reset scope
		whileScope.Reset()
	}
}

func (v *ReplVisitor) VisitForStmt(ctx *compiler.ForStmtContext) interface{} {

	// Push scope
	// mainForScope := v.ScopeTrace.PushScope("for")
	// TODO: implement first 'iterable' values as: strings, vectors and ranges
	// TODO: handle break and continue statements from call stack

	return nil
}

func (v *ReplVisitor) VisitNumericRange(ctx *compiler.NumericRangeContext) interface{} {
	// TODO: implement range
	return v.VisitChildren(ctx)
}

func (v *ReplVisitor) VisitGuardStmt(ctx *compiler.GuardStmtContext) interface{} {

	condition := v.Visit(ctx.Expr()).(value.IVOR)

	if condition.Type() != value.IVOR_BOOL {
		log.Fatal("Condition must be a boolean")
	}

	if !condition.(*value.BoolValue).InternalValue {

		// Push scope
		v.ScopeTrace.PushScope("guard")

		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}

		// Pop scope
		v.ScopeTrace.PopScope()
	}

	return nil
}

func (v *ReplVisitor) VisitReturnStmt(ctx *compiler.ReturnStmtContext) interface{} {

	exits, item := v.CallStack.IsReturnEnv()

	if !exits {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La sentencia return debe estar dentro de una funcion")
		return nil
	}

	item.ReturnValue = value.DefaultNilValue
	item.Action = ReturnItem

	if ctx.Expr() != nil {
		item.ReturnValue = v.Visit(ctx.Expr()).(value.IVOR)
	}

	panic(item)
}

func (v *ReplVisitor) VisitBreakStmt(ctx *compiler.BreakStmtContext) interface{} {

	exits, item := v.CallStack.IsBreakEnv()

	if !exits {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La sentencia break debe estar dentro de un ciclo o un switch")
		return nil
	}

	item.Action = BreakItem
	panic(item)
}

func (v *ReplVisitor) VisitContinueStmt(ctx *compiler.ContinueStmtContext) interface{} {

	exits, item := v.CallStack.IsContinueEnv()

	if !exits {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La sentencia continue debe estar dentro de un ciclo")
		return nil
	}

	item.Action = ContinueItem
	panic(item)
}

func (v *ReplVisitor) VisitFuncCall(ctx *compiler.FuncCallContext) interface{} {

	funcName := v.Visit(ctx.Id_pattern()).(string)

	funcObj := v.ScopeTrace.GetFunction(funcName)

	if funcObj == nil {
		log.Fatal("Function not found")
	}

	args := make([]*Argument, 0)

	if ctx.Arg_list() != nil {
		args = v.Visit(ctx.Arg_list()).([]*Argument)
	}

	switch funcObj.Type() {
	case value.IVOR_BUILTIN_FUNCTION:
		returnValue, ok, msg := funcObj.(*BuiltInFunction).Exec(v.GetReplContext(), args)

		if !ok {
			log.Fatal(msg)
		}

		return returnValue

	case value.IVOR_FUNCTION:
		funcObj.(*Function).Exec(v, args, ctx.GetStart())
		return funcObj.(*Function).ReturnValue

	default:
		log.Fatal("Function type not found")
	}

	return value.DefaultNilValue
}

func (v *ReplVisitor) VisitArgList(ctx *compiler.ArgListContext) interface{} {

	args := make([]*Argument, 0)

	for _, arg := range ctx.AllFunc_arg() {
		args = append(args, v.Visit(arg).(*Argument))
	}

	return args

}

func (v *ReplVisitor) VisitFuncArg(ctx *compiler.FuncArgContext) interface{} {

	argName := ""
	passByReference := false
	argValue, ok := v.Visit(ctx.Expr()).(value.IVOR)

	if !ok {
		panic(ctx.Expr().GetText())
	}

	if ctx.ID() != nil {
		argName = ctx.ID().GetText()
	}

	if ctx.ANPERSAND() != nil {
		passByReference = true
	}

	return &Argument{
		Name:            argName,
		Object:          argValue,
		PassByReference: passByReference,
		Token:           ctx.GetStart(),
	}

}

func (v *ReplVisitor) VisitFuncDecl(ctx *compiler.FuncDeclContext) interface{} {

	// TODO: supoort for structs functions
	if v.ScopeTrace.CurrentScope != v.ScopeTrace.GlobalScope {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Las funciones solo pueden ser declaradas en el scope global")
	}

	funcName := ctx.ID().GetText()
	params := v.Visit(ctx.Param_list()).([]*Param)
	returnType := value.IVOR_NIL

	if ctx.Primitive_type() != nil {
		returnType = ctx.Primitive_type().GetText()
	}

	body := ctx.AllStmt()

	function := &Function{ // pointer ?
		Name:       funcName,
		Param:      params,
		ReturnType: returnType,
		Body:       body,
		DeclScope:  v.ScopeTrace.CurrentScope,
	}

	v.ScopeTrace.AddFunction(funcName, function)

	return nil
}

func (v *ReplVisitor) VisitParamList(ctx *compiler.ParamListContext) interface{} {

	params := make([]*Param, 0)

	for _, param := range ctx.AllFunc_param() {
		params = append(params, v.Visit(param).(*Param))
	}

	return params
}

func (v *ReplVisitor) VisitFuncParam(ctx *compiler.FuncParamContext) interface{} {

	externName := ""

	if ctx.ID(0) != nil {
		externName = ctx.ID(0).GetText()
	}

	innerName := ctx.ID(1).GetText()

	passByReference := false

	if ctx.INOUT_KW() != nil {
		passByReference = true
	}

	// todo: Change primitive type to a more general type
	paramType := ctx.Primitive_type().GetText()

	return &Param{
		ExternName:      externName,
		InnerName:       innerName,
		PassByReference: passByReference,
		Type:            paramType,
	}

}
