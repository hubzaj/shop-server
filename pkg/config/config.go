package config

import (
	"fmt"
	"github.com/hubzaj/golang-component-test/pkg/utils"
	"os"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var k = koanf.New(".")

var Config *GeneralConfig

func InitConfig() error {
	return initConfig(getConfigName())
}

func initConfig(configName string) error {
	if err := loadConfig("default"); err != nil {
		return err
	}
	if err := loadConfig(configName); err != nil {
		return err
	}
	SetGeneralConfig()
	return nil
}

func getConfigName() string {
	if configName := strings.ToLower(os.Getenv("SHOP_CONFIG")); configName != "" {
		return configName
	}
	return "local"
}
func getConfigPath(configName string) string {
	projectAbsolutePath, _ := os.Getwd()
	return fmt.Sprintf("%s/config/%s.yaml", projectAbsolutePath, configName)
}

func loadConfig(configName string) error {
	if err := k.Load(file.Provider(getConfigPath(configName)), yaml.Parser()); err != nil {
		return fmt.Errorf("fatal error loading default config file: %s", err)
	}
	return nil
}

func SetGeneralConfig() {
	generalConfig := &GeneralConfig{}
	if err := k.Unmarshal("", generalConfig); err != nil {
		utils.ExitOnError(err, "error unmarshalling shop k")
	}
	Config = generalConfig
}
