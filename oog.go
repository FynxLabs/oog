package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("bot")
	viper.AddConfigPath("$HOME/.oog/")
	viper.AddConfigPath("/opt/oog/conf/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/v1/stream", stream)
	router.Run() // listen and serve on 0.0.0.0:8080
}

// Assign context to data and forward all messages to listener
func stream(data *gin.Context) {
	switch {
	case strings.Contains(data.Param("text"), "Ping"):
		fmt.Println("one")
	case strings.Contains(data.Param("text"), "Help"):
		fmt.Println("two")
	case strings.Contains(data.Param("text"), "Reload"):
		fmt.Println("three")
	default:
		fmt.Println("four")
	}
}
