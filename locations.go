package gtmetrix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LocationsResponse struct {
	Data []struct {
		Type       string `json:"type"`
		Id         string `json:"id"`
		Attributes struct {
			Name     string   `json:"name"`
			Default  bool     `json:"default"`
			Browsers []string `json:"browsers"`
		} `json:"attributes"`
	} `json:"data"`
}

// GetLocations - Get a list of available test locations
func (c *Client) GetLocations() (*LocationsResponse, error) {
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

	var data LocationsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %v", err)
	}

	return &data, nil
}
