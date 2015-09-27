package search

import (
	"testing"
)

var nums = []int{
	1, 2, 3, 5, 6, 7, 8, 9,
}

func TestBinaryRecursive(t *testing.T) {

	b := binary{nums, 8}
	key := b.Recursive(0, len(b.a)-1)
	if key != 6 {
		t.Errorf("Expected: 6; but got %d", key)
	}
}

func TestBinaryIterative(t *testing.T) {

	b := binary{nums, 5}
	key := b.Iterative(0, len(b.a)-1)
	if key != 3 {
		t.Errorf("Expected: 3; but got %d", key)
	}
}

func TestBinaryForInsertionSort(t *testing.T) {

	b := binary{nums, 6}
	key := b.GoWay(0, len(b.a)-1, func(i int) bool { return b.a[i] >= b.x })
	if key != 4 {
		t.Errorf("Expected: 4; but got %d", key)
	}
}
