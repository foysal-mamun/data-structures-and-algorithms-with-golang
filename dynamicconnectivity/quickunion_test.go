package dynamicconnectivity

import (
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
