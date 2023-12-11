package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hubzaj/golang-component-test/pkg/config"
	"github.com/hubzaj/golang-component-test/pkg/shop/controller"
	"github.com/hubzaj/golang-component-test/pkg/shop/controller/album"
	"github.com/hubzaj/golang-component-test/pkg/shop/controller/health"
	"net/http"
)

func InitRouter(dependencies *controller.Dependencies) *gin.Engine {
	router := gin.Default()
	NewHandler(router, dependencies)
	return router
}

func NewHandler(router *gin.Engine, dependencies *controller.Dependencies) {
	shopRouterGroup := router.Group(config.Config.Shop.Api.BaseURL)

	addHealthRoutes(shopRouterGroup)
	addAlbumRoutes(shopRouterGroup, dependencies)

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
}

func addHealthRoutes(routerGroup *gin.RouterGroup) {
	healthController := health.NewController()
	routerGroup.GET("/health", healthController.GetHealthStatus)
}

func addAlbumRoutes(routerGroup *gin.RouterGroup, dependencies *controller.Dependencies) {
	albumController := album.NewController(dependencies)
	routerGroup.GET("/albums", albumController.GetAlbums)
	routerGroup.POST("/album", albumController.PostAlbum)
}
