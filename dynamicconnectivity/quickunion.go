package dynamicconnectivity

type quickunion struct {
	id []int
}

func (qu *quickunion) Root(i int) int {
	for i != qu.id[i] {
		i = qu.id[i]
	}

	return i
}

/**
 * Not perfect yet
 */
func (w *quickunion) Childs(p int) []int {
	itms := make([]int, len(w.id))
	itmsIndex := 0

	for p != w.id[p] {

		p = w.id[p]
		itms[itmsIndex] = p
		itmsIndex += 1
	}

	return itms
}

func (qu *quickunion) Connected(p, q int) bool {
	return qu.Root(p) == qu.Root(q)
}

func (qu *quickunion) Union(p, q int) {

	rp := qu.Root(p)
	rq := qu.Root(q)

	if rp == rq {
		return
	}
	qu.id[rp] = rq
}
