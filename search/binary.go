package search

import (
//"fmt"
)

type Binary struct {
	A []int
	X int
}

func (b *Binary) Recursive(min, max int) int {

	if max < min {
		return -1
	}

	mid := min + (max-min)/2
	if b.A[mid] > b.X {
		return b.Recursive(min, mid-1)
	} else if b.A[mid] < b.X {
		return b.Recursive(mid+1, max)
	} else {
		return mid
	}

}

func (b *Binary) Iterative(min, max int) int {

	for min <= max {

		mid := min + (max-min)/2
		if b.A[mid] == b.X {
			return mid
		}

		if b.A[mid] < b.X {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return -1
}

// return index where to insert
func (b *Binary) SearchForInsertion(min, max int, f func(int) bool) int {

	for min < max {
		mid := min + (max-min)/2
		if f(mid) {
			max = mid
		} else {
			min = mid + 1
		}
	}

	return min
}

// // return index where to insert by recursive
// func (b *Binary) Recursive(min, max int) int {

// 	if min >= max {
// 		return min
// 	}

// 	mid := min + (max-min)/2
// 	if b.A[mid] >= b.X {
// 		return b.Recursive(min, mid)
// 	} else if b.A[mid] < b.X {
// 		return b.Recursive(mid+1, max)
// 	}

// 	return min
// }
