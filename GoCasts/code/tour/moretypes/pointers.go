package main

import "fmt"

func main() {
	i, j := 42, 2701

	//The & operator generates a pointer to its operand.
	//The * operator denotes the pointer's underlying value.
	p := &i		// point to i
	fmt.Println(*p) // read i through the pointer
	//This is known as "dereferencing" or "indirecting".
	*p = 21		// set i through the pointer. Now the value i has change because p is a reference to pointer
	fmt.Println(i)


	p = &j 		// point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j)
}
