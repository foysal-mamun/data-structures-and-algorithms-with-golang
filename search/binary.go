package search

import (
//"fmt"
)

type binary struct {
	a []int
	x int
}

func (b *binary) Recursive(min, max int) int {

	if max < min {
		return -1
	}

	mid := min + (max-min)/2
	if b.a[mid] > b.x {
		return b.Recursive(min, mid-1)
	} else if b.a[mid] < b.x {
		return b.Recursive(mid+1, max)
	} else {
		return mid
	}

}

func (b *binary) Iterative(min, max int) int {

	for min <= max {

		mid := min + (max-min)/2
		if b.a[mid] == b.x {
			return mid
		}

		if b.a[mid] < b.x {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return -1
}

// return index where to insert
func (b *binary) SearchForInsertion(min, max int, f func(int) bool) int {

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
// func (b *binary) Recursive(min, max int) int {

// 	if min >= max {
// 		return min
// 	}

// 	mid := min + (max-min)/2
// 	if b.a[mid] >= b.x {
// 		return b.Recursive(min, mid)
// 	} else if b.a[mid] < b.x {
// 		return b.Recursive(mid+1, max)
// 	}

// 	return min
// }
