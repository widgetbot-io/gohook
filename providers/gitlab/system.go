package gitlab

import (
	webhook "gopkg.in/go-playground/webhooks.v5/gitlab"
	"lab.venix.dev/disweb/gohook/structs"
)

func SystemHandler(c structs.EventContext) error {
	payload := c.Payload.(webhook.SystemHookPayload)

	switch payload.ObjectKind {

	}

	return nil
}
