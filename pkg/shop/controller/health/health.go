package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl *Controller) GetHealthStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hi! I am alive!")
}
