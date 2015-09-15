package dynamicconnectivity

import (
	"fmt"
	"testing"
)

func TestQuConnected(t *testing.T) {
	qu := quickunion{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	qu.Union(2, 4)
	qu.Union(9, 4)
	if !qu.Connected(2, 9) {
		t.Errorf("Expected: true; got false")
	}
}

func BenchmarkQuUnion(t *testing.B) {
	id := make([]int, 1000000)
	for i, v := range id {
		id[i] = v
	}

	qu := quickunion{id}
	for _, v := range id[:1000] {
		if v%2 == 0 {
			qu.Union(v, v+1)
		}
	}

}

/**
 * Not perfect yet
 */
func TestWQUChilds(t *testing.T) {
	id := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}

	w := quickunion{id}
	w.Union(0, 1)
	w.Union(0, 2)
	w.Union(0, 3)
	w.Union(0, 4)
	w.Union(0, 5)

	w.Union(1, 2)
	w.Union(2, 3)
	w.Union(3, 4)
	w.Union(4, 5)

	w.Union(4, 9)
	w.Union(9, 14)
	w.Union(14, 13)
	w.Union(13, 18)
	w.Union(18, 23)

	w.Union(21, 22)
	w.Union(22, 23)
	w.Union(23, 24)
	w.Union(24, 25)

	w.Union(26, 21)
	w.Union(26, 22)
	w.Union(26, 23)
	w.Union(26, 24)
	w.Union(26, 25)

	p, q := 0, 0

	fmt.Println("childs::")
	//for _, v := range w.id {
	for _, v := range w.Childs(p) {
		fmt.Print(v, ", ")
	}

	if !w.Connected(p, q) {
		t.Errorf("Expected true; Got false %d - %d", w.Root(p), w.Root(q))
	} else {
		fmt.Printf("%d - %d\n", w.Root(p), w.Root(q))
	}
}
