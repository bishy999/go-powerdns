package powerdns_test

import (
	"log"
	"os"
	"testing"

	"github.com/bishy999/go-powerdns/pkg/powerdns"
)

func TestNewClientConnAdd(t *testing.T) {

	defer os.Unsetenv("PDNS_URL")
	defer os.Unsetenv("PDNS_APIKEY")
	defer os.Unsetenv("PDNS_APIPASSWD")

	// Simulate setting required env variables
	os.Setenv("PDNS_URL", "http://myenv:8081")
	os.Setenv("PDNS_APIKEY", "X-API-Key")
	os.Setenv("PDNS_APIPASSWD", "1234BadPasswd")

	// Simulate user input on cli
	os.Args = []string{"/fake/loc/main", "add", "-domain=example.org", "-record=mytest", "-ttl=3600", "-ip=10.0.0.1"}

	t.Run("NewClientConnAdd", func(t *testing.T) {
		pdns, err := powerdns.NewClientConn()
		if err != nil {
			t.Fatalf("Could not read response %v correctly", err)
		}
		log.Printf("Response: %v", pdns)
	})

}

func TestNewClientConnDelete(t *testing.T) {

	defer os.Unsetenv("PDNS_URL")
	defer os.Unsetenv("PDNS_APIKEY")
	defer os.Unsetenv("PDNS_APIPASSWD")

	// Simulate setting required env variables
	os.Setenv("PDNS_URL", "http://myenv:8081")
	os.Setenv("PDNS_APIKEY", "X-API-Key")
	os.Setenv("PDNS_APIPASSWD", "1234BadPasswd")

	// Simulate user input on cli
	os.Args = []string{"/fake/loc/main", "delete", "-domain=example.org", "-record=mytest"}

	t.Run("NewClientConnDelete", func(t *testing.T) {
		pdns, err := powerdns.NewClientConn()
		if err != nil {
			t.Fatalf("Could not read response %v correctly", err)
		}
		log.Printf("Response: %v", pdns)
	})

}
