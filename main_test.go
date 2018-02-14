package main

import (
	"testing"
	"os"
	"log"
)

func TestSmokeMainVersion(t *testing.T) {
	os.Args = make([]string, 2);
	os.Args[0] = "haddock"
	os.Args[1] = "--version"
 	main()
	log.Print("done")
}
