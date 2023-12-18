package test

import (
	"github.com/hubzaj/golang-component-test/component-test/client"
	"github.com/hubzaj/golang-component-test/component-test/config"
	"github.com/hubzaj/golang-component-test/component-test/runner"
	"github.com/hubzaj/golang-component-test/component-test/stub"
	"github.com/hubzaj/golang-component-test/component-test/stub/shopstub"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestHealthEndpoints(t *testing.T) {
	t.Parallel()

	cfg := config.CreateDefaultConfig(ctx)

	stubs := stub.InitStubs(cfg)
	t.Cleanup(stubs.Cleanup)

	runner.StartShop(cfg)

	t.Run("should respond to health", func(test *testing.T) {
		test.Parallel()
		// When
		actualStatusCode, actualResponse := stubs.ShopClient.Health.GetHealthStatus(test)

		// Then
		require.Equal(test, http.StatusOK, actualStatusCode)
		require.Equal(test, "Hi! I am alive!", actualResponse)
	})
}

func TestHealthEndpointsWithoutStorageConnection(t *testing.T) {
	t.Parallel()

	cfg := config.CreateDefaultConfig(ctx)

	stubs := stub.Stub{
		ShopClient: shopstub.NewShopClient(client.NewHTTPClient(cfg)),
	}

	runner.StartShop(cfg)

	t.Run("should respond to health", func(test *testing.T) {
		test.Parallel()
		// When
		actualStatusCode, actualResponse := stubs.ShopClient.Health.GetHealthStatus(test)

		// Then
		require.Equal(test, http.StatusOK, actualStatusCode)
		require.Equal(test, "Hi! I am alive!", actualResponse)
	})
}
