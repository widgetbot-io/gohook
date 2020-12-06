package datadog

import (
	"fmt"
	"github.com/widgetbot-io/gohook/structs"
	"github.com/widgetbot-io/gohook/utils"
)

func MetricHandler(c structs.EventContext) error {
	payload := c.Payload.(structs.DataDogMain)

	embed := utils.NewEmbed().
		SetTitle(fmt.Sprintf("%s", payload.Title)).
		SetFooter(c.Provider.Logo).
		SetAuthor(payload.Org.Name, c.Provider.Logo).
		//SetDescription(payload.Status).
		SetTimestamp().
		SetURL(payload.Link).
		AddField("Tags", payload.Tags, true).
		AddField("Metric", payload.Metric, true).
		SetColour(0x00ff00)

	switch payload.Transition {
	case "Recovered":
		{
			embed.SetColour(0x00ff00)
		}
	case "Triggered":
		{
			embed.SetColour(0xff0000)
		}
	}

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
