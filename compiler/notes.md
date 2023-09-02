# Pending
1. Variable assignment
    1.1. Type checking
        1.1.1. Int literal can be assigned to a float variable, but not vice versa.                                                 ✅
        1.1.2. String literal of length 1 can be assigned to a char variable. IMPLEMENT CHAR LITERAL AS STRING LITERAL OF LENGTH 1? ✅
        1.1.3. Simple types are assign by value. Complex types are assign by reference. ????
    1.2 Constant can't be reassigned.   ✅
2. Scope related
    2.1 Variable can't be declared twice in the same scope.     ✅
    2.2 Search in parent scope if not found in current scope.   ✅
3. Semantic visitor
    3.1 validate transfer stmt on guard stmt 🕐
4. Operators
    4.1 Null makes all result null      ✅
5. Switch
    5.1. Implicit break on case stmt    ✅
6. Funcs
    6.1. Check inner, outer, and positional args. ✅
    6.2. Reference IVOR ✅
7. Vectors
    7.1 empty [] implicit case ✅
8. Structs
    8.1. Dcl visit 🕐
    8.2. Pointers of structs ✅
    8.3. Reuse func grammar ✅
    8.4. a[0][0].prop ✅
    8.5. Auto Referece 🕐
    8.6. Mutual Reference 🕐
9. Lexical and Syntax errors

# Next Steps
1. if stmt              ✅
2. switch stmt          ✅
3. while stmt           ✅
4. for stmt             ✅
5. guard stmt           ✅
6. callstack            ✅
    6.1. break          ✅  
    6.2. continue       ✅
7. function
    7.1. builtins       ✅
    7.2. declaration    ✅
    7.3. call           ✅
    7.4. return         ✅
    ? pass by reference ✅
8. Complex types
    8.1. vector         ✅
        8.1.1. iterator ✅
        8.1.2. builtins ✅
        8.1.2. access   ✅
    8.2. matrix ✅
    8.3. struct ✅


# Improvements

1. Eval items on call stack instead of removing them ✅
2. validate it first on the semachecker visitor      

# Notes
stackable elements: switch, while, for, function




# Matrix Grammar Of Python
list:
    | '[' [star_named_expressions] ']' 


// Match one or more occurrences of e, separated by s. The generated parse tree does not include the separator. This is otherwise identical to (e (s e)*).
star_named_expressions: ','.star_named_expression+ [','] 

star_named_expression:
    | '*' bitwise_or 
    | named_expression

named_expression:
    | assignment_expression
    | expression !':=' ->list

--- java ---

ElementValueArrayInitializer:
    { [ElementValues] [,] }

ElementValues:
    ElementValue { , ElementValue }

