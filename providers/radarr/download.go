package radarr

import (
	"github.com/widgetbot-io/gohook/structs"
	"github.com/widgetbot-io/gohook/utils"
)

func DownloadHandler(c structs.EventContext) error {
	payload := c.Payload.(structs.RadarrDownload)

	embed := utils.NewEmbed().
		SetColour(utils.RandomColor()).
		SetFooter(c.Provider.Logo).
		SetTimestamp()

	embed.SetAuthor("Film Downloaded", c.Provider.Logo)
	embed.SetTitle(payload.Movie.Title)
	embed.AddField("Quality", payload.MovieFile.Quality, true)
	embed.AddField("Released on", payload.Movie.ReleaseDate, true)

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
