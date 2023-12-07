package test

import (
	"github.com/hubzaj/golang-component-test/component-test/config"
	"github.com/hubzaj/golang-component-test/component-test/runner"
	"github.com/hubzaj/golang-component-test/component-test/stub"
	"github.com/hubzaj/golang-component-test/component-test/stub/shopstub"
	"github.com/hubzaj/golang-component-test/component-test/utils"
	"github.com/hubzaj/golang-component-test/pkg/shop/model"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestShopAlbumEndpoints(t *testing.T) {
	t.Parallel()

	cfg := config.CreateDefaultConfig(ctx)

	stubs := stub.InitStubs(cfg)
	t.Cleanup(stubs.Cleanup)

	runner.StartShop(cfg)

	t.Run("should append new album into existing ones", func(test *testing.T) {
		test.Parallel()
		// Given
		album := createNewAlbum(test, stubs.ShopClient)

		// When
		actualStatusCode, actualAlbums := stubs.ShopClient.Album.GetAvailableAlbums(test)

		// Then
		require.Equal(test, http.StatusOK, actualStatusCode)

		actualAlbum := utils.FindAlbumByTitle(actualAlbums, album.Title)
		require.NotNil(test, actualAlbum)
		require.NotNil(test, actualAlbum.ID)
		actualAlbum.ID = nil
		require.Equal(test, album, actualAlbum)
	})
}

func createNewAlbum(t *testing.T, shopClient *shopstub.ShopClient) *model.Album {
	album := &model.Album{
		Title:  utils.GenerateRandomString(10),
		Artist: utils.GenerateRandomString(10),
		Price:  77.77,
	}
	actualStatusCode := shopClient.Album.CreateNewAlbum(t, album)
	require.Equal(t, http.StatusCreated, actualStatusCode)
	return album
}
