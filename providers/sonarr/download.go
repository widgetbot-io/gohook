package sonarr

import (
	"git.deploys.io/disweb/gohook/structs"
	"git.deploys.io/disweb/gohook/utils"
)

func DownloadHandler(c structs.EventContext) error {
	payload := c.Payload.(structs.SonarrDownload)

	embed := utils.NewEmbed().
		SetColour(utils.RandomColor()).
		SetFooter(c.Provider.Logo).
		SetTimestamp()

	if len(payload.Episodes) <= 1 {
		embed.SetAuthor("Episode Downloaded!", "https://avatars1.githubusercontent.com/u/1082903?s=400&v=4")
		embed.SetTitle(utils.FormatSonarrTitle(payload.Series, payload.Episodes[0]))
		embed.AddField("Quality", payload.Episodes[0].Quality, true)
		embed.AddField("Aired On", payload.Episodes[0].AirDate, true)
	} else {
		return nil
	}

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
