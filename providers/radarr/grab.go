package radarr

import (
	"lab.venix.dev/disweb/gohook/structs"
	"lab.venix.dev/disweb/gohook/utils"
)

func GrabHandler(c structs.EventContext) error {
	payload := c.Payload.(structs.RadarrGrab)

	embed := utils.NewEmbed().
		SetColour(utils.RandomColor()).
		SetFooter(c.Provider.Logo).
		SetTimestamp()

	embed.SetAuthor("Film Downloading!", c.Provider.Logo)
	embed.SetTitle(payload.Movie.Title)
	embed.AddField("Quality", payload.Release.Quality, true)
	embed.AddField("Released on", payload.Movie.ReleaseDate, true)

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
