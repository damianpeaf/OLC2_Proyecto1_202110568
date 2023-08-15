package repl

import (
	"fmt"
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
	DclScan    bool
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

	variable, msg := v.ScopeTrace.AddVariable(varName, varType, varValue, isConst, false, ctx.GetStart())

	// Variable already exists
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
	}

	return nil
}

func (v *ReplVisitor) VisitValueDecl(ctx *compiler.ValueDeclContext) interface{} {

	isConst := isDeclConst(ctx.Var_type().GetText())
	varName := ctx.ID().GetText()
	varValue := v.Visit(ctx.Expr()).(value.IVOR)
	// todo: Change primitive type to a more general type
	varType := varValue.Type()

	variable, msg := v.ScopeTrace.AddVariable(varName, varType, varValue, isConst, false, ctx.GetStart())

	// Variable already exists
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
	}
	return nil
}

func (v *ReplVisitor) VisitTypeDecl(ctx *compiler.TypeDeclContext) interface{} {

	isConst := isDeclConst(ctx.Var_type().GetText())
	varName := ctx.ID().GetText()
	// todo: Change primitive type to a more general type
	varType := ctx.Primitive_type().GetText()

	if isConst {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Las constantes Deben tener un valor asignado")
		return nil
	}

	variable, msg := v.ScopeTrace.AddVariable(varName, varType, value.DefaultNilValue, isConst, true, ctx.GetStart())

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
	}

	return nil
}

func (v *ReplVisitor) VisitVectorDecl(ctx *compiler.VectorDeclContext) interface{} {

	isConst := isDeclConst(ctx.Var_type().GetText())
	varName := ctx.ID().GetText()
	// todo: Change primitive type to a more general type
	varType := value.IVOR_VECTOR
	varAuxType := ctx.Primitive_type().GetText()

	varValue := v.Visit(ctx.Vector_expr()).(value.IVOR)

	variable, msg := v.ScopeTrace.AddVector(varName, varType, varAuxType, varValue, isConst, false, ctx.GetStart()) // ? ALLOW NIL

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
	}

	return nil
}

func (v *ReplVisitor) VisitVectorItemList(ctx *compiler.VectorItemListContext) interface{} {

	var vectorItems []value.IVOR

	for _, item := range ctx.AllExpr() {
		vectorItems = append(vectorItems, v.Visit(item).(value.IVOR))
	}

	var itemType = value.IVOR_NIL

	if ctx.Expr(0) != nil {
		itemType = vectorItems[0].Type()

		for _, item := range vectorItems {
			if item.Type() != itemType {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), "Todos los items del vector deben ser del mismo tipo")
				return value.DefaultNilValue
			}
		}
	}

	return NewVectorValue(vectorItems, itemType)
}

func (v *ReplVisitor) VisitVectoReferece(ctx *compiler.VectoRefereceContext) interface{} {

	varName := ctx.Id_pattern().GetText()

	variable := v.ScopeTrace.GetVariable(varName)

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable "+varName+" no encontrada")
		return value.DefaultNilValue
	}

	if variable.Type != value.IVOR_VECTOR {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La variable "+varName+" no es un vector")
		return value.DefaultNilValue
	}

	// copy vector
	return variable.Value.Copy()
}

func (v *ReplVisitor) VisitDirectAssign(ctx *compiler.DirectAssignContext) interface{} {

	varName := v.Visit(ctx.Id_pattern()).(string)
	varValue := v.Visit(ctx.Expr()).(value.IVOR)

	variable := v.ScopeTrace.GetVariable(varName)

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable "+varName+" no encontrada")
	} else {
		ok, msg := variable.Assign(varValue)

		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		}
	}

	return nil

}

