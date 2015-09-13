package dynamicconnectivity

type quickfind struct {
	id []int
}

func (qf *quickfind) Connected(p, q int) bool {
	return qf.id[p] == qf.id[q]
}

func (qf *quickfind) Union(p, q int) {

	for i, v := range qf.id {
		if v == qf.id[p] {
			qf.id[i] = qf.id[q]
		}
	}
}
