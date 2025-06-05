package config

type Config struct {
	ScannerConfig struct {
		MaxFileSize    int64  `json:"maxFileSize"`
		QuarantinePath string `json:"quarantinePath"`
		ThreadCount    int    `json:"threadCount"`
		ScanTimeout    int    `json:"scanTimeout"`
	}

	ServerConfig struct {
		Port    string `json:"port"`
		Host    string `json:"host"`
		LogPath string `json:"logPath"`
	}
}

func NewDefaultConfig() *Config {
	cfg := &Config{}
	cfg.ScannerConfig.MaxFileSize = 100 * 1024 * 1024 // 100MB
	cfg.ScannerConfig.QuarantinePath = "./quarantine"
	cfg.ScannerConfig.ThreadCount = 4
	cfg.ScannerConfig.ScanTimeout = 30

	cfg.ServerConfig.Port = "8080"
	cfg.ServerConfig.Host = "localhost"
	cfg.ServerConfig.LogPath = "./logs"

	return cfg
}
