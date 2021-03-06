package datadog

import (
	"github.com/widgetbot-io/gohook/structs"
	"github.com/widgetbot-io/gohook/utils"
)

func SyntheticsHandler(c structs.EventContext) error {
	payload := c.Payload.(structs.DataDogMain)

	embed := utils.NewEmbed().
		SetFooter(c.Provider.Logo).
		SetDescription(payload.Body).
		SetTitle(payload.Title).
		SetTimestamp().
		SetURL("https://google.com").
		SetColour(0x00ff00)

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
