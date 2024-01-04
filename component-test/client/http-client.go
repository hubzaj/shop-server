package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hubzaj/golang-component-test/component-test/config"
	"github.com/hubzaj/golang-component-test/pkg/shop/controller"
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

func (c *HTTPClient) SendPostRequest(
	t *testing.T,
	endpoint string,
	body proto.Message,
	contentType controller.ContentType) *http.Response {
	switch contentType {
	case controller.JSON:
		return c.SendJSONPostRequest(t, endpoint, body)
	case controller.PROTOBUF:
		return c.SendProtoPostRequest(t, endpoint, body)
	default:
		panic("%s is not supported Content-Type")
	}
}

func (c *HTTPClient) SendJSONPostRequest(t *testing.T, endpoint string, body proto.Message) *http.Response {
	payload, err := json.Marshal(body)
	require.NoError(t, err)
	return c.sendPostRequest(t, endpoint, payload, controller.JSON)
}

func (c *HTTPClient) SendProtoPostRequest(t *testing.T, endpoint string, body proto.Message) *http.Response {
	payload, err := proto.Marshal(body)
	require.NoError(t, err)
	return c.sendPostRequest(t, endpoint, payload, controller.PROTOBUF)
}

func (c *HTTPClient) sendPostRequest(
	t *testing.T,
	endpoint string,
	payload []byte,
	contentType controller.ContentType) *http.Response {
	request, err := http.NewRequest(http.MethodPost, c.getShopUrlWithEndpoint(endpoint), bytes.NewReader(payload))
	require.NoError(t, err)
	request.Header.Set("Content-Type", string(contentType))
	response, err := c.client.Do(request)
	require.NoError(t, err)
	return response
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
