package search

import (
	"testing"
)

var nums = []int{
	1, 2, 3, 5, 6, 7, 8, 9,
}

func TestBinaryRecursive(t *testing.T) {

	b := binary{nums, 6}
	key := b.Recursive(0, len(b.a)-1)
	if key != 4 {
		t.Errorf("Expected: 4; but got %d", key)
	}
}

func TestBinaryIterative(t *testing.T) {

	b := binary{nums, 2}
	key := b.Iterative(0, len(b.a)-1)
	if key != 1 {
		t.Errorf("Expected: 1; but got %d", key)
	}
}

func TestBinarySearchForInsertion(t *testing.T) {

	b := binary{nums, 6}
	key := b.SearchForInsertion(0, len(b.a)-1, func(i int) bool { return b.a[i] >= b.x })
	if key != 4 {
		t.Errorf("Expected: 4; but got %d", key)
	}
}
