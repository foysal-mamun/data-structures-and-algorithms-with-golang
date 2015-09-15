package dynamicconnectivity

type wquickunion struct {
	id []int
	sz []int
}

func (w *wquickunion) Root(i int) int {

	for i != w.id[i] {
		i = w.id[i]
	}

	return i
}

func (w *wquickunion) Connected(p, q int) bool {
	return w.Root(p) == w.Root(q)
}

func (w *wquickunion) Union(p, q int) {
	rp := w.Root(p)
	rq := w.Root(q)

	if rp == rq {
		return
	}
	if w.sz[rp] < w.sz[rq] {
		w.id[rp] = rq
		w.sz[rq] += w.sz[rp]
	} else {
		w.id[rq] = rp
		w.sz[rp] += w.sz[rq]
	}

}
