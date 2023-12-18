package config

import (
	"fmt"
	"github.com/hubzaj/golang-component-test/pkg/utils"
	"os"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
)

var k = koanf.New(".")

var Config *GeneralConfig

func InitConfig(cfg *GeneralConfig) error {
	return initConfigWithConfig(cfg, getConfigName())
}

func initConfigWithConfig(cfg *GeneralConfig, configName string) error {
	if cfg != nil {
		if err := loadConfigFromStruct(cfg); err != nil {
			return err
		}
	} else {
		if err := loadConfig("default"); err != nil {
			return err
		}
		if err := loadConfig(configName); err != nil {
			return err
		}
	}
	if err := loadConfigFromEnvVariables(); err != nil {
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
	return fmt.Sprintf("%s/config/%s.yaml", utils.GetProjectRootAbsolutePath(), configName)
}

func loadConfig(configName string) error {
	if err := k.Load(file.Provider(getConfigPath(configName)), yaml.Parser()); err != nil {
		return fmt.Errorf("fatal error loading default config file: %s", err)
	}
	return nil
}

func loadConfigFromStruct(cfg *GeneralConfig) error {
	if err := k.Load(structs.Provider(cfg, "koanf"), nil); err != nil {
		return fmt.Errorf("fatal error loading config from struct: %s", err)
	}
	return nil
}

func loadConfigFromEnvVariables() error {
	if err := k.Load(env.Provider("SHOP_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(s), "_", ".", -1)
	}), nil); err != nil {
		return fmt.Errorf("fatal error loading config from env variables: %s", err)
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
