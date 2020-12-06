package sonarr

import (
	"encoding/json"
	"fmt"
	"github.com/widgetbot-io/gohook/structs"
)

func Handler(c structs.ProviderContext) error {
	payload, _ := eventParsing(c.EventName, c.Payload)

	return c.Event.Handler(structs.EventContext{
		ID:       c.ID,
		Secret:   c.Secret,
		Options:  c.Options,
		Event:    c.Event,
		Provider: c.Provider,
		Payload:  payload,
	})
}

func eventParsing(event string, payload []byte) (interface{}, error) {
	switch event {
	case "Test":
		var pl structs.SonarrTest
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case "Grab":
		var pl structs.SonarrGrab
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case "Download":
		var pl structs.SonarrDownload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err
	default:
		return nil, fmt.Errorf("unknown event %s", event)
	}
}
