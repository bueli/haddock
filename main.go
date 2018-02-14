package haddock

import (
	"fmt"
	"log"
	"flag"
	"github.com/bueli/vertec"
	"os"
	"os/user"
)

const SOLLZEITQUERY = `<Query>
<Selection>
<objref>273953</objref>
<ocl>getsollzeit(encodedate(2017,1,1),encodedate(2017,1,31))</ocl>
</Selection>
</Query>`

const BUILDID = `manual build`

func main() {
 
	user, _ := user.Current()
	username := flag.String("u", user.Username, "username, defaults to USERNAME from OS environment")
	password := flag.String("p", "", "password")
	url := flag.String("h", "http://localhost:8090/xml", "Vertec server URL")
	showVersion := flag.Bool("version", true, "print version")

	flag.Parse()
	
	if *showVersion {
		fmt.Printf("haddock version: 0.0.1, vertec access lib: %s, build: %s\n", vertec.Version(), BUILDID)
		os.Exit(0)
	}
	
    fmt.Println("Hunderttausend heulende und jaulende Höllenhunde!")
	fmt.Println("Version:", vertec.Version())
	
	var settings vertec.Settings
	
	settings.Username = *username
	settings.Password = *password
	settings.URL = *url
	log.Printf("accessing %s as user %s", settings.URL, settings.Username)

	response, err := vertec.Query(SOLLZEITQUERY, settings)
	if err != nil {
		log.Fatal("no respnse on query", err)
	}
	log.Print("response:\n", response)

	log.Printf("formatted response is:\n%s", response)

	log.Printf("ended gracefully")
}