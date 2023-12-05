package shop

import "github.com/hubzaj/golang-component-test/pkg/storage"

type Config struct {
	HTTPServer *HTTPServer            `koanf:"httpServer"`
	Api        *API                   `koanf:"api"`
	Storage    *storage.StorageConfig `koanf:"storage"`
}

type HTTPServer struct {
	Host            string `koanf:"host"`
	Port            int    `koanf:"port"`
	ShutdownTimeout int    `koanf:"shutdownTimeout"`
}

type API struct {
	BaseURL string `koanf:"baseUrl"`
}
