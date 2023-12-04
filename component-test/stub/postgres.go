package stub

import (
	"fmt"
	"github.com/hubzaj/golang-component-test/component-test/config"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"time"
)

type StorageStub struct {
	postgres *postgres.PostgresContainer

	cfg *config.TestConfig
}

func InitPostgresStub(cfg *config.TestConfig) *StorageStub {
	postgresContainer, err := postgres.RunContainer(cfg.Ctx,
		testcontainers.WithImage("docker.io/postgres:15.2-alpine"),
		//postgres.WithInitScripts(filepath.Join("testdata", "init-user-db.sh")),
		//postgres.WithConfigFile(filepath.Join("testdata", "my-postgres.conf")),
		postgres.WithDatabase(cfg.ShopConfig.Shop.Storage.DBName),
		postgres.WithUsername(cfg.ShopConfig.Shop.Storage.User),
		postgres.WithPassword(cfg.ShopConfig.Shop.Storage.Password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		panic(fmt.Errorf("error setting up postgres test container: %s", err))
	}

	return &StorageStub{
		postgres: postgresContainer,
		cfg:      cfg,
	}
}

func (stub *StorageStub) CleanupPostgresTestContainer() {
	if err := stub.postgres.Terminate(stub.cfg.Ctx); err != nil {
		log.Fatalf("error terminating postgres stub: %s", err)
	}
}
