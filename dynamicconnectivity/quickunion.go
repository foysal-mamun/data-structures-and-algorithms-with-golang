package dynamicconnectivity

type quickunion struct {
	id []int
}

func (qu *quickunion) root(i int) int {
	for i != qu.id[i] {
		i = qu.id[i]
	}

	return i
}

func (qu *quickunion) Connected(p, q int) bool {
	return qu.root(p) == qu.root(q)
}

func (qu *quickunion) Union(p, q int) {

	qu.id[qu.root(p)] = qu.root(q)
}
