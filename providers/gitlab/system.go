package gitlab

import (
	"github.com/widgetbot-io/gohook/structs"
	webhook "github.com/widgetbot-io/gohook/webhook/gitlab"
)

func SystemHandler(c structs.EventContext) error {
	payload := c.Payload.(webhook.SystemHookPayload)

	switch payload.ObjectKind {

	}

	return nil
}
