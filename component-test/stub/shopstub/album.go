package shopstub

import (
	"github.com/hubzaj/golang-component-test/component-test/client"
	"github.com/hubzaj/golang-component-test/component-test/endpoint"
	"github.com/hubzaj/golang-component-test/component-test/utils"
	"github.com/hubzaj/golang-component-test/pkg/shop/model"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

type AlbumEndpoints struct {
	httpClient *client.HTTPClient
}

func (stub *AlbumEndpoints) GetAvailableAlbums(t *testing.T) (int, []*model.Album) {
	response := stub.httpClient.SendGetRequest(t, endpoint.GetAvailableAlbums)
	defer response.Body.Close()
	albums := utils.UnmarshalArrayResponseBody(response.Body, []*model.Album{})
	return response.StatusCode, albums
}

func (stub *AlbumEndpoints) CreateNewAlbum(t *testing.T, body *model.Album) *model.Album {
	response := stub.httpClient.SendPostRequest(t, endpoint.CreateNewAlbum, body)
	defer response.Body.Close()
	require.Equal(t, http.StatusCreated, response.StatusCode)
	album := utils.UnmarshalResponseBody(response.Body, &model.Album{})
	require.NotNil(t, album)
	return album
}
