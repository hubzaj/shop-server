package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hubzaj/golang-component-test/pkg/shop"
	"net/http"
)

func GetAlbums(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			handlePanic(c, err)
		}
	}()
	processGetAlbumRequest(c)
}

func processGetAlbumRequest(c *gin.Context) {
	albums := []shop.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func handlePanic(c *gin.Context, err interface{}) {
	c.AbortWithStatus(http.StatusInternalServerError)
}
