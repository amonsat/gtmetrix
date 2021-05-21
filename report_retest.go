package gtmetrix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// RetestReport - Retest a report with the same page settings. This is semantically identical to starting
//a new test using the same parameters as the test that generated this report.
func (c *Client) RetestReport(reportID string) (*TestResponse, error) {
	endpoint := c.opt.ApiUrl + "/reports/" + reportID + "/retest"

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

	var data TestResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %v", err)
	}

	return &data, nil
}
