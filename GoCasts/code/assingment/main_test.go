package main

import (
	"fmt"
	"testing"
)

func TestGetOddNumbers_notEmpty(t *testing.T) {

	all := []int{1,2,3,4,5,7}

	odds := [2]int{2, 4}
	res := getOddNumbers(all)

	if odds != res {

	}

	pair1 := [2]int {4, 2}
	pair2 := [2]int {2, 4}
	if pair1 != pair2 {
		fmt.Println("different pair")
	}
	t.Error()

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)
}
