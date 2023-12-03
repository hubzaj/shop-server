package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hubzaj/golang-component-test/pkg/shop/model"
	"github.com/hubzaj/golang-component-test/pkg/utils"

	"net/http"
)

func GetAlbums(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			utils.HandlePanic(c, err)
		}
	}()
	processGetAlbumRequest(c)
}

func processGetAlbumRequest(c *gin.Context) {
	// TODO: Replace for database - postgres
	albums := []model.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	c.IndentedJSON(http.StatusOK, albums)
}
