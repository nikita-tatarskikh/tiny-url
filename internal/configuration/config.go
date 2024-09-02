package configuration

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerConfig   ServerConfig   `mapstructure:"server"`
	RedisConfig    RedisConfig    `mapstructure:"redis"`
	PostgresConfig PostgresConfig `mapstructure:"postgres"`
	KafkaConfig    KafkaConfig    `mapstructure:"kafka"`
}

type ServerConfig struct {
	ServerAddress string `mapstructure:"address"`
	ServerPort    string `mapstructure:"port"`
}

type RedisConfig struct {
	RedisAddress  string `mapstructure:"address"`
	RedisPort     string `mapstructure:"port"`
	RedisPassword string `mapstructure:"password"`
	RedisDB       int    `mapstructure:"db"`
}

type PostgresConfig struct {
	PostgresAddress  string `mapstructure:"address"`
	PostgresPort     string `mapstructure:"port"`
	PostgresUser     string `mapstructure:"user"`
	PostgresPassword string `mapstructure:"password"`
	PostgresDatabase string `mapstructure:"database"`
}

type KafkaConfig struct {
	Producers string `mapstructure:"producers"`
}

func NewConfig() (Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("conf")

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("error reading config file %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return config, nil
}
