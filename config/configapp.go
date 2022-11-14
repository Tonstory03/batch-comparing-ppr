package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var AppConfig Config

const CURRENT_FILE = "configcmd"

func LoadingConfig() {

	viper.SetConfigName("/configmap/config")

	viper.AddConfigPath(".")

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		getDefaultConfigYML()
	} else {
		err := viper.Unmarshal(&AppConfig)

		if err != nil {
			getDefaultConfigYML()
		}
	}

	fmt.Println("Load config success profile:", GetApplication().Profile)

}

func getDefaultConfigYML() {

	env := os.Getenv("ENV")

	configName := getConfigName(env)

	viper.SetConfigName(configName)

	viper.AddConfigPath(".")

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(&AppConfig)

	if err != nil {
		panic(err)
	}
}

func GetConfig() Config {
	return AppConfig
}

func GetApplication() Application {
	return GetConfig().Application
}

func GetServer() Server {
	return GetConfig().Server
}

func GetCronJobs() []Cronjob {
	return GetConfig().Cronjob
}

func GetElasticConfig() Elastic {
	return AppConfig.Elastic
}

func getConfigName(env string) string {

	configName := "./config"

	switch env {
	case "local":
		configName += ".local"
	}
	return configName
}
