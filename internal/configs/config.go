package configs

import (
	"log"
	"sync"

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
	}
}

var once sync.Once

func GetConfig() *Config {
	conf := &Config{}
	once.Do(func() {
		if err := cleanenv.ReadConfig("config.yaml", conf); err != nil {
			log.Fatal(err)
		}
	})
	log.Println(conf.ParamsProvider)
	return conf
}
