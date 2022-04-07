package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	var err error
	config = viper.New()
	config.SetDefault("http.port", "8080")
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")

	config.AutomaticEnv()
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = config.ReadInConfig()
	if err != nil {
		fmt.Println("Could not parse configuration file, using environment variables")
	}
}

func GetConfig() *viper.Viper {
	return config
}
