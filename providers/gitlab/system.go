package gitlab

import (
	"lab.venix.dev/widgetbot/gohook/structs"
	webhook "lab.venix.dev/widgetbot/gohook/webhook/gitlab"
)

func SystemHandler(c structs.EventContext) error {
	payload := c.Payload.(webhook.SystemHookPayload)

	switch payload.ObjectKind {

	}

	return nil
}
