package main

import (

	"log"
	"os"

	"github.com/ernestojavier17/goinaction/code/chap2/search"

)

func init() {
	//Change the device for logging to stdout
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
