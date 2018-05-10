package adapter

import "gopkg.in/resty.v1"

// Load - adapter loader
func Load() {
	// Load up a adapter from ooglins
}

// Message - place to send messages
func Message(msgtype string, text string, attachment map[string]string) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(&AuthSuccess{}).
		Post(adapter_url)
}

// Channel - Interact with Channels/Rooms
func Channel() {

}
