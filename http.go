package doh

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// exchangeHTTPS sends a given query to a given resolver using a DoH POST
// request as described in RFC 8484, and returns the response's body.
// Returns an error if there was an issue sending the request or reading the
// response body.
func exchangeHTTPS(q []byte, resolver string, c *http.Client) (a []byte, err error) {
	url := fmt.Sprintf("https://%s/dns-query", resolver)
	body := bytes.NewBuffer(q)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return
	}

	req.Header.Add("Accept", "application/dns-message")
	req.Header.Add("Content-Type", "application/dns-message")

	resp, err := c.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("HTTPS server returned with non-OK code %d", resp.StatusCode)
		return
	}

	return ioutil.ReadAll(resp.Body)
}
