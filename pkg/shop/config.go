package shop

type Config struct {
	HTTPServer HTTPServer `koanf:"httpServer"`
}

type HTTPServer struct {
	Host            string `koanf:"host"`
	Port            int    `koanf:"port"`
	ShutdownTimeout int    `koanf:"shutdownTimeout"`
}
