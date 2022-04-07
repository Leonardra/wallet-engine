package configs

import (
	"github.com/spf13/viper"
	"os"
	"walletEngine/util"
)


type Config struct {
	DB_URL      string `mapstructure:"DB_URL"`
	SERVER_PORT  string `mapstructure:"PORT"`
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

func EnvHTTPPort() string {
	config, _ := LoadConfig("C:\\Users\\ADMIN\\GolandProjects\\walletEngine")
	return os.Getenv(config.SERVER_PORT)
}