func (v *ReplVisitor) VisitArithmeticAssign(ctx *compiler.ArithmeticAssignContext) interface{} {
	varName := v.Visit(ctx.Id_pattern()).(string)

	variable := v.ScopeTrace.GetVariable(varName)

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable "+varName+" no encontrada")
	} else {

		leftValue := variable.Value
		rightValue := v.Visit(ctx.Expr()).(value.IVOR)

		op := string(ctx.GetOp().GetText()[0])

		strat, ok := BinaryStrats[op]

		if !ok {
			log.Fatal("Binary operator not found")
		}

		ok, msg, varValue := strat.Validate(leftValue, rightValue)

		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
			return nil
		}

		ok, msg = variable.Assign(varValue)

		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		}
	}

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
	stringVal := ctx.GetText()[1 : len(ctx.GetText())-1]

	// Todo: scape sequences

	// Character literal
	if len(stringVal) == 1 {
		return &value.CharacterValue{
			InternalValue: stringVal,
		}
	}

	// String literal
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
	varName := ctx.Id_pattern().GetText()

	variable := v.ScopeTrace.GetVariable(varName)

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable "+varName+" no encontrada")
		return value.DefaultNilValue
	}

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
		v.ErrorTable.NewSemanticError(ctx.GetOp(), msg)
		return value.DefaultNilValue
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
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condicion del if debe ser un booleano")
		return false

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

			return // break
		}
	}()

	visited := false

	// evaluate cases
	for _, switchCase := range ctx.AllSwitch_case() {

		caseValue := v.GetCaseValue(switchCase)

		// ? use binary strat
		if caseValue.Type() != mainValue.Type() {
			// warning
			continue
		}

		if caseValue.Value() == mainValue.Value() {
			v.Visit(switchCase)
			visited = true
			break // implicit break
		}

	}

	// evaluate default
	if ctx.Default_case() != nil && !visited {
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
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condicion del ciclo debe ser un booleano")
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
				condition = v.Visit(ctx.Expr()).(value.IVOR)             // update condition
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
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condicion del ciclo debe ser un booleano")
			return
		}

		// reset scope
		whileScope.Reset()
	}
}

func (v *ReplVisitor) VisitForStmt(ctx *compiler.ForStmtContext) interface{} {

	varName := ctx.ID().GetText()
	var iterableItem *VectorValue = DefaultEmptyVectorValue

	if ctx.Range_() != nil {
		iterableItem = v.Visit(ctx.Range_()).(*VectorValue)
	}

	if ctx.Expr() != nil {
		iterableValue := v.Visit(ctx.Expr()).(value.IVOR)

		if iterableValue.Type() == value.IVOR_VECTOR {
			iterableItem = iterableValue.(*VectorValue)
		} else if iterableValue.Type() == value.IVOR_STRING {
			iterableItem = StringToVector(iterableValue.(*value.StringValue))
		} else {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "El valor del rango debe ser un vector o una cadena")
			return nil
		}
	}

	if iterableItem.Size() == 0 {
		return nil
	}

	// Push scope outer scope
	outerForScope := v.ScopeTrace.PushScope("outer_for")

	// create the associated variable to the iterable
	iterableVariable, msg := outerForScope.AddVariable(varName, iterableItem.ItemType, iterableItem.Current(), true, false, ctx.ID().GetSymbol())

	if iterableVariable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		log.Fatal("This should not happen")
		return nil
	}

	// Push forItem to call stack [breakable, continuable]

	forItem := &CallStackItem{
		ReturnValue: value.DefaultNilValue,
		Type: []string{
			BreakItem,
			ContinueItem,
		},
	}

	v.CallStack.Push(forItem)

	// Push inner for scope
	innerForScope := v.ScopeTrace.PushScope("inner_for")

	v.VisitInnerFor(ctx, outerForScope, innerForScope, forItem, iterableItem, iterableVariable)

	v.ScopeTrace.PopScope()    // pop inner for scope
	v.ScopeTrace.PopScope()    // pop outer for scope
	v.CallStack.Clean(forItem) // ? clean item if it's still in call stack

	return nil
}

