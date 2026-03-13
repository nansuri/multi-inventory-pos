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

	// Set Defaults
	viper.SetDefault("DB_HOST", "172.18.0.2")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "Sv2K8OD313HnLEIb")
	viper.SetDefault("DB_NAME", "palugada_dashboard")
	viper.SetDefault("DB_SSLMODE", "disable")
	viper.SetDefault("REDIS_HOST", "localhost")
	viper.SetDefault("REDIS_PORT", "6379")
	viper.SetDefault("APP_PORT", "85")
	viper.SetDefault("GIN_MODE", "debug")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Note: .env file not found, using environment variables or defaults")
	}

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
