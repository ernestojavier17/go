package main

import (
	"fmt"
)

func main() {
	channel := make(chan int, 1)// we create a buffer channel of size 1, this is bidirectional channel. Which means that inside this function, we can put in the channel and pull from the channel

	channel <- 42 // put number 42
	//This code will run because we created a buffer channel and we are allowing to pass 1 value to the channel
	//Bill Kennedy said: try to stay away from buffered channels
	fmt.Println(<- channel)
}
