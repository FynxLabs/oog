package router

import (
	"fmt"
	"internal/botcore"
	"strings"

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
	adapter := botcore.adapter()
	brain := botcore.brain()
	plugin := botcore.plugin()

	// Bot specific routes
	botRG := router.Group("/v1/bot")
	{
		botRG.POST("/stream", stream)
		botRG.POST("/reload", plugin.load())
		botRG.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	// Adapter specific routes
	adapterRG := router.Group("/v1/adapter")
	{
		adapterRG.POST("/channel", adapter.channel())
		adapterRG.POST("/message", adapter.message())
	}

	// Brain specific routes
	brainRG := router.Group("/v1/brain")
	{
		brainRG.POST("/save", brain.save())
		brainRG.POST("/delete", brain.delete())
		brainRG.POST("/query", brain.query())
	}
	router.Run() // listen and serve on 0.0.0.0:8080
}

// Assign context to data and forward all messages to listener
func stream(data *gin.Context) {
	switch {
	case strings.Contains(data.Param("text"), "Ping"):
		adapter.message("plain", "ping")
	case strings.Contains(data.Param("text"), "Help"):
		plugin.help_list()
	case strings.Contains(data.Param("text"), "Reload"):
		plugin.load()
	default:
		plugin.exec()
	}
}
