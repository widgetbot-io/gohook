package radarr

import (
	"lab.venix.dev/disweb/gohook/structs"
	"lab.venix.dev/disweb/gohook/utils"
)

func TestHandler(c structs.EventContext) error {
	embed := utils.NewEmbed().
		SetTitle("Radarr test!").
		SetAuthor("Test", c.Provider.Logo).
		SetFooter(c.Provider.Logo).
		SetDescription("Beep beep").
		SetColour(utils.RandomColor()).
		SetTimestamp()

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
