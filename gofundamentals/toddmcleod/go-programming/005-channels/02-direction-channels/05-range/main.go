package main

import (
	"fmt"
)

func main() {
	c := make(chan int) //bidirectional channel
	go func() {
		for i := 0; i <=5; i++ {
			c <- i
		}
		close(c)//if we don't close the channel will have a deadlock
	}()
	for v := range c {// Pull off a channel until the channel is closed. When there is not more values on the channel it leave the range loop
		fmt.Println(v)
	}
	fmt.Println("About to exit")
}

