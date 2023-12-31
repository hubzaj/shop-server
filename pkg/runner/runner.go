package runner

import (
	"fmt"
	"github.com/hubzaj/golang-component-test/pkg/config"
	"github.com/hubzaj/golang-component-test/pkg/router"
	"github.com/hubzaj/golang-component-test/pkg/shop/controller"
	"github.com/hubzaj/golang-component-test/pkg/shop/services"
	"github.com/hubzaj/golang-component-test/pkg/shop/services/album"
	"github.com/hubzaj/golang-component-test/pkg/storage"
	"github.com/hubzaj/golang-component-test/pkg/utils"
	"log"
	"net"
	"net/http"
	"time"
)

func StartShop(ctx *ShopServerContext) *http.Server {
	return StartShopWithConfig(ctx, nil)
}

func StartShopWithConfig(appContext *ShopServerContext, cfg *config.GeneralConfig) *http.Server {
	initConfig(cfg)
	shopStorage := storage.InitStorage(config.Config.Shop.Storage)

	albumService := &album.Service{
		Storage: shopStorage,
	}

	shopRouter := router.InitRouter(&controller.Dependencies{
		ShopService: &services.ShopService{
			AlbumService: albumService,
		},
	})

	server := &http.Server{
		Addr: fmt.Sprintf(
			"%s:%d",
			config.Config.Shop.HTTPServer.Host,
			config.Config.Shop.HTTPServer.Port,
		),
		ReadHeaderTimeout: 1 * time.Second,
		Handler:           shopRouter,
	}

	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		log.Fatalf("Cannot create http-server listener: %s", err)
	}
	if server.Addr != listener.Addr().String() {
		server.Addr = listener.Addr().String()
	}

	appContext.AddTask()
	go func() {
		defer appContext.TaskDone()

		log.Printf("Shop http-server is listening on port [%s]", server.Addr)
		if err := server.Serve(listener); err != nil {
			log.Fatalf("Shop http-server error: %s", err)
		}
	}()

	return server
}

func initConfig(cfg *config.GeneralConfig) {
	if err := config.InitConfig(cfg); err != nil {
		utils.ExitOnError(err, "error initializing configuration")
	}
}
