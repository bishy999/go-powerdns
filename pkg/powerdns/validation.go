package powerdns

import (
	"errors"
	"flag"
	"os"
)

const (
	// CreateUsage message identify what input is expected
	createUsage = `
	#####################################################################################################
	#                                                                                                   #
	#  Usage:                                                                                           #
	#      ./go run powerdns-client.go add -domain=aws.xcl.ie -record=mjbtest -ttl=3600 -ip=10.0.0.1    #    
	#                                                                                                   #
	#####################################################################################################
	`

	// DeleteUsgage message identify what input is expected
	deleteUsgage = `
	########################################################################
	#                                                                      #
	#                                                                      #
	#  Usage:                                                              #
	#      ./powerdns-client delete -name=dev99                            #
	#                                                                      #
	########################################################################
	`

	createArg = "add"
)

// CheckUserInput check cli input provided by user meets requirements and return input in map if it does
func CheckUserInput() (map[string]string, error) {

	if len(os.Args) < 2 {
		usage()
		msg := "add or delete sub command is required"
		return nil, errors.New(msg)
	}

	input := map[string]string{}
	input["url"] = os.Getenv("PDNS_URL")
	input["apikey"] = os.Getenv("PDNS_APIKEY")
	input["apipasswd"] = os.Getenv("PDNS_APIPASSWD")

	if input["url"] == "" || input["apikey"] == "" || input["apipasswd"] == "" {
		usage()
		msg := "PDNS_URL, PDNS_APIKEY & PDNS_APIPASSWD need to set as environmental variables"
		return nil, errors.New(msg)
	}

	createCommand := flag.NewFlagSet(createArg, flag.ExitOnError)
	domainPtr := createCommand.String("domain", "", "Domain where the record is to be modified e.g. example.org ")
	recordPtr := createCommand.String("record", "", " The record to add to the domain e.g. demo.example.org")
	ttlPtr := createCommand.String("ttl", "", "DNS ttl in seconds to set e.g. 3600")
	ipPtr := createCommand.String("ip", "", "The ip to assign to the record")

	switch os.Args[1] {
	case createArg:
		err := createCommand.Parse(os.Args[2:])
		LogErr(err)
	default:
		usage()
		os.Exit(1)
	}

	if createCommand.Parsed() {
		if *domainPtr == "" {
			createCommand.PrintDefaults()
			msg := "domain needs to be provided"
			return nil, errors.New(msg)
		}
		input["domain"] = *domainPtr
		if *recordPtr == "" {
			createCommand.PrintDefaults()
			msg := "name needs to be provided"
			return nil, errors.New(msg)
		}
		input["record"] = *recordPtr + "." + input["domain"] + "."
		if *ttlPtr == "" {
			createCommand.PrintDefaults()
			msg := "ttl needs to be provided"
			return nil, errors.New(msg)
		}
		input["ttl"] = *ttlPtr
		if *ipPtr == "" {
			createCommand.PrintDefaults()
			msg := "ip address needs to be provided"
			return nil, errors.New(msg)
		}
		input["ip"] = *ipPtr
	}
	return input, nil
}
