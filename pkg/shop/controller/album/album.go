package album

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/hubzaj/golang-component-test/pkg/shop/controller"
	"github.com/hubzaj/golang-component-test/pkg/shop/model"
	"github.com/hubzaj/golang-component-test/pkg/shop/services"
	"github.com/hubzaj/golang-component-test/pkg/utils"

	"net/http"
)

type Controller struct {
	shopService *services.ShopService
}

func NewController(dependencies *controller.Dependencies) *Controller {
	return &Controller{
		shopService: dependencies.ShopService,
	}
}

func (ctrl *Controller) GetAlbums(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			utils.HandlePanic(c, err)
		}
	}()
	ctrl.processGetAlbumRequest(c)
}

func (ctrl *Controller) processGetAlbumRequest(c *gin.Context) {
	albums := ctrl.shopService.AlbumService.GetAvailableAlbums()
	c.IndentedJSON(http.StatusOK, albums)
}

func (ctrl *Controller) PostAlbum(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			utils.HandlePanic(c, err)
		}
	}()
	ctrl.processPostAlbum(c)
}

func (ctrl *Controller) processPostAlbum(c *gin.Context) {
	var newAlbum model.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	newAlbum.ID = generateRandomString(10)
	// TODO: Replace for database - postgres
	albums := []model.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// TODO: Remove this method and replace with UUID generator
func generateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
