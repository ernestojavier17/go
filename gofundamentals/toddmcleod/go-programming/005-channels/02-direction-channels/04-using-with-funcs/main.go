package main

import (
	"fmt"
)

func main() {
	c := make(chan int) //bidirectional channel
	go foo(c)
	bar(c)
	fmt.Println("About to exit")
}
// Send: argument is directional channel that only sends
func foo(c chan<- int) {
	c <- 42
}
// Receive: directional channel that only receives
func bar(c <-chan int) {
	fmt.Println(<-c)
}
