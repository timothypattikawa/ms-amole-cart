package config

import (
	"fmt"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func LoadViper() *viper.Viper {

	env := os.Getenv("ENV")

	v := viper.New()
	v.SetConfigName(fmt.Sprintf("application-%s", env))
	v.AddConfigPath(".")
	v.SetConfigType("yml")

	err := v.ReadInConfig()
	if err != nil {
		log.Errorf("fail to open config file err{%e}", err)
	}

	return v
}
