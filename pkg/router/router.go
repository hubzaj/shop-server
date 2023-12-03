package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hubzaj/golang-component-test/pkg/config"
	"github.com/hubzaj/golang-component-test/pkg/shop/controller"
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
	routerGroup.GET("/albums", controller.GetAlbums)
}
