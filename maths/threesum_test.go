package maths

import (
	"fmt"
	"testing"
)

func TestSumByTwoSum(t *testing.T) {
	ts := threesum{[]int{-25, -10, -7, -3, 2, 4, 8, 10}}
	nums := ts.SumByTwoSum()

	for k, v := range nums {
		fmt.Println(k, v, -(v + k))
	}

	numsLen := len(nums)
	if numsLen != 6 {
		t.Errorf("Expected: 6; but got %d", numsLen)
	}
}
