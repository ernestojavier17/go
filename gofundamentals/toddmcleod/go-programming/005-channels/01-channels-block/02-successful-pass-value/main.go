package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)// makes a channel where a can put integers
	go func() {
		channel <- 42 // put number 42
	}()
	fmt.Println(<- channel)
}
