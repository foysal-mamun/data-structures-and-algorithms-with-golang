// Copyright 2015 Foysal Mamun. :)

// Implements of swap solutions.
package swap

import ()

// add swap rarely use
// can error due to integer overflow, so need a chekcing if (a != b)
//
func ArithmeticSwap(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b

	return a, b
}

// simplest and most widely use method
// need extra memory
//
func TempVariableSwap(a, b int) (int, int) {
	temp := a
	a = b
	b = temp
	return a, b
}

// faster then Temp Variable swap
// less readable
// on modern CPU this can be skiwer then Temp variable swap
//
func XorSwap(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b

	return a, b
}

// Parallel Assignment
// most simple way
//
func ParallelSwap(a, b int) (int, int) {
	return b, a
}
