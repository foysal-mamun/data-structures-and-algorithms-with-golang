package maths

import (
	"fmt"
)

type threesum struct {
	nums []int
}

// 3 sum using 2 sum to make it O(n^2)
// // this given wrong result
func (ts *threesum) SumByTwoSum() map[int]int {

	output := make(map[int]int)

	twoSum := make(map[int]int)
	N := len(ts.nums)
	for i := 0; i < N; i += 1 {
		for j := i + 1; j < N; j += 1 {
			twoSum[ts.nums[i]+ts.nums[j]] = ts.nums[j]
		}
	}

	for _, v := range ts.nums {
		if s2, ok := twoSum[-(v)]; ok {
			output[v] = s2
		}

	}

	return output
}

func (ts *threesum) PrintAllThreeSum() {

	N := len(ts.nums)
	for i := 0; i < N; i += 1 {
		for j := i; j < N; j += 1 {
			for k := j; k < N; k += 1 {
				if ts.nums[i]+ts.nums[j]+ts.nums[k] == 0 {
					fmt.Println(ts.nums[i], ts.nums[j], ts.nums[k])
				}
			}
		}
	}
}

func (ts *threesum) TriplesCount() int {
	cnt := 0
	N := len(ts.nums)
	for i := 0; i < N; i += 1 {
		for j := i; j+1 < N; j += 1 {
			for k := j + 1; k < N; k += 1 {
				if ts.nums[i]+ts.nums[j]+ts.nums[k] == 0 {
					cnt += 1
				}
			}
		}
	}

	return cnt
}

func (ts *threesum) TwoSum() int {
	N := len(ts.nums)
	cnt := 0

	for i := 0; i < N; i += 1 {
		for j := i + 1; j < N; j += 1 {
			if ts.nums[i]+ts.nums[j] == 0 {
				cnt++
			}
		}
	}

	return cnt
}