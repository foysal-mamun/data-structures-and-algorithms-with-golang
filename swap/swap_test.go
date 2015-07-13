package swap

import "testing"

func TestArithmeticSwap(t *testing.T) {
	a, b := ArithmeticSwap(1, 2)
	if a != 2 || b != 1 {
		t.Errorf("Expected:: 2 then 1, got ", a, b)
	}
}

func TestTempVariableSwap(t *testing.T) {
	a, b := TempVariableSwap(1, 2)

	if a != 2 || b != 1 {
		t.Errorf("Expected:: 2 then 1, got ", a, b)
	}

}

func TestXorSwap(t *testing.T) {
	a, b := XorSwap(1, 2)
	if a != 2 || b != 1 {
		t.Errorf("Expected:: 2 then 1, got ", a, b)
	}
}

func TestParallelSwap(t *testing.T) {
	a, b := ParallelSwap(1, 2)
	if a != 2 || b != 1 {
		t.Errorf("Expected:: 2 then 1, got ", a, b)
	}
}
