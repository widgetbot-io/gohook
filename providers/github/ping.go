package github

import (
	webhook "gopkg.in/go-playground/webhooks.v5/github"
	"github.com/widgetbot-io/gohook/structs"
	"github.com/widgetbot-io/gohook/utils"
)

func PingHandler(c structs.EventContext) error {
	payload := c.Payload.(webhook.PingPayload)

	embed := utils.NewEmbed().
		SetTitle("Ping event recieved!").
		SetAuthor(payload.Sender.Login, payload.Sender.AvatarURL).
		SetFooter(c.Provider.Logo).
		SetDescription("Webhook successfully linked to repository.").
		SetURL(payload.Repository.HTMLURL).
		SetColour(utils.RandomColor()).
		SetTimestamp()

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
