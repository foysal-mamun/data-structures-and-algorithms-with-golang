package dynamicconnectivity

import (
	"testing"
)

func TestWQUPCroot(t *testing.T) {
	w := wqupc{[]int{0, 1, 9, 1, 4, 5, 6, 7, 8, 3}, []int{}}
	r := w.Root(9)
	if r != 1 {
		t.Errorf("Expected 3; Got %d", r)
	}
}

func TestWQUPCConnected(t *testing.T) {
	w := wqupc{[]int{0, 0, 0, 1, 1, 1, 5, 5, 5, 5}, make([]int, 10)}
	w.Union(5, 9)
	if !w.Connected(9, 1) {
		t.Errorf("Expected true; Got false %d", w.Root(9))
	}
}
