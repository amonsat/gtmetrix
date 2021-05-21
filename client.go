package gtmetrix

import (
	"fmt"
	"net/http"
	"time"
)

const (
	defaultApiUrl = "https://gtmetrix.com/api/2.0"
	httpTimeout   = 10 * time.Second
)

type Options struct {
	ApiUrl string
	ApiKey string
}

type Client struct {
	opt        *Options
	httpClient *http.Client
}

//NewClient - new gtmetrix client
func NewClient(opt *Options) (*Client, error) {
	opt.setDefaults()

	err := opt.checkErrors()
	if err != nil {
		return nil, err
	}

	return &Client{
		opt: opt,
		httpClient: &http.Client{
			Timeout: httpTimeout,
		},
	}, nil
}

func (opt *Options) setDefaults() {
	if opt.ApiUrl == "" {
		opt.ApiUrl = defaultApiUrl
	}
}

func (opt *Options) checkErrors() error {
	if opt.ApiKey == "" {
		return fmt.Errorf("api key must not be empty")
	}
	return nil
}
