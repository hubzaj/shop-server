package test

import (
	"github.com/hubzaj/golang-component-test/component-test/client"
	"github.com/hubzaj/golang-component-test/component-test/config"
	"github.com/hubzaj/golang-component-test/component-test/runner"
	"github.com/hubzaj/golang-component-test/component-test/utils"
	"github.com/hubzaj/golang-component-test/pkg/shop"
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
		// Given
		// TODO: Create album via POST
		album := &shop.Album{
			Title:  "Blue Train",
			Artist: "John Coltrane",
			Price:  56.99,
		}

		// When
		response := httpClient.SendGetRequest(test, "albums")
		defer response.Body.Close()

		// Then
		require.Equal(test, http.StatusOK, response.StatusCode)

		actualAlbums := utils.UnmarshalResponseBodyToArray(response.Body, []*shop.Album{})

		actualAlbum := utils.FindAlbumByTitle(actualAlbums, album.Title)
		require.NotNil(test, actualAlbum)
		require.NotEmpty(test, actualAlbum.ID)
		actualAlbum.ID = ""
		require.Equal(test, album, actualAlbum)
	})
}
