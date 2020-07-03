package powerdns

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

const (
	zoneEndpoint string = "/api/v1/servers/localhost/zones/"
)

// timeout for context
const (
	timeout = 30
)

// RRset structure with JSON API metadata
type RRset struct {
	Name       string     `json:"name,omitempty"`
	Type       string     `json:"type,omitempty"`
	TTL        int        `json:"ttl,omitempty"`
	Changetype string     `json:"changetype,omitempty"`
	Records    []Records  `json:"records,omitempty"`
	Comments   []Comments `json:"comments,omitempty"`
}

// Records structure with JSON API metadata
type Records struct {
	Content  string `json:"content,omitempty"`
	Disabled bool   `json:"disabled"`
	SetPTR   bool   `json:"set-ptr,omitempty"`
}

// Comments structure with JSON API metadata
type Comments struct {
	Content    string `json:"content,omitempty"`
	Account    string `json:"account,omitempty"`
	ModifiedAT int64  `json:"modified_at,omitempty"`
}

// RRsets structure with JSON API metadata
type RRsets struct {
	Sets []RRset `json:"rrsets,omitempty"`
}

// AddRecord add a record to a domain
func (c *ClientConn) AddRecord() error {

	log.Printf(" Adding DNS Record ")
	jsonData, err := json.MarshalIndent(c.records, "", "    ")
	if err != nil {
		return err
	}

	apiURL := generateAPIURL(c.input["url"], zoneEndpoint, c.input["domain"])

	req, err := http.NewRequest(http.MethodPatch, apiURL.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	//req.Header.Set(c.input["apikey"], c.input["apipasswd"])
	req.Header["X-API-Key"] = []string{c.input["apipasswd"]}

	data, err := httputil.DumpRequest(req, true)
	if err != nil {
		return err
	}
	log.Printf(" Http Request \n%s", string(data))

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return &Error{
			Status: resp.Status,
		}
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		body, _ := ioutil.ReadAll(resp.Body)

		return &Error{
			Status:  resp.Status,
			Message: string(body),
		}
	}

	defer resp.Body.Close()

	return err
}

func generateAPIURL(baseURL, path, domain string) *url.URL {

	u, err := url.Parse(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	u.Path = fmt.Sprintf("%s%s%s", path, domain, ".")
	u.Scheme = "http"

	return u
}
