package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Database DBConfig
}

type AppConfig struct {
	Port        int    `mapstructure:"PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`
}

type DBConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

type JWTConfig struct {
	Secret   string `mapstructure:"JWT_SECRET"`
	ExpireAt string `mapstructure:"JWT_EXPIRE"`
}

func LoadConfig() (cfg *Config, err error) {
	// Setting the default name and path of the config file
	viper.AddConfigPath(".env")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Unmarshal configuration into the Config struct
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil

}
