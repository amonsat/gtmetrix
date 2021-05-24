package gtmetrix

import (
	"fmt"
	"testing"
	"time"

	"github.com/amonsat/go-env"
	"github.com/stretchr/testify/assert"
)

func TestGTMetrix(t *testing.T) {
	apiKey := env.GetString("API_KEY", "")
	//apiKey = "123"
	c, err := NewClient(&Options{ApiKey: apiKey})
	if !assert.NoError(t, err) {
		return
	}

	if status, err := c.GetAccountStatus(); assert.NoError(t, err) {
		assert.Equal(t, status.Data.Id, apiKey)
	} else {
		return
	}

	start := time.Now()
	params := &TestRequestParams{Url: "https://golang.org", Report: ReportTypeLighthouseAndLegacy}
	if resp, err := c.StartTest(params); assert.NoError(t, err) {
		assert.Equal(t, resp.Data.Type, "test")
		assert.Equal(t, resp.Data.Attributes.Source, "api")

		fmt.Printf("test started - resp: %+v\n", resp)

		var reportLink string
		waitTime := 30 * time.Second

		for {
			testResp, err := c.GetTest(resp.Data.Id)
			if assert.NoError(t, err) && testResp.Data.Attributes.State == TestStateTypeCompleted {
				reportLink = testResp.Data.Links.Report
				fmt.Printf("test completed %+v\n", testResp)
				break
			}
			fmt.Printf("test status %v -> wait %v\n", testResp.Data.Attributes.State, waitTime)
			time.Sleep(waitTime)
		}

		if resp, err := c.GetReportByLink(reportLink); assert.NoError(t, err) {
			fmt.Printf("report [%v]: %+v\n", time.Since(start), resp.Data)
		}
	}

}
