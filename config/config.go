package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		DBType string
	}
	Server struct {
		Port string
	}
	ProjectID string
}

func ReadConfig() *Config {
	c := Config{}

	setViper(&c)
	// setDotEnv(&c)

	os.Setenv("PROJECT_ID", c.ProjectID)
	return &c
}

func setDotEnv(c *Config) {
	c.Database.DBType = os.Getenv("DB_TYPE")
	c.Server.Port = os.Getenv("PORT")
	c.ProjectID = os.Getenv("PROJECT_ID")
}

func setViper(c *Config) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join("$GOPATH", "src", "app", "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading envvars: %s", err)
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("Error unmarshalling envvars: %v", err)
	}
}
