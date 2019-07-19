package gitlab

import (
	"git.deploys.io/disweb/gohook/structs"
	webhook "gopkg.in/go-playground/webhooks.v5/gitlab"
)

func Handler(c structs.ProviderContext) error {
	hook, _ := webhook.New()
	payload, _ := hook.Parse(c.Context.Request, webhook.PushEvents)
	// event will be found because we find it earlier

	return c.Event.Handler(structs.EventContext{
		Payload: payload,
	})
}
