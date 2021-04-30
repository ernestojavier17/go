package search

/* When you import code from the standard library, you only need to reference the name of the package, unlike when you import code from
outside of the standard library. The compiler will always look for the packages you import at the locations ref- erenced by
the GOROOT and GOPATH environment variables. */

import (
	"fmt"
	"log"
	"sync"
)

//A map is a reference type that you’re required to make in Go
var matcherMap = make(map[string]Matcher)

//Run performs the search logic
func Run (s string) {
	// Retrieve the list of feeds to search through
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	//Creates an unbuffered channel to receive match results.
	results := make(chan *Result)

	//Setup a wait group so we can process all the feeds
	var waitGroup sync.WaitGroup

	//Set the number of goroutines we need to wait for while they process the individual feeds
	waitGroup.Add(len(feeds))

	fmt.Println(s)
}
