package main

func main() {
	c := make(chan <- int, 2) //directional channel, we can only put in this channel. This a receiver
	c <- 42
	c <- 43
	//fmt.Print(<-c) does not compile because invalid operation, we cannot read from
	go func() {
		//fmt.Print(<-c) won't compile either
	}()

}
