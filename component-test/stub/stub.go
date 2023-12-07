package stub

import (
	"github.com/hubzaj/golang-component-test/component-test/client"
	"github.com/hubzaj/golang-component-test/component-test/config"
)

type Stub struct {
	ShopClient *ShopClient
	Storage    *StorageStub
}

func InitStubs(cfg *config.TestConfig) *Stub {
	httpClient := client.NewHTTPClient(cfg)

	return &Stub{
		ShopClient: NewShopClient(httpClient),
		Storage:    InitPostgresStub(cfg),
	}
}

func (stub *Stub) Cleanup() {
	stub.Storage.CleanupPostgresTestContainer()
}
