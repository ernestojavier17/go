package maps

func FindOdd(seq []int) int {
	// your code here
	var temp map[int]int = make(map[int]int)

	for _, v := range seq {
		temp[v] = temp[v] + 1
	}
	res := make([]int, 10)
	for _, v := range temp {
		if v % 2 != 0 {
			res = append(res, v)
		}
	}
	return 0
}
