package arrayFunctions

import (
	"testing"
)

func TestFrequency(t *testing.T) {

	var a1 = []int{10, 12, 12, 11, 10}
	result1 := Frequency(a1)
	if result1[12] != 2 && result1[10] != 2 && result1[11] != 11 {
		t.Errorf("Result not correct ", result1[12], result1[11], result1[10])
	}

}

func TestSum(t *testing.T) {

	var a1 = []int{10, 12, 12, 11, 10}
	result := Sum(a1)
	if result != 55 {
		t.Errorf("want %d, got %d ", 55, result)
	}

}
