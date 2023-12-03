package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hubzaj/golang-component-test/pkg/config"
	controllers "github.com/hubzaj/golang-component-test/pkg/controllers/shop"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	NewHandler(router)
	return router
}

func NewHandler(router *gin.Engine) {
	shopRouterGroup := router.Group(config.Config.Shop.Api.BaseURL)

	addAlbumRoutes(shopRouterGroup)

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
}

func addAlbumRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/albums", controllers.GetAlbums)
}
