package app

type Config struct {
}

func InitConfig() (*Config, error) {
	return &Config{}, nil
}
