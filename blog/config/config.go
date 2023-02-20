// Package config is used to load the configuration file and provide the configuration to the application.
// It is used by the server, database, models, controllers, and main packages.
// It is defined in blog/config/config.go.
// It contains the Config struct and the Load function.
package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Cfg is a pointer to the Config struct.
// It is used to access the configuration values.
var Cfg *Config

// Config is the struct that contains the configuration values.
// It is defined in blog/config/config.go.
// It contains the Server and Database structs.
type Config struct {
	// Server is the struct that contains the server configuration values.
	Server struct {
		// Port is the port that the server will listen on.
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	// Database is the struct that contains the database configuration values.
	Database struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		DBName   string `mapstructure:"dbname"`
	} `mapstructure:"database"`
}

// Load loads the configuration file and returns a pointer to a Config.
// It is defined in blog/config/config.go.
func Load() *Config {
	viper.SetConfigName("config")   // Name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // Path to look for the config file in

	// If a config file is found, read it in.
	// If a config file is not found, log the error and exit the program.
	// For more information on viper, see:
	// https://github.com/spf13/viper
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Unmarshal the config file into a Config struct.
	// If the config file fails to unmarshal, log the error and exit the program.
	var cfg *Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	// Set Cfg to the Config struct.
	Cfg = cfg

	return cfg
}

// LoadDBUrl returns the database URL.
func (c *Config) LoadDBUrl() string {
	// The fmt.Sprintf function is used to format the database URL.
	// For more information on fmt.Sprintf, see:
	// https://golang.org/pkg/fmt/#Sprintf
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName,
	)
}

// ServerPort returns the server port.
func (c *Config) ServerPort() string {
	// The fmt.Sprintf function is used to format the server port.
	return fmt.Sprintf(":%d", c.Server.Port)
}
