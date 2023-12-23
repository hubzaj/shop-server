package shopstub

import (
	"github.com/hubzaj/golang-component-test/component-test/client"
	"github.com/hubzaj/golang-component-test/component-test/endpoint"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

type HealthEndpoints struct {
	httpClient *client.HTTPClient
}

func (stub *HealthEndpoints) GetHealthStatus(t *testing.T) (int, string) {
	response := stub.httpClient.SendGetRequest(t, endpoint.GetHealthStatus)
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	require.NoError(t, err)
	return response.StatusCode, string(responseBody)
}
