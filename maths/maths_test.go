package maths

import "testing"

func TestFactorial(t *testing.T) {
	fact := Factorial(6)
	if fact != 720 {
		t.Errorf("Expect: 720; but got: ", fact)
	}
}

func TestSin(t *testing.T) {
	result := Sin(23)
	if result != 0.39054379164003017 {
		t.Errorf("expect: 0.39054379164003017; got: ", result)
	}
}

func TestFibonacci(t *testing.T) {
	fib := Fibonacci(12)
	if fib != 144 {
		t.Errorf("expect: 144; got: ", fib)
	}
}

func TestReverseInt(t *testing.T) {
	rev := ReverseInt(1234)
	if rev != 4321 {
		t.Errorf("expect: 4321; got: ", rev)
	}
}

func TestDicimalToOctal(t *testing.T) {
	o := DicimalToOctal(93)
	if o != 135 {
		t.Errorf("expect: 135; got: ", o)
	}
}

func TestCharacterToNumber(t *testing.T) {
	num := CharacterToNumber("1234")
	if num != 1234 {
		t.Errorf("expect: 1234; got: ", num)
	}
}

func TestSquareRoot(t *testing.T) {
	sr := SquareRoot(16)
	if sr != 4 {
		t.Errorf("Expected: 2; Got: ", sr)
	}
}

func TestSdivisor(t *testing.T) {
	d := Sdivisor(1013)
	if d != 1 {

		t.Errorf("Expected: 1; Got: ", d)
	}
}

func TestGcd(t *testing.T) {
	d := Gcd(18, 30)
	if d != 6 {
		t.Errorf("Expected: 6; Got: ", d)
	}
}
