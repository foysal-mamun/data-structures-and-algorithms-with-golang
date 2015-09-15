package dynamicconnectivity

type wqupc struct {
	id []int
	sz []int
}

func (w *wqupc) Root(i int) int {

	for i != w.id[i] {
		w.id[i] = w.id[w.id[i]]
		i = w.id[i]
	}
	return i
}

func (w *wqupc) Connected(p, q int) bool {
	return w.Root(p) == w.Root(q)
}

func (w *wqupc) Union(p, q int) {
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
		w.sz[rp] += rq
	}
}
