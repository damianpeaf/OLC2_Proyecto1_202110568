package value

const (
	IVOR_INT              = "Int"
	IVOR_FLOAT            = "Float"
	IVOR_STRING           = "String"
	IVOR_BOOL             = "Bool"
	IVOR_CHARACTER        = "Character"
	IVOR_NIL              = "nil"
	IVOR_BUILTIN_FUNCTION = "builtin_function"
)

type IVOR interface {
	Value() interface{}
	Type() string
}
