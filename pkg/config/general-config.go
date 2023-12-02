package config

import (
	"github.com/hubzaj/golang-component-test/pkg/shop"
)

type GeneralConfig struct {
	Shop shop.Config `koanf:"shop"`
}
