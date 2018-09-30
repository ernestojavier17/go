package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range nums {
		if v % 2 == 0 {
			fmt.Printf("%v is even \n", v)
		} else {
			fmt.Printf("%v is odd \n", v)
		}
	}
}

func getOddNumbers(numbs []int) [10]int {
	odds := [10]int{}

	for i, v := range numbs {
		if v % 2 == 0 {
			//odds = append(odds, v)
			odds[i] = v

		}
	}

	return odds
}