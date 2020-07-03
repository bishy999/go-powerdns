
# go-powerdns

go-powerdns is a Go client library for accessing the PowerDNS API.

You can view the PowerDNS Http API docs here: [https://doc.powerdns.com/authoritative/http-api/index.html/](https://doc.powerdns.com/authoritative/http-api/index.html)

You can view the client API docs by serving the docs from this repository : [http://localhost:6060/pkg/](http://localhost:6060/pkg/)
```go
 godoc -http :6060
```

## Status
[![Build Status](https://travis-ci.com/bishy999/go-powerdns.svg?branch=master)](https://travis-ci.com/go-powerdns)
[![Go Report Card](https://goreportcard.com/badge/github.com/bishy999/go-powerdns)](https://goreportcard.com/report/github.com/bishy999/go-powerdns)
[![GoDoc](https://godoc.org/github.com/bishy999/go-powerdns/pkg/powerdns?status.svg)](https://godoc.org/github.com/bishy999/go-powerdns/pkg/powerdns))
![GitHub Repo size](https://img.shields.io/github/repo-size/bishy999/go-powerdns)
[![GitHub Tag](https://img.shields.io/github/tag/bishy999/go-powerdns.svg)](https://github.com/bishy999/go-powerdns/releases/latest)
[![GitHub Activity](https://img.shields.io/github/commit-activity/m/bishy999/go-powerdns)](https://github.com/bishy999/go-powerdns)
[![GitHub Contributors](https://img.shields.io/github/contributors/bishy999/go-powerdns)](https://github.com/bishy999/go-powerdns)


## Usage (package)

### Download package
```go
 go get github.com/bishy999/go-powerdns
 ```


### Use package
```go 
import (
	"log"

	"github.com/bishy999/go-powerdns/pkg/powerdns"
)
```

### Authentication
You will need an api key with sufficent priviliges to perform actions against the powerdns api.

You can then create a new client to add and delete dns records. An example of a client is stored under the cmd directory in this repository

```go

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




```

## Usage (binary)

Download the client binary from the repository and compile it with version 

Go get will download from the master, as such when we download it give it the tag verison from the master

```go
go get -ldflags "-X main.version=v1.0.0 -X main.buildstamp=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'`)" github.com/bishy999/go-powerdns/cmd/powerdns-client

./powerdns-client add run powerdns-client.go add -domain=aws.xcl.ie -record=jbtest -ttl=3600 -ip=10.0.0.1

//./powerdns-client delete -domain=aws.xcl.ie -record=mjbtest
```


## Contributing

We love pull requests! Please see the [contribution guidelines](CONTRIBUTING.md).
