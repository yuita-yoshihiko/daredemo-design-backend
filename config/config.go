package config

type Config struct {
	EnvConfig
	DBConfig
}

var Conf Config

type EnvConfig struct {
	Environment string `env:"ENVIRONMENT,required"`
}

type DBConfig struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
}
