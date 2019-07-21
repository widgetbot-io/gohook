package sonarr

import (
	"git.deploys.io/disweb/gohook/structs"
	"git.deploys.io/disweb/gohook/utils"
	"github.com/sirupsen/logrus"
)

func GrabHandler(c structs.EventContext) error {
	payload := c.Payload.(structs.SonarrTest)

	embed := utils.NewEmbed().
		SetColour(utils.RandomColor()).
		SetTimestamp()

	if len(payload.Episodes) <= 1 {
		logrus.Info(payload.Episodes[0].Title)
		embed.SetAuthor("Episode Downloading!", "https://avatars1.githubusercontent.com/u/1082903?s=400&v=4")
		embed.SetTitle(utils.FormatSonarrTitle(payload.Series, payload.Episodes[0]))
		embed.AddField("Quality", payload.Episodes[0].Quality, true)
		embed.AddField("Aired On", payload.Episodes[0].AirDate, true)
	}

	embed.SetDescription("TBC")

	return utils.SendToDiscord(c.ID, c.Secret, embed)
}