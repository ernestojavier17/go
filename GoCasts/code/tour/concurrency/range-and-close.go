package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("c <- %v \n", x)
		c <- x
		x, y = y, x+y
		//time.Sleep(3)
	}
	close(c)
	fmt.Println("Closing channel")
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	time.Sleep(1)
/*	for i := range c {
		fmt.Println(i)
	}*/

	//Checking if a channel is closed
	_, ok := <- c //ok is false if there are no more values to receive AND the channel is closed.
	if !ok {
		fmt.Println("Channel is close and There aren't more values")
		//Receiving from a close channel returns the zero value immediately
		fmt.Println(<- c)
		//Sending to a close channel will panic
		//c <- 55
	} else {
		fmt.Println("Channel is open OR there are some values")
		//If channel is close this will panic
		c <- 55
	}
}