package sonarr

import (
	"git.deploys.io/disweb/gohook/structs"
	"git.deploys.io/disweb/gohook/utils"
)

func DownloadHandler(c structs.EventContext) error {
	payload := c.Payload.(structs.SonarrTest)

	embed := utils.NewEmbed().
		SetTitle(utils.FormatSonarrTitle(payload.Series, payload.Episodes[0])).
		SetFooter(c.Provider.Logo).
		SetAuthor("Test", "https://avatars1.githubusercontent.com/u/1082903?s=400&v=4").
		SetDescription("Beep beep").
		SetColour(utils.RandomColor()).
		SetTimestamp()

	return utils.SendToDiscord(c.ID, c.Secret, embed)
}
