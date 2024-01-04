package test

import (
	"github.com/hubzaj/golang-component-test/component-test/config"
	"github.com/hubzaj/golang-component-test/component-test/runner"
	"github.com/hubzaj/golang-component-test/component-test/stub"
	"github.com/hubzaj/golang-component-test/component-test/utils"
	"github.com/hubzaj/golang-component-test/pkg/shop/controller"
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
		album := stubs.ShopClient.Album.CreateNewAlbum(t, &model.Album{
			Title:  utils.GenerateRandString(),
			Artist: utils.GenerateRandString(),
			Price:  utils.GenerateRandFloat(),
		},
			controller.PROTOBUF,
		)

		// When
		actualStatusCode, actualAlbums := stubs.ShopClient.Album.GetAvailableAlbums(test)

		// Then
		require.Equal(test, http.StatusOK, actualStatusCode)

		actualAlbum := utils.FindAlbumByTitle(actualAlbums, album.Title)
		require.Equal(test, album, actualAlbum)
	})
}
