package repl

import (
	"testing"
)

func TestIsVector(t *testing.T) {

	matrix := map[string]bool{
		"[int]":   true,
		"int":     false,
		"[[int]]": false,
	}

	for k, v := range matrix {
		if IsVectorType(k) != v {
			t.Errorf("isVector(%s) != %t", k, v)
		}
	}

}
