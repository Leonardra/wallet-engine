package configs

import (
	"github.com/spf13/viper"
	"walletEngine/util"
)


type Config struct {
	DBURL      string `mapstructure:"DB_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")


	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		util.ApplicationLog.Fatal("Error loading config file %v\n", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		util.ApplicationLog.Fatal("Error unmarshalling %v\n", err)
	}
	return
}
