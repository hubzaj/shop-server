package stub

import (
	"github.com/hubzaj/golang-component-test/component-test/config"
)

type Stub struct {
	Storage *StorageStub
}

func InitStubs(cfg *config.TestConfig) *Stub {
	return &Stub{
		Storage: InitPostgresStub(cfg),
	}
}

func (stub *Stub) Cleanup() {
	stub.Storage.CleanupPostgresTestContainer()
}
