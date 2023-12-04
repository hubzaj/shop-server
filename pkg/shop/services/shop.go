package services

import (
	"github.com/hubzaj/golang-component-test/pkg/shop/services/album"
)

type ShopService struct {
	AlbumService *album.Service
}
