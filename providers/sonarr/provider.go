package sonarr

import (
	"git.deploys.io/disweb/gohook/structs"
	"git.deploys.io/disweb/gohook/utils"
	"net/http"
)

func Handler(c structs.ProviderContext) error {
	if c.Context.Request.Method != http.MethodPost {
		return utils.ErrInvalidHTTPMethod
	}
	payload := ""

	return c.Event.Handler(structs.EventContext{
		ID:       c.ID,
		Secret:   c.Secret,
		Event:    c.Event,
		Provider: c.Provider,
		Payload:  payload,
	})
}
