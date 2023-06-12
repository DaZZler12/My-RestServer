package database

import (
	"context"
	"fmt"

	masteryaml "github.com/Dazzler/My-RestServer/pkg/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

var (
	client *mongo.Client
)

func ConnectToMongoDB() (*mongo.Database, error) {

	// viper.SetConfigFile("../config/master.yaml")
	// err = viper.ReadInConfig()
	// if err != nil {

	// 	return nil, fmt.Errorf("failed to read the file")
	// }

	// config := Config{}
	// fmt.Println("viper.ReadConfig()", config)

	// err = viper.Unmarshal(&config)
	config, err := masteryaml.ExtractYamlForDB()
	if err != nil {

		return nil, fmt.Errorf("failed to unmarshal")
	}

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%v", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port)
	fmt.Println("connectionString", connectionString)
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Print(err)
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	// here also we will try to ping the mongo-db before any operation
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	fmt.Println("Connection established with MongoDB.")
	db := client.Database(config.Database.DbName)
	return db, nil
}

func DisconnectMongoDB(context context.Context) {
	client.Disconnect(context)
	fmt.Println("Connection Closed")
}
