package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	var err error
	config = viper.New()
	config.SetDefault("http.port", "8080")
	config.SetConfigType("yaml")
	config.Set("environment", env)
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")

	config.AutomaticEnv()
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = config.ReadInConfig()
	if err != nil {
		log.Warning("Could not parse configuration file, using environment variables")
	}
}

func GetConfig() *viper.Viper {
	return config
}
