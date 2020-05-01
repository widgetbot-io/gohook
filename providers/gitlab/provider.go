package gitlab

import (
	"lab.venix.dev/widgetbot/gohook/structs"
	webhook "lab.venix.dev/widgetbot/gohook/webhook/gitlab"
)

func Handler(c structs.ProviderContext) error {
	hook, _ := webhook.New()
	payload, _ := hook.Parse(c.Context.Request, webhook.PushEvents, webhook.TagEvents, webhook.IssuesEvents, webhook.CommentEvents, webhook.SystemHookEvents, webhook.PipelineEvents, webhook.JobEvents)
	// event will be found because we find it earlier

	return c.Event.Handler(structs.EventContext{
		ID:       c.ID,
		Secret:   c.Secret,
		Options:  c.Options,
		Event:    c.Event,
		Provider: c.Provider,
		Payload:  payload,
	})
}
