package shopstub

import (
	"github.com/hubzaj/golang-component-test/component-test/client"
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
