package album

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/hubzaj/golang-component-test/pkg/shop/controller"
	"github.com/hubzaj/golang-component-test/pkg/shop/model"
	"github.com/hubzaj/golang-component-test/pkg/shop/services"
	"github.com/hubzaj/golang-component-test/pkg/utils"
	"io"
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
	switch c.Request.Header.Get("Content-Type") {
	case "application/protobuf":
		bindProto(c, newAlbum)
		break
	case "application/json":
		if err := c.BindJSON(&newAlbum); err != nil {
		}
	}
	newAlbum.Id = utils.CreateNewUUID().String()
	ctrl.shopService.AlbumService.RegisterNewAlbum(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func bindProto(c *gin.Context, album *model.Album) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "Error reading request body", err)
	}
	if err = proto.Unmarshal(body, album); err != nil {
		c.String(http.StatusBadRequest, "Error unmarshalling Protobuf request", err)
	}
}
