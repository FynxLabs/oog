package oog

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func newConfig() {
	log := log.With().
		Str("component", "container").
		Logger()
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Error().Err(err).Msg("Failed to find home directory")
		panic(err)
	}

	// Search config in home and other directories with name ".oog" (without extension).
	viper.SetConfigName(".oog")
	viper.AddConfigPath(home)
	viper.AddConfigPath("/opt/oog/conf/")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Err(err).Msg("Failed to read config file")
		panic(err)
	}
}
