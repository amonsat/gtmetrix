package gtmetrix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TestGetResponse struct {
	Data struct {
		Id         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Source   string `json:"source"`
			Created  int    `json:"created"`
			Location int    `json:"location"`
			Browser  int    `json:"browser"`
			Started  int    `json:"started"`
			Finished int    `json:"finished"`
			State    string `json:"state"`
		} `json:"attributes"`
		Links struct {
			Report string `json:"report"`
		} `json:"links"`
	} `json:"data"`
}

// GetTest - get test or report if ready
func (c *Client) GetTest(testID string) (*TestGetResponse, error) {
	endpoint := c.opt.ApiUrl + "/tests/" + testID

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

	if resp.StatusCode != 200 && resp.StatusCode != 303 {
		var data ErrorResponse
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, fmt.Errorf("unmarshal error response failed: %v, %v", resp.Status, err)
		}
		return nil, fmt.Errorf("request failed: %v, %v", resp.Status, data)
	}

	var data TestGetResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %v", err)
	}

	return &data, nil
}
