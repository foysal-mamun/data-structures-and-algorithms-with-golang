package dynamicconnectivity

import (
	"fmt"
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

func TestWQUPCPercolation(t *testing.T) {
	id := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}
	sz := make([]int, 27)
	w := wqupc{id, sz}
	w.Union(1, 0)
	w.Union(2, 0)
	w.Union(3, 0)
	w.Union(4, 0)
	w.Union(5, 0)

	w.Union(21, 26)
	w.Union(22, 26)
	w.Union(23, 26)
	w.Union(24, 26)
	w.Union(25, 26)

	w.Union(4, 9)
	w.Union(14, 9)
	w.Union(13, 14)
	w.Union(13, 18)
	w.Union(18, 23)

	p, q := 26, 0

	if !w.Connected(p, q) {
		t.Errorf("Expected true; Got false %d - %d", w.Root(p), w.Root(q))
	} else {
		fmt.Printf("%d - %d\n", w.Root(p), w.Root(q))
	}
}
