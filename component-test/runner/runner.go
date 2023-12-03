package runner

import (
	"fmt"
	"github.com/hubzaj/golang-component-test/pkg/config"
	shopRunner "github.com/hubzaj/golang-component-test/pkg/runner"
	"strconv"
	"strings"
)

func StartShop(cfg *config.GeneralConfig) {
	server := shopRunner.StartShopWithConfig(cfg)
	port, err := strconv.Atoi(strings.Split(server.Addr, ":")[1])
	if err != nil {
		panic(fmt.Errorf("error during parsing server port: %s", err))
	}
	cfg.Shop.HTTPServer.Port = port
}
