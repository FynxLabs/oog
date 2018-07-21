package adapter

import "gopkg.in/resty.v1"

// Load - adapter loader
func Load() {
	// Load up a adapter from ooglins
}

// Message - place to send messages
func Message(adapterURL string, msgtype string, text string, attachment map[string]string) {
	payload := make(map[string]interface{})
	payload["type"] = msgtype
	payload["text"] = text
	payload["attachment"] = attachment
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(&AuthSuccess{}).
		Post(adapterURL)
}

// Channel - Interact with Channels/Rooms
func Channel() {

}
