package sonarr

import (
	"encoding/json"
	"git.deploys.io/disweb/gohook/structs"
)

func Handler(c structs.ProviderContext) error {
	var TestPayload structs.SonarrTest

	_ = json.Unmarshal([]byte(c.Payload), &TestPayload)

	return c.Event.Handler(structs.EventContext{
		ID:       c.ID,
		Secret:   c.Secret,
		Event:    c.Event,
		Provider: c.Provider,
		Payload:  TestPayload,
	})
}
