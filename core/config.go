package core

import "github.com/spf13/viper"

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     int    `mapstructure:"DB_PORT"`
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_PORT")
	viper.AutomaticEnv()
	err := viper.Unmarshal(config)
	return config, err
}
