package client

import (
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hubzaj/golang-component-test/pkg/config"
	"github.com/stretchr/testify/require"
	"testing"

	"net/http"
	"time"
)

type HTTPClient struct {
	client  *http.Client
	shopURL string
}

func NewHTTPClient(cfg *config.GeneralConfig) *HTTPClient {
	httpClient := &HTTPClient{}
	httpClient.initHTTPClient()
	httpClient.setShopURL(cfg)
	return httpClient
}

func (c *HTTPClient) initHTTPClient() {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 3
	retryClient.RetryWaitMin = 500 * time.Millisecond
	retryClient.RetryWaitMax = 3 * time.Second
	c.client = retryClient.StandardClient()
}

func (c *HTTPClient) setShopURL(cfg *config.GeneralConfig) {
	c.shopURL = getShopURL(cfg)
}

func getShopURL(cfg *config.GeneralConfig) string {
	return fmt.Sprintf("%s:%d/%s", cfg.Shop.HTTPServer.Host, cfg.Shop.HTTPServer.Port, cfg.Shop.Api.BaseURL)
}

func (c *HTTPClient) SendGetRequest(t *testing.T, endpoint string) *http.Response {
	request, err := http.NewRequest(http.MethodGet, c.getShopUrlWithEndpoint(endpoint), http.NoBody)
	require.NoError(t, err)
	response, err := c.client.Do(request)
	require.NoError(t, err)
	return response
}

func (c *HTTPClient) getShopUrlWithEndpoint(endpoint string) string {
	return fmt.Sprintf("%s/%s", c.shopURL, endpoint)
}
