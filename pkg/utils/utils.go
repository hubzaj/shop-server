package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
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

func GetProjectRootAbsolutePath() string {
	initialAbsolutePath, _ := os.Getwd()
	currentAbsolutePath := initialAbsolutePath
	fileInProjectRootPath := "go.mod"
	basePath := "/"
	for currentAbsolutePath != basePath {
		if _, err := os.Stat(fmt.Sprintf("%s/%s", currentAbsolutePath, fileInProjectRootPath)); !os.IsNotExist(err) {
			return currentAbsolutePath
		}
		currentAbsolutePath = filepath.Dir(currentAbsolutePath)
	}
	return initialAbsolutePath
}
