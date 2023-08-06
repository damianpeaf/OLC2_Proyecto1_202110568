package value

func IsPrimitiveType(t string) bool {
	switch t {
	case IVOR_BOOL, IVOR_INT, IVOR_FLOAT, IVOR_STRING, IVOR_NIL:
		return true
	default:
		return false
	}
}
