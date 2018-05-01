package main

import (
	"fmt"
	"strings"

	"github.com/fynxlabs/oog/internal/adapter"
	"github.com/fynxlabs/oog/internal/datastore"
	"github.com/fynxlabs/oog/internal/plugin"
	"github.com/gin-gonic/gin"
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

func main() {
	router := gin.Default()
	client := adapter.Load()  // Create a new client by connecting via client
	brain := datastore.Load() // Load up a brain

	// Bot specific routes
	botRG := router.Group("/v1/bot")
	{
		botRG.POST("/stream", stream)        // Accept all incoming messages and send to stream function to sort
		botRG.POST("/reload", plugin.Load()) // Reload all plugins
		botRG.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	// Adapter specific routes
	adapterRG := router.Group("/v1/adapter")
	{
		adapterRG.POST("/channel", client.Channel()) // Endpoint to interact with Channels/Rooms
		adapterRG.POST("/message", client.Message()) // Endpoint to send message via client
	}

	// Brain specific routes
	brainRG := router.Group("/v1/brain")
	{
		brainRG.POST("/save", brain.Save())     // Endpoint to save data to selected datastore
		brainRG.POST("/delete", brain.Delete()) // Endpoint to delete data from selected datastore
		brainRG.POST("/query", brain.Query())   // Endpoint to query selected datastore
	}
	router.Run() // listen and serve on 0.0.0.0:8080
}

// Assign context to data and forward all messages to listener
func stream(data *gin.Context) {
	switch {
	case strings.Contains(data.Param("text"), "Ping"):
		adapter.Message("plain", "ping")
	case strings.Contains(data.Param("text"), "Help"):
		plugin.HelpList()
	case strings.Contains(data.Param("text"), "Reload"):
		plugin.Load()
	default:
		plugin.Exec()
	}
}
