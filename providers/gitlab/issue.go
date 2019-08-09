package gitlab

import (
	"fmt"
	"git.deploys.io/disweb/gohook/structs"
	"git.deploys.io/disweb/gohook/utils"
	webhook "gopkg.in/go-playground/webhooks.v5/gitlab"
	"strings"
)

func IssueHandler(c structs.EventContext) error {
	payload := c.Payload.(webhook.IssueEventPayload)

	embed := utils.NewEmbed().
		SetTitle(fmt.Sprintf("%s issue #%d on %s", strings.Title(payload.ObjectAttributes.State), payload.ObjectAttributes.IID, payload.Project.Name)).
		SetFooter(c.Provider.Logo).
		SetURL(fmt.Sprintf("%s/issues/%d", payload.Project.WebURL, payload.ObjectAttributes.IID)).
		SetAuthor(payload.User.Name, payload.User.AvatarURL).
		SetTimestamp()

	if payload.ObjectAttributes.Description != "" {
		if len(payload.ObjectAttributes.Description) > 1024 {
			embed.AddField(payload.ObjectAttributes.Title, fmt.Sprintf("%s \u2026", payload.ObjectAttributes.Description[:1024]), false)
		} else {
			embed.AddField(payload.ObjectAttributes.Title, payload.ObjectAttributes.Description, false)
		}
	} else {
		embed.AddField(payload.ObjectAttributes.Title, "This issue has no description..", false)
	}

	switch payload.ObjectAttributes.State {
	case "opened":
		{
			embed.SetColour(0x36ff00)
		}
	case "closed":
		{
			embed.SetColour(0xff0000)
		}
	}

	return utils.SendToDiscord(c.ID, c.Secret, embed)
}
