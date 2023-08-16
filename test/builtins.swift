var w = Int("10")
var x =  Int("10.001")
var x1 = Int(10.999999)
var y = Int("Q10.00")

var res = x + w + x1 // 30
print(res)

var w = Float("10")
var x = Float("10.001")
var y = Float("Q10.00")

var res = x + w // 20.001
print(res)

print( String(10) + iota(3.5)) //imprime 103.5000
print( String( true )) //true
cadena = String(true) + "->" + String(3.504) //
print(cadena); // imprime true->3.50400000