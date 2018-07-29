package main

import (
	"strings"

	"github.com/FynxLabs/oog/lib/botcore"
	"github.com/gin-gonic/gin"
)

func run() {
	router := gin.Default()
	client := botcore.chatClient() // Create a new client by connecting via client
	brain := botcore.dataClient()  // Load up a brain

	// Bot specific routes
	botRG := router.Group("/v1/bot")
	{
		botRG.POST("/stream", stream)               // Accept all incoming messages and send to stream function to sort
		botRG.POST("/reload", botcore.pluginLoad()) // Reload all plugins
		botRG.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	// Adapter specific routes
	adapterRG := router.Group("/v1/adapter")
	{
		adapterRG.POST("/channel", botcore.Channel(client)) // Endpoint to interact with Channels/Rooms
		adapterRG.POST("/message", botcore.Message(client)) // Endpoint to send message via client
	}

	// Brain specific routes
	brainRG := router.Group("/v1/brain")
	{
		brainRG.POST("/save", botcore.Save(brain))     // Endpoint to save data to selected datastore
		brainRG.POST("/delete", botcore.Delete(brain)) // Endpoint to delete data from selected datastore
		brainRG.POST("/query", botcore.Query(brain))   // Endpoint to query selected datastore
	}
	router.Run() // listen and serve on 0.0.0.0:8080
}

// Assign context to data and forward all messages to listener
func stream(data *gin.Context) {
	switch {
	case strings.Contains(data.Param("text"), "Ping"):
		botcore.Message(client, "plain", "ping")
	case strings.Contains(data.Param("text"), "Help"):
		botcore.HelpList()
	case strings.Contains(data.Param("text"), "Reload"):
		botcore.pluginLoad()
	default:
		botcore.pluginExec()
	}
}
