package gtmetrix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TestRequest struct {
	Data struct {
		Type       string            `json:"type"`
		Attributes TestRequestParams `json:"attributes"`
	} `json:"data"`
}

type TestRequestParams struct {
	//The URL of the page to test
	Url string `json:"url"`
	//Location ID
	Location int `json:"location,omitempty"`
	//Browser ID
	Browser int `json:"browser,omitempty"`
	//Report type
	//
	//    lighthouse for Lighthouse
	//    legacy for PageSpeed/YSlow
	//    lighthouse,legacy for both
	//    none for a metrics-only report
	//
	//This parameter will vary in credit costs.
	Report ReportType `json:"report,omitempty"`
	//Choose how long the report will be retained and accessible - This parameter may incur additional credit costs.
	Retention int `json:"retention,omitempty"`
	//Username for the test page HTTP access authentication
	//This is not the API authentication.
	HttpauthUsername string `json:"httpauth_username,omitempty"`
	//Password for the test page HTTP access authentication
	//This is not the API authentication.
	HttpauthPassword string `json:"httpauth_password,omitempty"`
	//Enable AdBlock (default: 0)  (0, 1)
	Adblock int `json:"adblock,omitempty"`
	//Specify cookies to supply with test page requests https://gtmetrix.com/faq.html#faq-cookies.
	Cookies string `json:"cookies,omitempty"`
	//Enable generation of video (default: 0) (0, 1)
	//This parameter incurs additional credit costs.
	Video int `json:"video,omitempty"`
	//Stop the test at window.onload instead of after the page has fully loaded
	//(ie. 2 seconds of network inactivity). (default: 0) (0, 1)
	StopOnload int `json:"stop_onload,omitempty"`
	//Throttle the connection. Speed measured in Kbps, latency in ms.
	//'down/up/latency' values in Kbps
	Throttle ThrottleType `json:"throttle,omitempty"`
	//Only load resources that match one of the URLs on this list.
	//This uses the same syntax as the web front end.
	AllowUrl string `json:"allow_url,omitempty"`
	//Prevent loading of resources that match one of the URLs on this list.
	//This occurs after the Only Allow URLs are applied.
	//This uses the same syntax as the web front end.
	BlockUrl string `json:"block_url,omitempty"`

	//GTMetrix Pro Plan

	//Use a custom DNS host and IP to run the test with. (host:ip_address)
	DNS string `json:"dns,omitempty"`
	//Simulate the display of your site on a variety of devices using a pre-selected combination
	//of Screen Resolutions, User Agents, and Device Pixel Ratios. (device Id)
	SimulateDevice DeviceType `json:"simulate_device,omitempty"`
	//Use a custom User Agent string
	//simulate_device overrides this parameter with preset values.
	UserAgent string `json:"user_agent,omitempty"`
	//Set the width of the viewport for the analysis. Also requires browser_height to be set.
	//simulate_device overrides this parameter with preset values. (pixels) (default: 1366)
	BrowserWidth int `json:"browser_width,omitempty"`
	//Set the height of the viewport for the analysis. Also requires browser_width to be set.
	//simulate_device overrides this parameter with preset values. (pixels) (default: 768)
	BrowserHeight int `json:"browser_height,omitempty"`
	//Set the device pixel ratio for the analysis. Decimals are allowed.
	//simulate_device overrides this parameter with preset values. (1 - 5)
	BrowserDppx int `json:"browser_dppx,omitempty"`
	//Swaps the width and height of the viewport for the analysis. (0, 1)
	BrowserRotate int `json:"browser_rotate,omitempty"`
}

type TestResponse struct {
	Data struct {
		Type       string `json:"type"`
		Id         string `json:"id"`
		Attributes struct {
			Source   string `json:"source"`
			Location int    `json:"location"`
			Browser  int    `json:"browser"`
			State    string `json:"state"`
			Created  int    `json:"created"`
		} `json:"attributes"`
	} `json:"data"`
	Meta struct {
		CreditsLeft float32 `json:"credits_left"`
		CreditsUsed float32 `json:"credits_used"`
	} `json:"meta"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

func newTestRequest(params *TestRequestParams) *TestRequest {
	r := &TestRequest{}
	r.Data.Type = "test"
	r.Data.Attributes = *params
	fmt.Println("req: ", r)
	return r
}

func (c *Client) StartTest(params *TestRequestParams) (*TestResponse, error) {
	endpoint := c.opt.ApiUrl + "/tests"

	jsonValue, err := json.Marshal(newTestRequest(params))
	if err != nil {
		return nil, fmt.Errorf("marshal json failed: %v", err)
	}

	req, err := c.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, fmt.Errorf("create request failed: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 202 {
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
