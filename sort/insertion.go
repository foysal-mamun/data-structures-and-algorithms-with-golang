package sort

import (
	//"fmt"
	"github.com/foysal-mamun/data-structures-and-algorithms-with-golang/search"
)

type Insertion struct {
	A []int
}

func (ins *Insertion) Sort() {
	l := len(ins.A)
	for i := 1; i < l; i++ {

		if ins.A[i-1] > ins.A[i] {

			x := ins.A[i]
			b := search.Binary{A: ins.A[:i-1], X: x}
			index := b.SearchForInsertion(0, len(b.A)-1, func(i int) bool { return b.A[i] >= b.X })

			// remove, insert and concat
			ins.A = append(ins.A[:index], append([]int{x}, append(ins.A[:i], ins.A[i+1:]...)[index:]...)...)

		}

	}
}
