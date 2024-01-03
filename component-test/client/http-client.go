package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	config "github.com/hubzaj/golang-component-test/component-test/config"
	"github.com/stretchr/testify/require"
	"testing"

	"net/http"
	"time"
)

type HTTPClient struct {
	client *http.Client

	cfg *config.TestConfig
}

func NewHTTPClient(cfg *config.TestConfig) *HTTPClient {
	httpClient := &HTTPClient{}
	httpClient.initHTTPClient()
	httpClient.cfg = cfg
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
	return response
}

func (c *HTTPClient) SendPostRequestWithHeaders(t *testing.T, endpoint string, payload interface{}, headers map[string]string) *http.Response {
	jsonPayload, err := json.Marshal(payload)
	require.NoError(t, err)
	request, err := http.NewRequest(http.MethodPost, c.getShopUrlWithEndpoint(endpoint), bytes.NewReader(jsonPayload))
	require.NoError(t, err)
	for header, value := range headers {
		request.Header.Set(header, value)
	}
	response, err := c.client.Do(request)
	require.NoError(t, err)
	return response
}

func (c *HTTPClient) SendPostRequest(t *testing.T, endpoint string, payload interface{}) *http.Response {
	return c.SendPostRequestWithHeaders(t, endpoint, payload, map[string]string{
		"Content-Type": "application/json",
	})
}

func (c *HTTPClient) getShopURL() string {
	return fmt.Sprintf("http://%s:%d/%s",
		c.cfg.ShopConfig.Shop.HTTPServer.Host,
		c.cfg.ShopConfig.Shop.HTTPServer.Port,
		c.cfg.ShopConfig.Shop.Api.BaseURL,
	)
}

func (c *HTTPClient) getShopUrlWithEndpoint(endpoint string) string {
	return fmt.Sprintf("%s/%s", c.getShopURL(), endpoint)
}
