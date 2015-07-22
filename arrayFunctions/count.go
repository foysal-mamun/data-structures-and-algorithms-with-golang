package arrayFunctions

func Frequency(list []int) map[int]int {

	var counts = make(map[int]int)
	for _, v := range list {
		counts[v] = counts[v] + 1
	}

	return counts
}

func Sum(list []int) (total int) {

	for _, v := range list {
		total = total + v
	}

	return
}
