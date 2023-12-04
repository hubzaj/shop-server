package utils

import (
	"github.com/hubzaj/golang-component-test/pkg/shop/model"
)

func FindAlbumByTitle(albums []*model.Album, title string) *model.Album {
	for _, album := range albums {
		if album.Title == title {
			return album
		}
	}
	return nil
}
