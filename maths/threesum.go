package maths

type threesum struct {
	nums []int
}

// 3 sum using 2 sum to make it O(n^2)
func (ts *threesum) SumByTwoSum() map[int]int {

	output := make(map[int]int)

	twoSum := make(map[int]int)
	for _, vi := range ts.nums {
		for _, vj := range ts.nums {
			twoSum[vi+vj] = vj
		}
	}

	for _, v := range ts.nums {
		if s2, ok := twoSum[-(v)]; ok {
			output[v] = s2
		}

	}

	return output
}
