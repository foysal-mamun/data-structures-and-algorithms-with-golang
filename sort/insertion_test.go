package sort

import (
	"reflect"
	"testing"
)

var nums = []int{
	11, 12, 8, 5, 6, 1, 3, 20, -1,
}

func TestInsertionSort(t *testing.T) {

	i := Insertion{nums}
	i.Sort()
	if !reflect.DeepEqual(i.A, []int{-1, 1, 3, 5, 6, 8, 11, 12, 20}) {
		t.Errorf("sort not working:: %v", i.A)
	}
}
