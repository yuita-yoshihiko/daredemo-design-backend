package config

type Config struct {
	EnvConfig
}

var Conf Config

type EnvConfig struct {
	Environment string `env:"ENVIRONMENT,required"`
}
