package sonarr

import (
	"github.com/sirupsen/logrus"
	"lab.venix.dev/disweb/gohook/structs"
	"lab.venix.dev/disweb/gohook/utils"
)

func TestHandler(c structs.EventContext) error {
	logrus.Info(c.Payload)
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
