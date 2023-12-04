package config

import (
	"github.com/hubzaj/golang-component-test/pkg/config"
	"github.com/hubzaj/golang-component-test/pkg/shop"
)

type TestConfig struct {
	ShopConfig *config.GeneralConfig
}

func CreateDefaultConfig() *TestConfig {
	return &TestConfig{
		ShopConfig: &config.GeneralConfig{
			Shop: &shop.Config{
				HTTPServer: &shop.HTTPServer{
					Host:            "localhost",
					Port:            0,
					ShutdownTimeout: 5,
				},
				Api: &shop.API{
					BaseURL: "api/v1/shop",
				},
			},
		},
	}
}
