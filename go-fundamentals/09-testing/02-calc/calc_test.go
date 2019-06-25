package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var result int
	result = Add(14, 6)
	if result != 20 {
		t.Error("Se esperaba un 20, no un ", result)
	}
}
func TestSubs(t *testing.T) {
	var result int
	result = Subs(10, 3)
	if result != 7 {
		t.Error("Se esperaba un 7, no un ", result)
	}
}
func TestMult(t *testing.T) {
	var result int
	result = Mult(11, 6)
	if result != 66 {
		t.Error("Se esperaba un 66, no un ", result)
	}
}
func TestDiv(t *testing.T) {
	var result int
	result = Div(10, 2)
	if result != 5 {
		t.Error("Se esperaba un 5, no un ", result)
	}
}
