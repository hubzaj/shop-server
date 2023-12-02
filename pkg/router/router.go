package router

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hubzaj/golang-component-test/pkg/controllers/shop"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	applyAlbumRoutes(router)

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	return router
}

func applyAlbumRoutes(router *gin.Engine) {
	router.Group("/")
	router.GET("albums", controllers.GetAlbums)
}
