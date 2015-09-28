package maths

import (
	"github.com/foysal-mamun/data-structures-and-algorithms-with-golang/search"
	"github.com/foysal-mamun/data-structures-and-algorithms-with-golang/sort"
)

type twosum struct {
	ids []int
}

func (ts *twosum) Count() int {

	cnt := 0
	for i, vi := range ts.ids {
		for _, vj := range ts.ids[i+1:] {
			if vi+vj == 0 {
				cnt += 1
			}
		}
	}

	return cnt
}

func (ts *twosum) CountFast() int {

	is := sort.Insertion{ts.ids}
	is.Sort()
	ts.ids = is.A

	cnt := 0

	for i, v := range ts.ids {
		bs := search.Binary{A: ts.ids, X: -v}

		index := bs.Recursive(0, len(bs.A)-1)

		if index > i {
			cnt += 1
		}
	}
	return cnt
}
