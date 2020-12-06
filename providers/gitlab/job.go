package gitlab

import (
	"fmt"
	"strings"

	//"fmt"
	//"github.com/sirupsen/logrus"
	"github.com/widgetbot-io/gohook/structs"
	"github.com/widgetbot-io/gohook/utils"
	webhook "github.com/widgetbot-io/gohook/webhook/gitlab"
)

func JobHandler(c structs.EventContext) error {
	payload := c.Payload.(webhook.BuildEventPayload)
	tagged := payload.Tag

	// TODO: SetURL is broken or the repo URL from gitlab is broken.
	embed := utils.NewEmbed().
		SetTitle(fmt.Sprintf("[%s:%s]", payload.Repository.Name, payload.Ref)).
		SetURL(fmt.Sprintf("%s/-/jobs/%d", payload.Repository.Homepage, payload.BuildID)).
		SetAuthor(payload.User.Name, payload.User.AvatarURL).
		SetFooter(c.Provider.Logo).
		SetColour(0x00ff00).
		SetTimestamp()

	description := ""
	switch payload.BuildStatus {
	case "failed":
		if payload.BuildAllowFailure {
			return nil
		}

		description = "The job has failed."
		embed.SetColour(0xff0000)
	case "canceled":
		description = "The job has been canceled."
		embed.SetColour(0xffff00)
	case "running":
		{
			embed.SetColour(0x0000ff)

			if strings.HasPrefix(payload.BuildName, "deploy-") {
				environment := strings.TrimPrefix(payload.BuildName, "deploy-")

				if len(environment) != 0 {
					if tagged {
						description = fmt.Sprintf("Version %s is deploying to %s...", payload.Ref, environment)
					} else {
						description = fmt.Sprintf("Deploying latest commit to %s...", environment)
					}
				} else {
					return nil
				}
			} else {
				return nil
			}
		}
	case "success":
		{
			if strings.HasPrefix(payload.BuildName, "deploy-") {
				environment := strings.TrimPrefix(payload.BuildName, "deploy-")

				if len(environment) != 0 {
					if tagged {
						description = fmt.Sprintf("Version %s has been deployed to %s...", payload.Ref, environment)
					} else {
						description = fmt.Sprintf("Deployed latest commit to %s...", environment)
					}
				} else {
					return nil
				}
			} else {
				return nil
			}
		}
	default:
		return nil
	}
	embed.SetDescription(description)

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
