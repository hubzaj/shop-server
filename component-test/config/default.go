package config

import (
	"github.com/hubzaj/golang-component-test/pkg/config"
	"github.com/hubzaj/golang-component-test/pkg/shop"
)

func CreateDefaultConfig() *config.GeneralConfig {
	return &config.GeneralConfig{
		Shop: &shop.Config{
			HTTPServer: &shop.HTTPServer{
				Host:            "localhost",
				Port:            0,
				ShutdownTimeout: 5,
			},
			Api: &shop.API{
				BaseURL: "/api/v1/shop",
			},
		},
	}
}
