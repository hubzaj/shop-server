package stub

import "github.com/hubzaj/golang-component-test/component-test/config"

type Stub struct {
}

func InitStubs(cfg *config.TestConfig) *Stub {
	InitPostgresStub(cfg)
	return &Stub{}
}
