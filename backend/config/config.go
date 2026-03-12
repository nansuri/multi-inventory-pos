package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSslMode  string `mapstructure:"DB_SSLMODE"`

	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`

	AppPort   string `mapstructure:"APP_PORT"`
	JWTSecret string `mapstructure:"JWT_SECRET"`
	GinMode   string `mapstructure:"GIN_MODE"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// If .env exists, read it. If not, it's okay because we use env vars.
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Note: .env file not found, using environment variables")
	}

	// This is key: tell viper how to unmarshal environment variables into the struct
	// by explicitly binding each field to its corresponding env var.
	// Alternatively, just use AutomaticEnv and Unmarshal with a decoder hook,
	// but simpler is to just bind or use defaults.
	
	// Bind all keys to environment variables
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_SSLMODE")
	viper.BindEnv("REDIS_HOST")
	viper.BindEnv("REDIS_PORT")
	viper.BindEnv("REDIS_PASSWORD")
	viper.BindEnv("APP_PORT")
	viper.BindEnv("JWT_SECRET")
	viper.BindEnv("GIN_MODE")

	err = viper.Unmarshal(&config)
	return
}
