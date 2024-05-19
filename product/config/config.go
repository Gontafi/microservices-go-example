package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

var AppConfig *Config

type Config struct {
	ServerPort       int           `mapstructure:"SERVER_PORT"`
	DBHost           string        `mapstructure:"DB_HOST"`
	DBPort           string        `mapstructure:"DB_PORT"`
	DBUser           string        `mapstructure:"DB_USER"`
	DBDriver         string        `mapstructure:"DB_DRIVER"`
	DBPassword       string        `mapstructure:"DB_PASSWORD"`
	DBName           string        `mapstructure:"DB_NAME"`
	DB_MAX_CONN      int           `mapstructure:"DB_MAX_CONN"`
	DB_MAX_IDLE_CONN int           `mapstructure:"DB_MAX_IDLE_CONN"`
	DB_MAX_IDLE_TIME time.Duration `mapstructure:"DB_MAX_IDLE_TIME"`
	GRPCPort         int           `mapstructure:"GRPC_PORT"`
	AuthGRPCPort     int           `mapstructure:"AUTH_GRPC_PORT"`
	AuthAppHost      string        `mapstructure:"AUTH_APP_HOST"`
}

func LoadConfig() error {
	err := os.Setenv("TZ", "Asia/Almaty")
	if err != nil {
		return err
	}

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("/app/docker")
	viper.AddConfigPath("/docker/")

	cfg := Config{}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper could not read config file:%v", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	AppConfig = &cfg

	return nil
}
