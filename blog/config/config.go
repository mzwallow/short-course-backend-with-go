package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	// Server
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	// Database
	Database struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
	} `mapstructure:"database"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}

func (c *Config) GetPort() string {
	return fmt.Sprintf(":%d", c.Server.Port)
}

func (c *Config) GetDatabaseURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName,
	)
}
