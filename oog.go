package oog

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var ()

Viper.SetConfigName("oog.config")
viper.AddConfigPath("$HOME/.oog")
viper.AddConfigPath("/opt/oog/conf/")
viper.AddConfigPath(".")
err := viper.ReadInConfig() // Find and read the config file
if err != nil { // Handle errors reading the config file
	panic(fmt.Errorf("Fatal error config file: %s \n", err))
}

// fun oog()
func oog() {

}
