package github

import (
	"fmt"
	webhook "gopkg.in/go-playground/webhooks.v5/github"
	"lab.venix.dev/widgetbot/gohook/structs"
	"lab.venix.dev/widgetbot/gohook/utils"
)

func RepositoryHandler(c structs.EventContext) error {
	payload := c.Payload.(webhook.RepositoryPayload)

	embed := utils.NewEmbed().
		SetAuthor(payload.Sender.Login, payload.Sender.AvatarURL).
		SetFooter(c.Provider.Logo).
		SetURL(payload.Repository.HTMLURL).
		SetTimestamp()

	switch payload.Action {
	case "created":
		embed.SetTitle(fmt.Sprintf("[%s] New repository created", payload.Repository.FullName)).
			SetColour(0x00ff00)
	case "deleted":
		embed.SetTitle(fmt.Sprintf("[%s] Repository was deleted", payload.Repository.FullName)).
			SetColour(0xff0000)
	case "archived":
		embed.SetTitle(fmt.Sprintf("[%s] Repository was archived", payload.Repository.FullName)).
			SetColour(0xffff00)
	case "unarchived":
		embed.SetTitle(fmt.Sprintf("[%s] Repository was unarchived", payload.Repository.FullName)).
			SetColour(0xffff00)
	case "edited":
		embed.SetTitle(fmt.Sprintf("[%s] Repository was edited", payload.Repository.FullName)).
			SetColour(0xffa500)
	case "renamed":
		embed.SetTitle(fmt.Sprintf("[%s] Repository was renamed", payload.Repository.FullName)).
			SetColour(0xffa500)
	case "transferred":
		embed.SetTitle(fmt.Sprintf("[%s] Repository was transferred", payload.Repository.FullName)).
			SetColour(0xff0000)
	case "publicized":
		embed.SetTitle(fmt.Sprintf("[%s] Repository was made public", payload.Repository.FullName)).
			SetColour(0x00ff00)
	case "privatized":
		embed.SetTitle(fmt.Sprintf("[%s] Repository was made private", payload.Repository.FullName)).
			SetColour(0x000000)
	}

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
