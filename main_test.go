package haddock

import (
	"testing"
	"os"
	"log"
)

func TestSmokeMainVersion(t *testing.T) {
	os.Args[0] = "haddock"
	os.Args[1] = "--version"
 	main()
	log.Print("done")
}
