package botcore

// Load plugins
func pluginLoad(images [:]) {
	for _, image := range images {
		pullImage(endpoint, image.name, image.tag)
	}

	if err != nil {
		log.Fatal(err)
	}
}

// Exec against plugins
func pluginExec() {

}

// HelpList Build and return help list
func HelpList() {

}