func (v *ReplVisitor) VisitInnerFor(ctx *compiler.ForStmtContext, outerForScope *BaseScope, innerForScope *BaseScope, forItem *CallStackItem, iterableItem *VectorValue, iterableVariable *Variable) {

	// reset scope
	innerForScope.Reset()

	// handle break and continue statements from call stack
	defer func() {

		if item, ok := recover().(*CallStackItem); item != nil && ok {

			// Not a for item, propagate panic
			if item != forItem {
				panic(item)
			}

			// Continue
			if item.IsAction(ContinueItem) {
				item.ResetAction()                                                                          // reset action, can be used again
				iterableItem.Next()                                                                         // next item
				v.VisitInnerFor(ctx, outerForScope, innerForScope, forItem, iterableItem, iterableVariable) // continue
			}

			// Break
			if item.IsAction(BreakItem) {
				return
			}

		}
	}()

	// iterableItem.Size()
	for iterableItem.CurrentIndex < iterableItem.Size() {

		// update variable value
		iterableVariable.Value = iterableItem.Current()

		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}

		iterableItem.Next()
	}
}

func (v *ReplVisitor) VisitNumericRange(ctx *compiler.NumericRangeContext) interface{} {

	leftExpr := v.Visit(ctx.Expr(0)).(value.IVOR)
	rightExpr := v.Visit(ctx.Expr(1)).(value.IVOR)

	if leftExpr.Type() != value.IVOR_INT || rightExpr.Type() != value.IVOR_INT {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Los valores de los rangos deben ser enteros")
		return value.DefaultNilValue
	}

	left := leftExpr.(*value.IntValue).InternalValue
	right := rightExpr.(*value.IntValue).InternalValue

	if left > right {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El valor izquierdo del rango debe ser menor o igual al valor derecho")
	}

	var values []value.IVOR

	for i := left; i <= right; i++ {
		values = append(values, &value.IntValue{
			InternalValue: i,
		})
	}

	return &VectorValue{
		InternalValue: values,
		CurrentIndex:  0,
		ItemType:      value.IVOR_INT,
	}
}

func (v *ReplVisitor) VisitGuardStmt(ctx *compiler.GuardStmtContext) interface{} {

	condition := v.Visit(ctx.Expr()).(value.IVOR)

	if condition.Type() != value.IVOR_BOOL {
		fmt.Println(condition)
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condicion del guard debe ser un booleano")
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

	funcObj, msg := v.ScopeTrace.GetFunction(funcName)

	if funcObj == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		return value.DefaultNilValue
	}

	args := make([]*Argument, 0)

	if ctx.Arg_list() != nil {
		args = v.Visit(ctx.Arg_list()).([]*Argument)
	}

	switch funcObj := funcObj.(type) {
	case *BuiltInFunction:
		returnValue, ok, msg := funcObj.Exec(v.GetReplContext(), args)

		if !ok {

			if msg != "" {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
			}

			return value.DefaultNilValue

		}

		return returnValue

	case *Function:
		funcObj.Exec(v, args, ctx.GetStart())
		return funcObj.ReturnValue

	case *ObjectBuiltInFunction:
		funcObj.Exec(v, args, ctx.GetStart())
		return funcObj.ReturnValue

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

	params := make([]*Param, 0)

	if ctx.Param_list() != nil {
		params = v.Visit(ctx.Param_list()).([]*Param)
	}

	if len(params) > 0 {

		baseParamType := params[0].ParamType()

		for _, param := range params {
			if param.ParamType() != baseParamType {
				v.ErrorTable.NewSemanticError(param.Token, "Todos los parametros de la funcion deben ser del mismo tipo")
				return nil
			}
		}
	}

	returnType := value.IVOR_NIL
	var returnTypeToken antlr.Token = nil

	if ctx.Primitive_type() != nil {
		returnType = ctx.Primitive_type().GetText()
		returnTypeToken = ctx.Primitive_type().GetStart()
	}

	body := ctx.AllStmt()

	function := &Function{ // pointer ?
		Name:            funcName,
		Param:           params,
		ReturnType:      returnType,
		Body:            body,
		DeclScope:       v.ScopeTrace.CurrentScope,
		ReturnTypeToken: returnTypeToken,
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
	innerName := ""

	// at least ID(0) is defined
	// only 1 ID defined
	if ctx.ID(1) == nil {
		// innerName : type
		// _ : type
		innerName = ctx.ID(0).GetText()
	} else {
		// externName innerName : type
		externName = ctx.ID(0).GetText()
		innerName = ctx.ID(1).GetText()
	}

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
		Token:           ctx.GetStart(),
	}

}
