package github

import (
	"git.deploys.io/disweb/gohook/structs"
	"git.deploys.io/disweb/gohook/utils"
	webhook "gopkg.in/go-playground/webhooks.v5/github"
)

func PushHandler(c structs.EventContext) error {
	payload := c.Payload.(webhook.PushPayload)

	embed := utils.NewEmbed().
		SetTitle("Ping event recieved!").
		SetAuthor(payload.Sender.Login, payload.Sender.AvatarURL).
		SetDescription("Webhook successfully linked to repository.").
		SetURL(payload.Repository.HTMLURL).
		SetColour(utils.RandomColor()).
		SetTimestamp()

	return utils.SendToDiscord(c.ID, c.Secret, embed)
}