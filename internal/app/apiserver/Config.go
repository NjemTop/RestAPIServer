package apiserver

type Config struct {
	BindAddr string `toml:"build_addr`
}

func NewConfig() *Config {
	return &Config{
		build_addr ":8070"
	}
}