package gitlab

import (
	"github.com/widgetbot-io/gohook/structs"
	webhook "github.com/widgetbot-io/gohook/webhook/gitlab"
)

func Handler(c structs.ProviderContext) error {
	hook, _ := webhook.New()
	payload, _ := hook.Parse(c.Context.Request, webhook.PushEvents, webhook.TagEvents, webhook.IssuesEvents, webhook.CommentEvents, webhook.SystemHookEvents, webhook.PipelineEvents, webhook.JobEvents, webhook.BuildEvents)
	// event will be found because we find it earlier
	println(payload)

	return c.Event.Handler(structs.EventContext{
		ID:       c.ID,
		Secret:   c.Secret,
		Options:  c.Options,
		Event:    c.Event,
		Provider: c.Provider,
		Payload:  payload,
	})
}
