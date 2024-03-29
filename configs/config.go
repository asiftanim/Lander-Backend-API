package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBSource             string `mapstructure:"DB_SOURCE"`
	SERVER               string `mapstructure:"SERVER"`
	JWT_SECRET           string `mapstructure:"JWT_SECRET"`
	JWT_TOKEN_EXPIRETIME string `mapstructure:"JWT_TOKEN_EXPIRETIME"`
	IMAGE_PATH           string `mapstructure:"IMAGE_PATH"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
