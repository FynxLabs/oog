package oog

import log "github.com/rs/zerolog/log"

// Load plugins
func pluginLoad() {
	log := log.With().
		Str("component", "plugin").
		Logger()

	log.Debug().Msg("Pulling plugin images")

	// for _, image := range images {
	// 	log.Debug().Msg("Pulling: %s", image)
	// 	pullImage(endpoint, image.name, image.tag)
	// }

	// if err != nil {
	// 	log.Fatal(err)
	// }
}

// Exec against plugins
func pluginExec() {
	log := log.With().
		Str("component", "plugin").
		Logger()

	log.Debug().Msg("Getting Help List")
}

// HelpList Build and return help list
func HelpList() {
	log := log.With().
		Str("component", "plugin").
		Logger()

	log.Debug().Msg("Getting Help List")
}
