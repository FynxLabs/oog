package oog

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type adapterContext struct {
	Type    string `json:"type"`
	User    string `json:"user,omitempty"`
	Text    string `json:"text,omitempty"`
	Channel string `json:"channel,omitempty"`
	URL     string `json:"url,omitempty"`
}

type adapterAttachment struct {
	attachment map[string]string `json:"attachment,omitempty"`
}

type adapterPayload struct {
	adapterContext
	adapterAttachment
}

type channelPayload struct {
	Type    string `json:"type"`
	User    string `json:"user,omitempty"`
	Text    string `json:"text,omitempty"`
	Channel string `json:"channel,omitempty"`
	URL     string `json:"url,omitempty"`
}

func run() {

	newConfig() // Loads new viper config found via config.go
	log.Debug().Msg("Loaded new config")

	router := gin.Default()

	log := log.With().
		Str("component", "container").
		Logger()

	// client := chatClient()
	// Open the bolt.db data file in your brain directory.
	// It will be created if it doesn't exist.
	// brain, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// Bot specific routes
	botRG := router.Group("/v1/bot")
	{
		botRG.POST("/stream", func(data *gin.Context) {
			stream(data)
		}) // Accept all incoming messages and send to stream function to sort
		botRG.POST("/reload", func(data *gin.Context) {
			pluginLoad()
		}) // Reload all plugins
		botRG.GET("/ping", func(data *gin.Context) {
			data.JSON(200, gin.H{
				"message": "pong",
			})
		}) // Return ping check
	}
	log.Debug().Msg("Loaded /v1/bot routes")

	// Adapter specific routes
	adapterRG := router.Group("/v1/adapter")
	{
		adapterRG.POST("/channel", func(data *gin.Context) {
			sendChannel(channelPayload{
				data.Param("type"),
				data.Param("user"),
				data.Param("text"),
				data.Param("channel"),
				"",
			}) // Endpoint to interact with Channels/Rooms
		})
		adapterRG.POST("/message", func(data *gin.Context) {
			sendMessage(adapterPayload{
				adapterContext{
					data.Param("type"),
					data.Param("user"),
					data.Param("text"),
					data.Param("channel"),
					"",
				},
				adapterAttachment{},
			}) // Endpoint to send message via client
		})
	}
	log.Debug().Msg("Loaded /v1/adapter routes")

	// Brain specific routes
	// brainRG := router.Group("/v1/brain")
	// {
	// 	brainRG.POST("/save", Save(brain))     // Endpoint to save data to selected datastore
	// 	brainRG.POST("/delete", Delete(brain)) // Endpoint to delete data from selected datastore
	// 	brainRG.POST("/query", Query(brain))   // Endpoint to query selected datastore
	// }

	router.Run() // listen and serve on 0.0.0.0:8080
}

// Assign context to data and forward all messages to listener
func stream(data *gin.Context) {

	log := log.With().
		Str("component", "container").
		Logger()
	switch {
	case strings.Contains(data.Param("text"), "Ping"):
		log.Debug().Msg("Calling pong check")
		sendMessage(adapterPayload{
			adapterContext{
				data.Param("type"),
				data.Param("user"),
				data.Param("text"),
				data.Param("channel"),
				"",
			},
			adapterAttachment{},
		})
	case strings.Contains(data.Param("text"), "Help"):
		log.Debug().Msg("Calling help list")
		HelpList()
	case strings.Contains(data.Param("text"), "Reload"):
		log.Debug().Msg("Loading plugins")
		pluginLoad()
	default:
		log.Debug().Msg("Attempting to exec against plugin")
		pluginExec()
	}
}
