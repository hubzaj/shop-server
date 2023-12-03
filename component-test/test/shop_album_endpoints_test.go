package test

import (
	"github.com/hubzaj/golang-component-test/component-test/client"
	"github.com/hubzaj/golang-component-test/component-test/config"
	"github.com/hubzaj/golang-component-test/component-test/runner"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestShopAlbumEndpoints(t *testing.T) {
	t.Parallel()

	cfg := config.CreateDefaultConfig()

	httpClient := client.NewHTTPClient(cfg)

	runner.StartShop(cfg)

	t.Run("should return all available albums", func(test *testing.T) {
		test.Parallel()
		// When
		response := httpClient.SendGetRequest(test, "albums")

		// Then
		require.Equal(test, http.StatusOK, response.StatusCode)
	})
}
