package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Cfg *Config

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Database struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
	} `mapstructure:"database"`
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	var cfg *Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	Cfg = cfg

	return cfg
}

func (c *Config) LoadDBUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
	)
}

func (c *Config) ServerPort() string {
	return fmt.Sprintf(":%d", c.Server.Port)
}
