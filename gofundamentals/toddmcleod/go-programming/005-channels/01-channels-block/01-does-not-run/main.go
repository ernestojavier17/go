package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)// makes a channel where a can put integers

	channel <- 42 // put number 42
	//This code does not run because channels blocks. If we run this code, we'll get this error: fatal error: all goroutines are asleep - deadlock!
	//We need to put the above code inside a goroutine, so when it block it will only block that goroutine.
	//When we send an receive
	fmt.Println(<- channel)
}
