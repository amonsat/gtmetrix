package gtmetrix

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ReportResponce struct {
	Data struct {
		Id         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			StructureScore           int     `json:"structure_score"`
			Source                   string  `json:"source"`
			SpeedIndex               int     `json:"speed_index"`
			OnloadTime               int     `json:"onload_time"`
			Browser                  int     `json:"browser"`
			RedirectDuration         int     `json:"redirect_duration"`
			FirstPaintTime           int     `json:"first_paint_time"`
			DomContentLoadedDuration int     `json:"dom_content_loaded_duration"`
			DomContentLoadedTime     int     `json:"dom_content_loaded_time"`
			DomInteractiveTime       int     `json:"dom_interactive_time"`
			PageRequests             int     `json:"page_requests"`
			PageBytes                int     `json:"page_bytes"`
			GtmetrixGrade            string  `json:"gtmetrix_grade"`
			Location                 int     `json:"location"`
			HtmlBytes                int     `json:"html_bytes"`
			FirstContentfulPaint     int     `json:"first_contentful_paint"`
			PerformanceScore         int     `json:"performance_score"`
			FullyLoadedTime          int     `json:"fully_loaded_time"`
			TotalBlockingTime        int     `json:"total_blocking_time"`
			LargestContentfulPaint   int     `json:"largest_contentful_paint"`
			TimeToInteractive        int     `json:"time_to_interactive"`
			TimeToFirstByte          int     `json:"time_to_first_byte"`
			RumSpeedIndex            int     `json:"rum_speed_index"`
			BackendDuration          int     `json:"backend_duration"`
			OnloadDuration           int     `json:"onload_duration"`
			CumulativeLayoutShift    float64 `json:"cumulative_layout_shift"`
			ConnectDuration          int     `json:"connect_duration"`
		} `json:"attributes"`
		Links struct {
			OptimizedImages string `json:"optimized_images"`
			ReportPdf       string `json:"report_pdf"`
			Har             string `json:"har"`
			Lighthouse      string `json:"lighthouse"`
			ReportUrl       string `json:"report_url"`
			Screenshot      string `json:"screenshot"`
		} `json:"links"`
	} `json:"data"`
}

// GetReport - get report by reportID
func (c *Client) GetReport(reportID string) (*ReportResponce, error) {
	endpoint := c.opt.ApiUrl + "/reports/" + reportID

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

	var data ReportResponce
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %v", err)
	}

	return &data, nil
}
