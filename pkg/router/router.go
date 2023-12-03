package router

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/hubzaj/golang-component-test/pkg/controllers/shop"
	"net/http"
)

type Config struct {
	router *gin.Engine
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	NewHandler(&Config{router: router})
	return router
}

func NewHandler(config *Config) {
	shopRouterGroup := config.router.Group("/api/v1/shop")

	addAlbumRoutes(shopRouterGroup)

	config.router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
}

func addAlbumRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/albums", controllers.GetAlbums)
}
