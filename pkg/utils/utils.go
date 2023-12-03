package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func ExitOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
		os.Exit(1)
	}
}

func HandlePanic(c *gin.Context, err interface{}) {
	panicErr := errors.New(fmt.Sprint(err))
	fmt.Printf("error during processing request: %s", panicErr)
	c.AbortWithStatus(http.StatusInternalServerError)
}
