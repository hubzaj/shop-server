package config

import (
	"context"
	"github.com/hubzaj/golang-component-test/pkg/config"
	"github.com/hubzaj/golang-component-test/pkg/shop"
	"github.com/hubzaj/golang-component-test/pkg/storage"
)

type TestConfig struct {
	ShopConfig *config.GeneralConfig

	Ctx context.Context
}

func CreateDefaultConfig(ctx context.Context) *TestConfig {
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
				Storage: &storage.StorageConfig{
					User:     "postgres",
					DBName:   "shop",
					Password: "postgres",
					Host:     "localhost",
					Port:     0,
					SSLMode:  "disable",
				},
			},
		},
		Ctx: ctx,
	}
}
