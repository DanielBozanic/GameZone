package config

import "github.com/spf13/viper"

type Config struct {
	DBURL     string `mapstructure:"DB_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile("./config.env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}