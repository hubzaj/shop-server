package album

import (
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
	var newAlbum *model.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	newAlbum.ID = utils.CreateNewUUID()
	ctrl.shopService.AlbumService.RegisterNewAlbum(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
