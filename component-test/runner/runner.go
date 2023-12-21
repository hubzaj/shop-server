package runner

import (
	"context"
	"fmt"
	"github.com/hubzaj/golang-component-test/component-test/config"
	shopRunner "github.com/hubzaj/golang-component-test/pkg/runner"
	"net/http"
	"strconv"
	"strings"
)

func StartShop(cfg *config.TestConfig) {
	ctx := shopRunner.NewShopServiceContext(context.Background())
	server := shopRunner.StartShopWithConfig(ctx, cfg.ShopConfig)
	updateConfig(cfg, server)
}

func updateConfig(cfg *config.TestConfig, server *http.Server) {
	updateShopServerPort(cfg, server)
}

func updateShopServerPort(cfg *config.TestConfig, server *http.Server) {
	port, err := strconv.Atoi(strings.Split(server.Addr, ":")[1])
	if err != nil {
		panic(fmt.Errorf("error during parsing server port: %s", err))
	}
	cfg.ShopConfig.Shop.HTTPServer.Port = port
}
