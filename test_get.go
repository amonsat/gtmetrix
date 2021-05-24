package gtmetrix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetTest - get test or report if ready
func (c *Client) GetTest(testID string) (*TestResponse, *ReportResponce, error) {
	endpoint := c.opt.ApiUrl + "/tests/" + testID

	req, err := c.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("create request failed: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("request failed: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		if resp.StatusCode == 303 {
			var data ReportResponce
			err = json.Unmarshal(body, &data)
			if err != nil {
				return nil, nil, fmt.Errorf("unmarshal response failed: %v", err)
			}
			return nil, &data, nil
		}
		var data ErrorResponse
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, nil, fmt.Errorf("unmarshal error response failed: %v, %v", resp.Status, err)
		}
		return nil, nil, fmt.Errorf("request failed: %v, %v", resp.Status, data)
	}

	var data TestResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, nil, fmt.Errorf("unmarshal response failed: %v", err)
	}

	return &data, nil, nil
}
