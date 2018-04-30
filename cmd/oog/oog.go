package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("bot")
	viper.AddConfigPath("$HOME/.oog/")
	viper.AddConfigPath("/opt/oog/conf/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}
