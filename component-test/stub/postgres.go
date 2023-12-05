package stub

import (
	"fmt"
	"github.com/hubzaj/golang-component-test/component-test/config"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"strconv"
	"time"
)

const (
	postgresDefaultPort = "5432/tcp"
)

type StorageStub struct {
	postgres *postgres.PostgresContainer

	cfg *config.TestConfig
}

func InitPostgresStub(cfg *config.TestConfig) *StorageStub {
	postgresContainer, err := postgres.RunContainer(cfg.Ctx,
		testcontainers.WithImage("docker.io/postgres:15.2-alpine"),
		postgres.WithInitScripts("/Users/hubert.zajac/Projects/Private/Golang/golang-component-test/config/resources/init.sql"),
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

	storageStub := &StorageStub{
		postgres: postgresContainer,
		cfg:      cfg,
	}
	storageStub.updateStorageConfig()
	return storageStub
}

func (stub *StorageStub) updateStorageConfig() {
	portMap, _ := stub.postgres.Ports(stub.cfg.Ctx)
	port, _ := strconv.Atoi(portMap[postgresDefaultPort][0].HostPort)
	stub.cfg.ShopConfig.Shop.Storage.Port = port
}

func (stub *StorageStub) CleanupPostgresTestContainer() {
	if err := stub.postgres.Terminate(stub.cfg.Ctx); err != nil {
		log.Fatalf("error terminating postgres stub: %s", err)
	}
}
