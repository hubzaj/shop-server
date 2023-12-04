package album

import (
	"database/sql"
	"github.com/hubzaj/golang-component-test/pkg/shop/model"
	"github.com/hubzaj/golang-component-test/pkg/storage"
	"log"
)

type Album interface {
	RegisterNewAlbum(album *model.Album)
	GetAvailableAlbums() []*model.Album
}

type Service struct {
	Album

	Storage *storage.Storage
}

func (s *Service) RegisterNewAlbum(album *model.Album) {
	_, err := s.Storage.DB.Exec("INSERT INTO  albums(id,title,artist,price) VALUES($1,$2,$3,$4)",
		album.ID,
		album.Title,
		album.Artist,
		album.Price,
	)
	if err != nil {
		log.Fatalf("error inserting new album: %s", err)
	} else {
		log.Printf("new album has been added: %v", album)
	}
}

func (s *Service) GetAvailableAlbums() []*model.Album {
	rows, err := s.Storage.DB.Query("SELECT * FROM albums")
	if err != nil {
		log.Fatalf("error reading albums from DB: %s", err)
	}
	return createAlbumsFromRows(rows)

}

func createAlbumsFromRows(rows *sql.Rows) []*model.Album {
	return []*model.Album{}
}
