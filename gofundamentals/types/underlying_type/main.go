package main

import (
	"fmt"
)

type hotdog int // we can create our own type in go

func main() {
	var a int = 42
	var b hotdog = 43

	fmt.Printf("value:%v type:%T \n", a, a)//Output int
	fmt.Printf("value:%v type:%T \n", b, b)//Output 43 main.hotdog
	//a = b Doesn't compile: cannot use b (type hotdog) as type int in assignment
	a = int(b) // but I can convert to the underlying type int. Conversion not casting
}
