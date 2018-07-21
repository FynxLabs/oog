package main

import (
	"strings"

	"internal/adapter"
	"internal/datastore"
	"internal/plugin"

	"github.com/gin-gonic/gin"
)

func run() {
	router := gin.Default()
	client := adapter.Load()  // Create a new client by connecting via client
	brain := datastore.Load() // Load up a brain

	// Bot specific routes
	botRG := router.Group("/v1/bot")
	{
		botRG.POST("/stream", stream)      // Accept all incoming messages and send to stream function to sort
		botRG.POST("/reload", plugin.Load) // Reload all plugins
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
