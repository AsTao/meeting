package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	port := viper.GetString("server.port")
	if port == "" {
		fmt.Printf("Server port not specified in the configuration file")
	}

}
