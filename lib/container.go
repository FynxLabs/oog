package botcore

import (
	docker "github.com/fsouza/go-dockerclient"
	"github.com/rs/zerolog/log"
)

func pullImage(endpoint string, image string, tag string) {
	log := log.With().
		Str("component", "container").
		Logger()

	client, err := docker.NewClient(endpoint)
	if err != nil {
		log.Error().Err(err).Msg("Failed to setup docker client")
		panic(err)
	}
	err = client.PullImage(
		docker.PullImageOptions{
			Repository: image,
			Tag:        tag,
		},
		docker.AuthConfiguration{},
	)
	if err != nil {
		log.Error().Err(err).Msg("Failed to pull image")
		panic(err)
	}
	log.Debug().Msg("Image succesfully pulled")
}

func runContainer(image string, cmd string) {
	//
	log := log.With().
		Str("component", "container").
		Logger()

	client, err := docker.NewClient(endpoint)
	if err != nil {
		log.Error().Err(err).Msg("Failed to setup docker client")
		panic(err)
	}

	log.Debug().Msg("Container succesfully stopped")
}

func stopContainer(container string) {
	log := log.With().
		Str("component", "container").
		Logger()

	client, err := docker.NewClient(endpoint)
	if err != nil {
		log.Error().Err(err).Msg("Failed to setup docker client")
		panic(err)
	}

	log.Debug().Msg("Container succesfully stopped")
}
