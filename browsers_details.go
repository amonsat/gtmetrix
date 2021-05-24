package gtmetrix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BrowserDetailsResponse struct {
	Data struct {
		Type       string `json:"type"`
		Id         string `json:"id"`
		Attributes struct {
			Dns        bool   `json:"dns"`
			Cookies    bool   `json:"cookies"`
			Adblock    bool   `json:"adblock"`
			HttpAuth   bool   `json:"http_auth"`
			Video      bool   `json:"video"`
			UserAgent  bool   `json:"user_agent"`
			Browser    string `json:"browser"`
			Name       string `json:"name"`
			Device     string `json:"device"`
			Lighthouse bool   `json:"lighthouse"`
			Resolution bool   `json:"resolution"`
			Filtering  bool   `json:"filtering"`
			Throttle   bool   `json:"throttle"`
			Platform   string `json:"platform"`
		} `json:"attributes"`
	} `json:"data"`
}

// GetBrowserDetails - Get details about a browser by browser ID. This list is limited to the test browsers currently available to you.
func (c *Client) GetBrowserDetails(id string) (*BrowserDetailsResponse, error) {
	endpoint := c.opt.ApiUrl + "/locations/" + id

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

	var data BrowserDetailsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %v", err)
	}

	return &data, nil
}
