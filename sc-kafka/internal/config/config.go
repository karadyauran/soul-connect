package config

import (
	"github.com/spf13/viper"
	"log"
)

type KafkaConfig struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	Brokers    string `mapstructure:"BROKERS"`
	Topic      string `mapstructure:"TOPIC"`
}

func LoadKafkaConfig(path string) (config KafkaConfig, err error) {
	viper.SetConfigFile(path + ".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("could not loadconfig: %v", err)
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("could not loadconfig: %v", err)
	}

	return
}
