package main

import (
	"log"

	"github.com/bishy999/go-powerdns/pkg/powerdns"
)

var (
	version    string
	buildstamp string
)

func main() {

	log.Printf("Version    : %s\n", version)
	log.Printf("Build Time : %s\n", buildstamp)

	pdns, err := powerdns.NewClientConn()

	if err != nil {
		log.Fatalf("error creating client: [ %v ]", err)
	}

	err = pdns.AddRecord()
	if err != nil {
		log.Printf("error adding record: [ %v ]", err)
	}

}
