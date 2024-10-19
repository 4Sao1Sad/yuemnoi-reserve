package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName         string   `yaml:"appName"`
	Port            int32    `yaml:"port"`
	ActivityLogPort int32    `yaml:"activityLogPort"`
	PostPort        int32    `yaml:"postPort"`
	Db              DBConfig `yaml:"DB"`
}

type DBConfig struct {
	Host     string `yaml:"host" default:"localhost"`
	Port     string `yaml:"port" default:"5432"`
	Username string `yaml:"username" default:"youruser"`
	Password string `yaml:"password" default:"yourpassword"`
	Database string `yaml:"database" default:"yourdbname"`
}

func Load() *Config {
	config := Config{}
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error occurs while reading the config. ", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("error occurs while unmarshalling the config. ", err)
	}
	return &config
}
