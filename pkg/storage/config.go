package storage

type StorageConfig struct {
	User     string `koanf:"user"`
	DBName   string `koanf:"dbName"`
	Password string `koanf:"password"`
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	SSLMode  string `koanf:"sslMode"`
}
