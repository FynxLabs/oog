package oog

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// Load - adapter loader
// func loadAdapter() {
// 	// Load up a adapter from ooglins
// 	pullImage()
// 	runContainer()
// }

// Message - place to send messages
func sendMessage(payload adapterPayload) {
	log := log.With().
		Str("component", "adapter").
		Logger()

	data, err := json.Marshal(payload)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshall json for payload")
		panic(err)
	}

	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "%s/v1/attachment"+payload.URL, body)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create new http request")
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to post message to adapter message endpoint")
		panic(err)
	}
	defer resp.Body.Close()
}

// Channel - Interact with Channels/Rooms
func sendChannel(payload channelPayload) {
	log := log.With().
		Str("component", "adapter").
		Logger()

	data, err := json.Marshal(payload)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshall json for payload")
		panic(err)
	}

	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "%s/v1/channel"+payload.URL, body)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create new http request")
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to post message to adapter channel endpoint")
		panic(err)
	}
	defer resp.Body.Close()
}
