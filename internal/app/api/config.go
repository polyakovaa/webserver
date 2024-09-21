package api

//general instance for API server of REST application

type Config struct {
	//port
	BindAddr string `toml:"bind_addr"`
	//logger
	LoggerLevel string `toml:"logger_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    "8080",
		LoggerLevel: "debug",
	}
}
