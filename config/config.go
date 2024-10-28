package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName        string   `yaml:"appName"`
	Port           int32    `yaml:"port"`
	ActivityLogUrl string   `yaml:"activityLogUrl"`
	PostUrl        string   `yaml:"postUrl"`
	UserInfoURL    string   `yaml:"userInfoUrl"`
	Db             DBConfig `yaml:"DB"`
	RabbitMQUrl    string   `yaml:"rabbitmqurl"`
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
