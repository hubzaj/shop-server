package runner

import (
	"fmt"
	"github.com/hubzaj/golang-component-test/pkg/config"
	"github.com/hubzaj/golang-component-test/pkg/router"
	"github.com/hubzaj/golang-component-test/pkg/utils"
	"net"
	"net/http"
	"time"
)

func StartShop() *http.Server {
	return StartShopWithConfig(nil)
}

func StartShopWithConfig(cfg *config.GeneralConfig) *http.Server {
	initConfig(cfg)
	//storage.InitStorage()

	server := &http.Server{
		Addr: fmt.Sprintf(
			"%s:%d",
			config.Config.Shop.HTTPServer.Host,
			config.Config.Shop.HTTPServer.Port,
		),
		ReadHeaderTimeout: 1 * time.Second,
		Handler:           router.InitRouter(),
	}

	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		fmt.Println("Cannot create http-server listener")
	}
	if server.Addr != listener.Addr().String() {
		server.Addr = listener.Addr().String()
	}

	go func() {
		fmt.Printf("Shop http-server is listening on port [%s]", server.Addr)
		if err := server.Serve(listener); err != nil {
			fmt.Println("Shop http-server error")
		}
	}()

	return server
}

func initConfig(cfg *config.GeneralConfig) {
	if err := config.InitConfig(cfg); err != nil {
		utils.ExitOnError(err, "error initializing configuration")
	}
}
