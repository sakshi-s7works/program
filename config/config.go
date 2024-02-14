// config/config.go

package config

import (
	"github.com/spf13/viper"
	"log"
)

// InitializeConfig initializes the configuration using Viper
func InitializeConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading configuration file:", err)
	}
}

// GetJWTSecretKey returns the JWT secret key from the configuration
func GetJWTSecretKey() string {
	secretKey := viper.GetString("jwt.secretKey")
	if secretKey == "" {
		log.Fatal("JWT secret key not found in configuration")
	}
	return secretKey
}
