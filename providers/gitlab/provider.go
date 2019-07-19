package gitlab

import (
	"git.deploys.io/disweb/gohook/structs"
	log "github.com/sirupsen/logrus"
	webhook "gopkg.in/go-playground/webhooks.v5/gitlab"
)

func Handler(c structs.ProviderContext) error {
	hook, _ := webhook.New()
	payload, _ := hook.Parse(c.Context.Request, webhook.PushEvents)
	// event will be found because we find it earlier

	switch payload.(type) {
	case webhook.PushEventPayload:
		release := payload.(webhook.PushEventPayload)
		c.Event.Handler(structs.EventContext{})
	}

	log.Info(c.ID)
	return nil
}
