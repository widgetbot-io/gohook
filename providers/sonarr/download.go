package sonarr

import (
	"lab.venix.dev/disweb/gohook/structs"
	"lab.venix.dev/disweb/gohook/utils"
)

func DownloadHandler(c structs.EventContext) error {
	payload := c.Payload.(structs.SonarrDownload)

	embed := utils.NewEmbed().
		SetColour(utils.RandomColor()).
		SetFooter(c.Provider.Logo).
		SetTimestamp()

	if len(payload.Episodes) <= 1 {
		embed.SetAuthor("Episode Downloaded", c.Provider.Logo)
		embed.SetTitle(utils.FormatSonarrTitle(payload.Series, payload.Episodes[0]))
		embed.AddField("Quality", payload.Episodes[0].Quality, true)
		embed.AddField("Aired On", payload.Episodes[0].AirDate, true)
	} else {
		return nil
	}

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
