package gitlab

import (
	"fmt"
	"git.deploys.io/disweb/gohook/structs"
	"git.deploys.io/disweb/gohook/utils"
	webhook "gopkg.in/go-playground/webhooks.v5/gitlab"
)

func TagHandler(c structs.EventContext) error {
	payload := c.Payload.(webhook.TagEventPayload)
	tag := utils.GetBranch(payload.Ref)

	embed := utils.NewEmbed().
		SetFooter("Gohook", "https://upload.wikimedia.org/wikipedia/commons/thumb/1/18/GitLab_Logo.svg/1108px-GitLab_Logo.svg.png").
		SetTimestamp().
		SetURL(payload.Project.WebURL).
		SetAuthor(payload.UserName, payload.UserAvatar)

	if payload.Before == "0000000000000000000000000000000000000000" {
		embed.SetTitle(fmt.Sprintf("[%s] - Tag Created: %s", payload.Project.Name, tag)).
			SetColour(0x00ff00)
	} else if payload.After == "0000000000000000000000000000000000000000" {
		embed.SetTitle(fmt.Sprintf("[%s] - Tag Removed: %s", payload.Project.Name, tag)).
			SetColour(0xff0000)
	}

	return utils.SendToDiscord(c.ID, c.Secret, embed)
}
