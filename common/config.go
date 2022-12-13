package common

import "github.com/ilyakaznacheev/cleanenv"

// Config for all
var Config config

type config struct {
	Port int    `yaml:"port"`
	DB   string `yaml:"db"`
}

// LoadConfig yaml from file
func LoadConfig(filename string) error {
	return cleanenv.ReadConfig(filename, &Config)
}
