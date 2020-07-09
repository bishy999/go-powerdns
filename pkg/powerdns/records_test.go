package powerdns_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bishy999/go-powerdns/pkg/powerdns"
)

func TestAddRecord(t *testing.T) {

	defer os.Unsetenv("PDNS_URL")
	defer os.Unsetenv("PDNS_APIKEY")
	defer os.Unsetenv("PDNS_APIPASSWD")

	// Simulate setting required env variables
	os.Setenv("PDNS_URL", "http://myenv:8081")
	os.Setenv("PDNS_APIKEY", "X-API-Key")
	os.Setenv("PDNS_APIPASSWD", "1234BadPasswd")

	// Simulate user input on cli
	os.Args = []string{"/fake/loc/main", "add", "-domain=example.org", "-record=mytest", "-ttl=3600", "-ip=10.0.0.1"}

	t.Run("AddRecord", func(t *testing.T) {

		srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			_, err := rw.Write([]byte("ok"))
			powerdns.LogErr(err)
		}))
		defer srv.Close()

		pdns := &powerdns.ClientConn{
			Client: *srv.Client(),
			Input:  map[string]string{"url": srv.URL},
		}

		err := pdns.UpdateARecord()
		if err != nil {
			log.Fatalf("error adding record: [ %v ]", err)
		}

		log.Printf("Response: %v", pdns)
	})

}