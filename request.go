package gtmetrix

import (
	"io"
	"net/http"
)

func (c *Client) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/vnd.api+json")
	req.SetBasicAuth(c.opt.ApiKey, "")
	//spew.Dump(req)
	return req, nil
}
