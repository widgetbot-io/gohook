package github

import (
	webhook "gopkg.in/go-playground/webhooks.v5/github"
	"github.com/widgetbot-io/gohook/structs"
)

func Handler(c structs.ProviderContext) error {
	hook, _ := webhook.New()
	payload, _ := hook.Parse(c.Context.Request, webhook.PushEvent, webhook.PingEvent, webhook.RepositoryEvent)
	// event will be found because we find it earlier

	return c.Event.Handler(structs.EventContext{
		ID:       c.ID,
		Secret:   c.Secret,
		Event:    c.Event,
		Options:  c.Options,
		Provider: c.Provider,
		Payload:  payload,
	})
}
