package gtmetrix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ReportDeleteResponce struct {
	Data struct {
		Id   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// DeleteReport - Delete a report and all of its resources.
func (c *Client) DeleteReport(reportID string) (*ReportDeleteResponce, error) {
	endpoint := c.opt.ApiUrl + "/reports/" + reportID

	req, err := c.NewRequest(http.MethodDelete, endpoint, nil)
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

	var data ReportDeleteResponce
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %v", err)
	}

	return &data, nil
}
