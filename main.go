package main

import (
	"fmt"
	log "github.com/inconshreveable/log15"
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

	logger := log.New()

	user, _ := user.Current()
	username := flag.String("u", user.Username, "username, defaults to USERNAME from OS environment")
	password := flag.String("p", "", "password")
	url := flag.String("h", "http://localhost:8090/xml", "Vertec server URL")
	showVersion := flag.Bool("version", false, "print version")
	useToken := flag.Bool("token", false, "use modern token login provided by Vertec Cloud Server")

	flag.Parse()

	if *showVersion {
		fmt.Printf("haddock version: 0.0.1, vertec access lib: %s, build: %s\n", vertec.Version(), BUILDID)
		os.Exit(0)
	}

	fmt.Println("Hunderttausend heulende und jaulende Höllenhunde!")
	logger.Info("connection properties", "url", *url, "username", *username)

	var settings vertec.Settings
	settings.URL = *url

	if *useToken {
		err := vertec.Login(settings, *username, *password)
		if err != nil {
			logger.Warn("cannot authenticate with token access method, falling", err)
		} else {
			settings.Username = *username
			settings.Password = *password
		}
	}

	logger.Info("query", "query", SOLLZEITQUERY)

	response, err := vertec.Query(SOLLZEITQUERY, settings)
	if err != nil {
		logger.Error("no response on query", "error", err, "query", SOLLZEITQUERY)
	} else {
		logger.Info("response", "response", response)
	}

	logger.Info("ended gracefully")
}