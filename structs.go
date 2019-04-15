package gtmetrix

// TestResponse response of check
type TestResponse struct {
	TestID       string `json:"test_id,omitempty"`        // The test ID
	PollStateURL string `json:"poll_state_url,omitempty"` // URL to use to poll test state
	CreditsLeft  int    `json:"credits_left,omitempty"`   // The number of API credits remaining after running this test
}

// TestError response error
type TestError struct {
	Error string `json:"error,omitempty"`
}

// TestState type
type TestState string

// Status of check
const (
	StatusQueued    TestState = "queued"
	StatusStarted   TestState = "started"
	StatusCompleted TestState = "completed"
	StatusError     TestState = "error"
)

// TestResult result of check
type TestResult struct {
	State   TestState `json:"state,omitempty"`
	Error   string    `json:"error,omitempty"`
	Results struct {
		ReportURL                string `json:"report_url,omitempty"`
		PagespeedScore           int    `json:"pagespeed_score,omitempty"`
		YSlowScore               int    `json:"yslow_score,omitempty"`
		HTMLBytes                int    `json:"html_bytes,omitempty"`
		HTMLLoadTime             int    `json:"html_load_time,omitempty"`
		PageBytes                int    `json:"page_bytes,omitempty"`
		PageLoadTime             int    `json:"page_load_time,omitempty"`
		PageElements             int    `json:"page_elements,omitempty"`
		RedirectDuration         int    `json:"redirect_duration,omitempty"`
		ConnectDuration          int    `json:"connect_duration,omitempty"`
		BackendDuration          int    `json:"backend_duration,omitempty"`
		FirstPaintTime           int    `json:"first_paint_time,omitempty"`
		FirstContentfulPaintTime int    `json:"first_contentful_paint_time,omitempty"`
		DomInteractiveTime       int    `json:"dom_interactive_time,omitempty"`
		DomContentLoadedTime     int    `json:"dom_content_loaded_time,omitempty"`
		DomContentLoadedDuration int    `json:"dom_content_loaded_duration,omitempty"`
		OnloadTime               int    `json:"onload_time,omitempty"`
		OnloadDuration           int    `json:"onload_duration,omitempty"`
		FullyLoadedTime          int    `json:"fully_loaded_time,omitempty"`
		RumSpeedIndex            int    `json:"rum_speed_index,omitempty"`
	} `json:"results,omitempty"`

	Resources struct {
		Screenshot     string `json:"screenshot,omitempty"`
		HAR            string `json:"har,omitempty"`
		Pagespeed      string `json:"pagespeed,omitempty"`
		PagespeedFiles string `json:"pagespeed_files,omitempty"`
		YSlow          string `json:"yslow,omitempty"`
		ReportPDF      string `json:"report_pdf,omitempty"`
		ReportPDFFull  string `json:"report_pdf_full,omitempty"`
		Video          string `json:"video,omitempty"`
		Filmstrip      string `json:"filmstrip,omitempty"`
	} `json:"resources,omitempty"`
}

// AccountStatus - GTMetrix account status
type AccountStatus struct {
	// api_credits	Amount of API credits remaining	integer
	APICredits int `json:"api_credits,omitempty"`
	// api_refill	Unix timestamp for next API refill	integer
	APIRefill int `json:"api_refill,omitempty"`
}
