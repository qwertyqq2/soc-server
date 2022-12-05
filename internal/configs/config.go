package configs

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ParamsProvider struct {
		ProviderURL          string `yaml:"provider_url"`
		GroupContractAddress string `yaml:"group_address"`
	} `yaml:"params_listner"`
	DB struct {
		DbUrl string `yaml:"db_url"`
	} `yaml:"db"`
	Server struct {
		Type   string `yaml:"type"`
		BindIp string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"server"`
}

func GetConfig() *Config {
	conf := &Config{}
	if err := cleanenv.ReadConfig("configs.yaml", conf); err != nil {
		log.Fatal(err)
	}
	return conf
}
