package gtmetrix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// APIURL - GTMetrix api url
const APIURL = "https://gtmetrix.com/api/0.1"

// Client - GTMetrix client
type Client struct {
	email      string
	apiKey     string
	httpClient *http.Client
}

// Init GTMetrix client
func Init(email, apiKey string) *Client {
	return &Client{
		email:      email,
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

// StartTest send url to check with GTMetrix
// Example:
// curl --user user@example.com:e8ddc55d93eb0e8281b255ea236dcc4f \
// --form url=http://example.com --form x-metrix-adblock=0 \
// https://gtmetrix.com/api/0.1/test
func (c *Client) StartTest(checkURL string) (*TestResponse, error) {
	apiURL := APIURL + "/test"
	vals := url.Values{
		"url": {checkURL},
	}

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(vals.Encode()))
	if err != nil {
		return nil, fmt.Errorf("create request failed: %v", err)
	}

	req.SetBasicAuth(c.email, c.apiKey)
	// logrus.Info("req: ", req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		var data TestError
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, fmt.Errorf("unmarshal response failed: %v", err)
		}

		return nil, fmt.Errorf("request failed: %v, %v", resp.Status, data.Error)

	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var data TestResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %v", err)
	}

	return &data, nil
}

// GetResult of checkID in GTMetrix
func (c *Client) GetResult(checkID string) (*TestResult, error) {
	apiURL := APIURL + "/test/" + checkID

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %v", err)
	}

	req.SetBasicAuth(c.email, c.apiKey)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed: %v", resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var data TestResult
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %v", err)
	}

	return &data, nil
}

// GetAccountStatus - get gtmetrix account status
func (c *Client) GetAccountStatus() (*AccountStatus, error) {
	apiURL := APIURL + "/status"

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %v", err)
	}
	// spew.Dump(req)

	req.SetBasicAuth(c.email, c.apiKey)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status code: %v", resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var data AccountStatus
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %v", err)
	}

	return &data, nil
}
