package test

import (
	"github.com/hubzaj/golang-component-test/component-test/client"
	"github.com/hubzaj/golang-component-test/component-test/config"
	"github.com/hubzaj/golang-component-test/component-test/endpoint"
	"github.com/hubzaj/golang-component-test/component-test/runner"
	"github.com/hubzaj/golang-component-test/component-test/utils"
	"github.com/hubzaj/golang-component-test/pkg/shop/model"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestShopAlbumEndpoints(t *testing.T) {
	t.Parallel()

	cfg := config.CreateDefaultConfig()

	httpClient := client.NewHTTPClient(cfg)
	//s := stub.InitStubs(cfg)

	runner.StartShop(cfg)

	t.Run("should append new album into existing ones", func(test *testing.T) {
		test.Parallel()
		// Given
		album := createNewAlbum(test, httpClient)

		// When
		response := httpClient.SendGetRequest(test, endpoint.GetAvailableAlbums)
		defer response.Body.Close()

		// Then
		require.Equal(test, http.StatusOK, response.StatusCode)

		actualAlbums := utils.UnmarshalResponseBodyToArray(response.Body, []*model.Album{})

		actualAlbum := utils.FindAlbumByTitle(actualAlbums, album.Title)
		require.NotNil(test, actualAlbum)
		require.NotEmpty(test, actualAlbum.ID)
		actualAlbum.ID = ""
		require.Equal(test, album, actualAlbum)
	})
}

func createNewAlbum(t *testing.T, c *client.HTTPClient) *model.Album {
	album := &model.Album{
		Title:  utils.GenerateRandomString(10),
		Artist: utils.GenerateRandomString(10),
		Price:  77.77,
	}
	response := c.SendPostRequest(t, endpoint.CreateNewAlbum, album)
	defer response.Body.Close()
	require.Equal(t, http.StatusCreated, response.StatusCode)
	return album
}
