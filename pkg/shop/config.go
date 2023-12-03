package shop

type Config struct {
	HTTPServer HTTPServer `koanf:"httpServer"`
	Api        API        `koanf:"api"`
}

type HTTPServer struct {
	Host            string `koanf:"host"`
	Port            int    `koanf:"port"`
	ShutdownTimeout int    `koanf:"shutdownTimeout"`
}

type API struct {
	BaseURL string `koanf:"baseUrl"`
}
