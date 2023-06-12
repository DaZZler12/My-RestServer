package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	DbName    string `yaml:"dbname"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Usertable string `yaml:"usertable"`
	Itemtable string `yaml:"itemtable"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

type Server struct {
	Port string `yaml:"port"`
}

type ServerConfig struct {
	Serverconfig Server `yaml:"serverconfig"`
}

var err error

func ExtractYamlForDB() (*Config, error) {
	viper.SetConfigFile("../config/master.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read the file: %w", err)
	}

	config := Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return &config, nil
}

func ExtractYamlForServer() (*ServerConfig, error) {
	viper.SetConfigFile("../config/master.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read the file: %w", err)
	}

	serverconfig := ServerConfig{}
	err = viper.Unmarshal(&serverconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	fmt.Println(serverconfig)
	return &serverconfig, nil
}
