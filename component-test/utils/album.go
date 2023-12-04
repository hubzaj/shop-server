package utils

import "github.com/hubzaj/golang-component-test/pkg/shop"

func FindAlbumByTitle(albums []*shop.Album, title string) *shop.Album {
	for _, album := range albums {
		if album.Title == title {
			return album
		}
	}
	return nil
}
