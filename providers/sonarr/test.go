package sonarr

import (
	"git.deploys.io/disweb/gohook/structs"
	"git.deploys.io/disweb/gohook/utils"
)

func TestHandler(c structs.EventContext) error {
	payload := c.Payload.(structs.SonarrTest)

	embed := utils.NewEmbed().
		SetTitle(utils.FormatSonarrTitle(payload.Series, payload.Episodes[0])).
		SetAuthor("Test", c.Provider.Logo).
		SetFooter(c.Provider.Logo).
		SetDescription("Beep beep").
		SetColour(utils.RandomColor()).
		SetTimestamp()

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
