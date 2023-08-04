# Pending
1. Variable assignment
    1.1. Type checking
        1.1.1. Int literal can be assigned to a float variable, but not vice versa.
        1.1.2. String literal of length 1 can be assigned to a char variable.
    1.2 Constant can't be reassigned.
2. Scope related
    2.1 Variable can't be declared twice in the same scope.
    2.2 Search in parent scope if not found in current scope.
3. Semantic visitor
    3.1 validate transfer stmt on guard stmt

# Next Steps
1. if stmt          ✅
2. switch stmt      ✅
3. while stmt       ▶
4. for stmt         ❌
5. guard stmt       
6. callstack
    6.1. break
    6.2. continue
7. function
    7.1. builtins
    7.2. return


# Notes

stackable elements: switch, while, for, function