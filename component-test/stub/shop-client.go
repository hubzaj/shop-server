package stub

import (
	"github.com/hubzaj/golang-component-test/component-test/client"
	"github.com/hubzaj/golang-component-test/component-test/endpoint"
	"github.com/hubzaj/golang-component-test/component-test/utils"
	"github.com/hubzaj/golang-component-test/pkg/shop/model"
	"testing"
)

type ShopClient struct {
	Album *AlbumEndpoints
}

func NewShopClient(httpClient *client.HTTPClient) *ShopClient {
	return &ShopClient{
		Album: &AlbumEndpoints{
			httpClient: httpClient,
		},
	}
}

type AlbumEndpoints struct {
	httpClient *client.HTTPClient
}

func (stub *AlbumEndpoints) GetAvailableAlbums(t *testing.T) (int, []*model.Album) {
	response := stub.httpClient.SendGetRequest(t, endpoint.GetAvailableAlbums)
	defer response.Body.Close()
	albums := utils.UnmarshalResponseBodyToArray(response.Body, []*model.Album{})
	return response.StatusCode, albums
}

func (stub *AlbumEndpoints) CreateNewAlbum(t *testing.T, album *model.Album) int {
	response := stub.httpClient.SendPostRequest(t, endpoint.CreateNewAlbum, album)
	defer response.Body.Close()
	return response.StatusCode
}
