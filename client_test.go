package gtmetrix

import (
	"fmt"
	"testing"

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

	params := &TestRequestParams{Url: "https://golang.org"}
	if resp, err := c.StartTest(params); assert.NoError(t, err) {
		assert.Equal(t, resp.Data.Type, "test")
		assert.Equal(t, resp.Data.Attributes.Source, "api")

		fmt.Printf("test started - resp: %+v\n", resp)

		if testResp, _, err := c.GetTest(resp.Data.Id); assert.NoError(t, err) {
			fmt.Println(testResp)
		}
	}
}
