package github

import (
	"git.deploys.io/disweb/gohook/structs"
	webhook "gopkg.in/go-playground/webhooks.v5/github"
)

func Handler(c structs.ProviderContext) error {
	hook, _ := webhook.New()
	payload, _ := hook.Parse(c.Context.Request, webhook.PushEvent)
	// event will be found because we find it earlier

	return c.Event.Handler(structs.EventContext{
		ID:       c.ID,
		Secret:   c.Secret,
		Event:    c.Event,
		Provider: c.Provider,
		Payload:  payload,
	})
}