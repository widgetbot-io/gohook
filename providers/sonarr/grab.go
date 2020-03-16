package sonarr

import (
	"lab.venix.dev/widgetbot/gohook/structs"
	"lab.venix.dev/widgetbot/gohook/utils"
)

func GrabHandler(c structs.EventContext) error {
	payload := c.Payload.(structs.SonarrGrab)

	embed := utils.NewEmbed().
		SetColour(utils.RandomColor()).
		SetFooter(c.Provider.Logo).
		SetTimestamp()

	if len(payload.Episodes) <= 1 {
		embed.SetAuthor("Episode Downloading!", c.Provider.Logo)
		embed.SetTitle(utils.FormatSonarrTitle(payload.Series, payload.Episodes[0]))
		embed.AddField("Quality", payload.Episodes[0].Quality, true)
		embed.AddField("Aired On", payload.Episodes[0].AirDate, true)
	} else {
		return nil
	}

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
