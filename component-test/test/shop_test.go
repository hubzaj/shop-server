package test

import (
	"github.com/hubzaj/golang-component-test/component-test/client"
	"github.com/hubzaj/golang-component-test/component-test/config"
	"github.com/hubzaj/golang-component-test/pkg/runner"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestCookieSyncing(t *testing.T) {
	t.Parallel()

	cfg := config.CreateDefaultConfig()

	httpClient := client.NewHTTPClient(cfg)

	runner.StartShopWithConfig(cfg)

	t.Run("config test", func(test *testing.T) {
		// When
		response := httpClient.SendGetRequest(test, "albums")

		// Then
		require.Equal(test, http.StatusOK, response.Status)
	})
}
