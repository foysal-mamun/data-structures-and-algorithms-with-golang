package dynamicconnectivity

import (
	"testing"
)

func TestConnected(t *testing.T) {
	qf := quickfind{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	qf.Union(1, 2)
	qf.Union(3, 2)
	con := qf.Connected(1, 3)
	if !con {
		t.Errorf("Expected: true; but got %t", con)
	}
}
