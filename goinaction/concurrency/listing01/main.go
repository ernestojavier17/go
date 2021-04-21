package main

import (
	"fmt"
	"runtime"
	"sync"
)

// This example demostrates the use of concurrency. Because we only use one logical processor the goroutines in this example won't
// run in parallel. Concurrency is not parallelism. Parallelism can only be achieved when multiple pieces o code are executing
// simultaneously against different physical processors.
func main() {
	// Allocate 1 logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	// wg is used to wait for the program to finish
	var wg sync.WaitGroup
	// Add a count of two, one for each goroutine
	wg.Add(2)

	fmt.Println("Start Goroutines")
	// Declare an anonymous function and create a goroutine
	go func() {
		// schedule the call to Done to tell main we are done.
		defer wg.Done()
		// diplay the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
			fmt.Println("Goroutine 1")
		}
	}()

	go func() {
		// schedule the call to Done to tell main we are done.
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
			fmt.Println("Goroutine 2")
		}
	}()

	fmt.Println("Waiting to finish")
	// Wait for the goroutines to finish. Without this the main function will terminate before the goroutines have a chance to run,
	// because the main thread does not wait by default.
	wg.Wait()
	// A WaitGroup is a counting semaphore that can be used to maintain a record of running goroutines.
	// When the value of a WaitGroup is greater than zero, the Wait method will block.
	// We set the value of WaitGroup to 2 for two goroutines.
	// To decrement the value of the WaitGroup and eventually release the main function, we made calls to Done within the scope of a defer statement

	fmt.Println("...Terminating Program")
}
