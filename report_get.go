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
			Browser                  string  `json:"browser"`
			Location                 string  `json:"location"`
			Source                   string  `json:"source"`
			GtmetrixGrade            string  `json:"gtmetrix_grade"`
			CumulativeLayoutShift    float64 `json:"cumulative_layout_shift"`
			StructureScore           int     `json:"structure_score"`
			SpeedIndex               int     `json:"speed_index"`
			OnloadTime               int     `json:"onload_time"`
			RedirectDuration         int     `json:"redirect_duration"`
			FirstPaintTime           int     `json:"first_paint_time"`
			DomContentLoadedDuration int     `json:"dom_content_loaded_duration"`
			DomContentLoadedTime     int     `json:"dom_content_loaded_time"`
			DomInteractiveTime       int     `json:"dom_interactive_time"`
			PageRequests             int     `json:"page_requests"`
			PageBytes                int     `json:"page_bytes"`
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
			ConnectDuration          int     `json:"connect_duration"`
			//legacy
			PagespeedScore int `json:"pagespeed_score"`
			YslowScore     int `json:"yslow_score"`
		} `json:"attributes"`
		Links struct {
			OptimizedImages string `json:"optimized_images"`
			ReportPdf       string `json:"report_pdf"`
			Har             string `json:"har"`
			Lighthouse      string `json:"lighthouse"`
			ReportUrl       string `json:"report_url"`
			Screenshot      string `json:"screenshot"`
			//legacy
			Pagespeed      string `json:"pagespeed"`
			PagespeedFiles string `json:"pagespeed_files"`
			Yslow          string `json:"yslow"`
		} `json:"links"`
	} `json:"data"`
}

// GetReport - get report by reportID
func (c *Client) GetReport(id string) (*ReportResponce, error) {
	endpoint := c.opt.ApiUrl + "/reports/" + id
	return c.GetReportByLink(endpoint)
}

// GetReportByLink - get report by reportID
func (c *Client) GetReportByLink(link string) (*ReportResponce, error) {
	req, err := c.NewRequest(http.MethodGet, link, nil)
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
