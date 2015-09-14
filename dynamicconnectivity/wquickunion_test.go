package dynamicconnectivity

import (
	"testing"
)

func TestWQUroot(t *testing.T) {
	w := wquickunion{[]int{0, 1, 9, 1, 4, 5, 6, 7, 8, 3}, []int{}}
	r := w.Root(9)
	if r != 1 {
		t.Errorf("Expected 3; Got %d", r)
	}
}

func TestWQUConnected(t *testing.T) {
	w := wquickunion{[]int{0, 1, 9, 1, 4, 5, 6, 7, 8, 3}, make([]int, 10)}
	w.Union(5, 9)
	if !w.Connected(9, 5) {
		t.Errorf("Expected true; Got false")
	}
}
