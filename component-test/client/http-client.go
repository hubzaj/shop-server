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
	client *http.Client

	shopConfig *config.GeneralConfig
}

func NewHTTPClient(cfg *config.GeneralConfig) *HTTPClient {
	httpClient := &HTTPClient{}
	httpClient.initHTTPClient()
	httpClient.shopConfig = cfg
	return httpClient
}

func (c *HTTPClient) initHTTPClient() {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 3
	retryClient.RetryWaitMin = 500 * time.Millisecond
	retryClient.RetryWaitMax = 3 * time.Second
	c.client = retryClient.StandardClient()
}

func (c *HTTPClient) SendGetRequest(t *testing.T, endpoint string) *http.Response {
	request, err := http.NewRequest(http.MethodGet, c.getShopUrlWithEndpoint(endpoint), http.NoBody)
	require.NoError(t, err)
	response, err := c.client.Do(request)
	require.NoError(t, err)
	defer response.Body.Close()
	return response
}

func (c *HTTPClient) getShopURL() string {
	return fmt.Sprintf("http://%s:%d/%s",
		c.shopConfig.Shop.HTTPServer.Host,
		c.shopConfig.Shop.HTTPServer.Port,
		c.shopConfig.Shop.Api.BaseURL,
	)
}

func (c *HTTPClient) getShopUrlWithEndpoint(endpoint string) string {
	return fmt.Sprintf("%s/%s", c.getShopURL(), endpoint)
}
