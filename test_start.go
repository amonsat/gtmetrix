package gtmetrix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ReportTypeLighthouse          = "lighthouse"
	ReportTypeLegacy              = "legacy"
	ReportTypeLighthouseAndLegacy = "lighthouse,legacy"
	ReportTypeNone                = "none"

	ThrottleTypeBroadbandFast = "20000/5000/25"
	ThrottleTypeBroadband     = "5000/1000/30"
	ThrottleTypeBroadbandSlow = "1500/384/50"
	ThrottleTypeLteMobile     = "15000/10000/100"
	ThrottleType3gMobile      = "1600/768/200"
	ThrottleType2gMobile      = "240/200/400"
	ThrottleType56kDialUp     = "50/30/125"

	//Devices

	//Phones

	//Apple iPhone X/XS/11/12/12 mini/12 Pro 	375x812 @ 3 DPR
	DeviceTypePhoneIphoneX = "iphone_x"
	//Apple iPhone XR 	414x896 @ 2 DPR
	DeviceTypePhoneIphoneXR = "iphone_xr"
	//Apple iPhone XS Max/11 Pro/11 Pro Max/12 Pro Max 	414x896 @ 3 DPR
	DeviceTypePhoneIphoneXsMax = "iphone_xs_max"
	//Apple iPhone 4/4S 	320x480 @ 2 DPR
	DeviceTypePhoneIphone4s = "iphone_4s"
	//Apple iPhone 5/5C/5S/SE (1st Gen) 	320x568 @ 2 DPR
	DeviceTypePhoneIphoneSE = "iphone_se"
	//Apple iPhone 6/6S/7/8 Plus 	414x736 @ 3 DPR
	DeviceTypePhoneIphone7plus = "iphone_7_plus"
	//Apple iPhone 6/6S/7/8/SE (2nd Gen) 	375x667 @ 2 DPR
	DeviceTypePhoneIphone7 = "iphone_7"
	//nexus_4 	Google Nexus 4 	384x640 @ 2 DPR
	DeviceTypePhoneNexus4 = "nexus_4"
	//nexus_5 	Google Nexus 5 	360x640 @ 3 DPR
	DeviceTypePhoneNexus5 = "nexus_5"
	//pixel 	Google Nexus 5X/Pixel/Pixel 2 	412x732 @ 2.625 DPR
	DeviceTypePhonePixel = "pixel"
	//pixel_xl 	Google Nexus 6/6P/Pixel XL/Pixel 2 XL 	412x732 @ 3.5 DPR
	DeviceTypePhonePixelXL = "pixel_xl"
	//pixel_3 	Google Pixel 3 	412x824 @ 2.625 DPR
	DeviceTypePhonePixel3 = "pixel_3"
	//pixel_3_xl 	Google Pixel 3 XL/3a XL 	412x847 @ 3.5 DPR
	DeviceTypePhonePixel3xl = "pixel_3_xl"
	//pixel_4 	Google Pixel 3a/4/4 XL 	412x869 @ 2.625 DPR
	DeviceTypePhonePixel4 = "pixel_4"
	//pixel_4a 	Google Pixel 4a/5 	412x892 @ 2.625 DPR
	DeviceTypePhonePixel4a = "pixel_4a"
	//lumia_520 	Nokia Lumia 520 	320x533 @ 1.5 DPR
	DeviceTypePhoneLumia520 = "lumia_520"
	//galaxy_note_3 	Samsung Galaxy Note 3 	360x640 @ 3 DPR
	DeviceTypePhoneGalaxyNote3 = "galaxy_note_3"
	//galaxy_note_5 	Samsung Galaxy Note 4/5 	412x732 @ 2.625 DPR
	DeviceTypePhoneGalaxyNote5 = "galaxy_note_5"
	//galaxy_note_8 	Samsung Galaxy Note 8/9 	412x846 @ 2.625 DPR
	DeviceTypePhoneGalaxyNote8 = "galaxy_note_8"
	//galaxy_note_10 	Samsung Galaxy Note 10/10+ 	412x869 @ 2.625 DPR
	DeviceTypePhoneGalaxyNote10 = "galaxy_note_10"
	//galaxy_note_20 	Samsung Galaxy Note 20/20 Ultra 	412x915 @ 2.625 DPR
	DeviceTypePhoneGalaxyNote20 = "galaxy_note_20"
	//galaxy_s5 	Samsung Galaxy S4/S5 	360x640 @ 3 DPR
	DeviceTypePhoneGalaxyS5 = "galaxy_s5"
	//galaxy_s7 	Samsung Galaxy S6/S7 	360x640 @ 4 DPR
	DeviceTypePhoneGalaxyS7 = "galaxy_s7"
	//galaxy_s8 	Samsung Galaxy S8/S8+/S9/S9+ 	360x740 @ 3 DPR
	DeviceTypePhoneGalaxyS8 = "galaxy_s8"
	//galaxy_s10 	Samsung Galaxy S10/S10+ 	360x760 @ 3 DPR
	DeviceTypePhoneGalaxyS10 = "galaxy_s10"
	//galaxy_s20 	Samsung Galaxy S20/S20+/S20 Ultra 	360x800 @ 3 DPR
	DeviceTypePhoneGalaxyS20 = "galaxy_s20"

	//tablets

	//ipad_2 	Apple iPad 2/Mini 	1024x768 @ 1 DPR
	DeviceTypeTabletIpad2 = "ipad_2"
	//ipad 	Apple iPad 3/4/Air/Air 2/2017 	1024x768 @ 2 DPR
	DeviceTypeTabletIpad = "ipad"
	//nexus_7 	Google Nexus 7 	960x600 @ 2 DPR
	DeviceTypeTabletNexus7 = "nexus_7"
	//nexus_10 	Google Nexus 10 	1280x800 @ 2 DPR
	DeviceTypeTabletNexus10 = "nexus_10"
	//galaxy_tab_a 	Samsung Galaxy Tab A 10.1 	960x600 @ 2 DPR
	DeviceTypeTabletGalaxyTabA = "galaxy_tab_a"
	//galaxy_tab_s3 	Samsung Galaxy Tab S3 	1024x768 @ 2 DPR
	DeviceTypeTabletGalaxyTabS3 = "galaxy_tab_s3"
	//galaxy_tab_s7 	Samsung Galaxy Tab S7/S7+ 	1280x800 @ 2 DPR
	DeviceTypeTabletGalaxyTabS7 = "galaxy_tab_s7"
	//galaxy_tab_4 	Samsung Galaxy Tab 4 	1280x800 @ 1 DPR
	DeviceTypeTabletGalaxyTab4 = "galaxy_tab_4"
)

type ReportType string
type ThrottleType string
type DeviceType string

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
		CreditsLeft int `json:"credits_left"`
		CreditsUsed int `json:"credits_used"`
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
