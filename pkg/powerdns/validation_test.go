package powerdns_test

import (
	"log"
	"os"
	"testing"

	"github.com/bishy999/go-powerdns/pkg/powerdns"
)

func TestUserInputAdd(t *testing.T) {

	defer os.Unsetenv("PDNS_URL")
	defer os.Unsetenv("PDNS_APIKEY")
	defer os.Unsetenv("PDNS_APIPASSWD")

	// Simulate setting required env variables
	os.Setenv("PDNS_URL", "http://myenv:8081")
	os.Setenv("PDNS_APIKEY", "X-API-Key")
	os.Setenv("PDNS_APIPASSWD", "1234BadPasswd")

	// Simulate user input on cli
	os.Args = []string{"/fake/loc/main", "add", "-domain=example.org", "-record=mytest", "-ttl=3600", "-ip=10.0.0.1"}

	t.Run("UserInputAdd", func(t *testing.T) {

		result, err := powerdns.CheckUserInput()
		if err != nil {
			log.Fatalf("error with user inoput: [ %v ]", err)
		}
		if v, found := result["url"]; !found {
			log.Fatalf("Input missing: [ %v ]", v)
		}
		if v, found := result["domain"]; !found {
			log.Fatalf("Input missing: [ %v ]", v)
		}

		log.Printf("Response: %v", result)

	})

}

func TestUserInputDelete(t *testing.T) {

	defer os.Unsetenv("PDNS_URL")
	defer os.Unsetenv("PDNS_APIKEY")
	defer os.Unsetenv("PDNS_APIPASSWD")

	// Simulate setting required env variables
	os.Setenv("PDNS_URL", "http://myenv:8081")
	os.Setenv("PDNS_APIKEY", "X-API-Key")
	os.Setenv("PDNS_APIPASSWD", "1234BadPasswd")

	// Simulate user input on cli
	os.Args = []string{"/fake/loc/main", "delete", "-domain=example.org", "-record=mytest"}

	t.Run("UserInputDelete", func(t *testing.T) {

		result, err := powerdns.CheckUserInput()
		if err != nil {
			log.Fatalf("error with user inoput: [ %v ]", err)
		}
		if v, found := result["url"]; !found {
			log.Fatalf("Input missing: [ %v ]", v)
		}
		if v, found := result["domain"]; !found {
			log.Fatalf("Input missing: [ %v ]", v)
		}
		if v, found := result["record"]; !found {
			log.Fatalf("Input missing: [ %v ]", v)
		}

		log.Printf("Response: %v", result)

	})

}
