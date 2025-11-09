package config

import (
	"log"

	"gopkg.in/yaml.v2"
)

func GenerateConfig(conf []byte) Config {
	confData := Config{}

	err := yaml.Unmarshal(conf, &confData)

	if err != nil {
		log.Println(err)
	}

	return confData

}
