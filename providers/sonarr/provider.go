package sonarr

import (
	"encoding/json"
	"git.deploys.io/disweb/gohook/structs"
	"git.deploys.io/disweb/gohook/utils"
	"net/http"
)

func Handler(c structs.ProviderContext) error {
	var TestPayload structs.SonarrTest

	if c.Context.Request.Method != http.MethodPost {
		return utils.ErrInvalidHTTPMethod
	}
	_ = json.Unmarshal([]byte(c.Payload), &TestPayload)

	return c.Event.Handler(structs.EventContext{
		ID:       c.ID,
		Secret:   c.Secret,
		Event:    c.Event,
		Provider: c.Provider,
		Payload:  TestPayload,
	})
}
