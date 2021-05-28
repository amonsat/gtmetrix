package gtmetrix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BrowsersResponse struct {
	Data []struct {
		Type       string `json:"type"`
		Id         string `json:"id"`
		Attributes struct {
			Name       string `json:"name"`
			Browser    string `json:"browser"`
			Platform   string `json:"platform"`
			Device     string `json:"device"`
			Dns        bool   `json:"dns"`
			Cookies    bool   `json:"cookies"`
			Adblock    bool   `json:"adblock"`
			HttpAuth   bool   `json:"http_auth"`
			Video      bool   `json:"video"`
			UserAgent  bool   `json:"user_agent"`
			Lighthouse bool   `json:"lighthouse"`
			Resolution bool   `json:"resolution"`
			Filtering  bool   `json:"filtering"`
			Throttle   bool   `json:"throttle"`
		} `json:"attributes"`
	} `json:"data"`
}

// GetBrowsers - Get a list of available browsers. This list is limited to the test browsers currently available to you.
func (c *Client) GetBrowsers() (*BrowsersResponse, error) {
	endpoint := c.opt.ApiUrl + "/locations"

	req, err := c.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var data ErrorResponse
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, fmt.Errorf("unmarshal error response failed: %v, %v", resp.Status, err)
		}
		return nil, fmt.Errorf("request failed: %v, %v", resp.Status, data)
	}

	var data BrowsersResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %v", err)
	}

	return &data, nil
}
