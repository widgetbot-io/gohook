package gitlab

import (
	"fmt"
	//"fmt"
	//"github.com/sirupsen/logrus"
	"lab.venix.dev/widgetbot/gohook/structs"
	"lab.venix.dev/widgetbot/gohook/utils"
	webhook "lab.venix.dev/widgetbot/gohook/webhook/gitlab"
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
		description = "The job has failed."
		embed.SetColour(0xff0000)
	case "canceled":
		description = "The job has been canceled."
		embed.SetColour(0xffff00)
	case "running":
		{
			embed.SetColour(0x0000ff)
			switch payload.BuildStage {
			case "production":
				{
					if tagged {
						description = fmt.Sprintf("Version %s is deploying to production...", payload.Ref)
					} else {
						description = "Deploying latest commit to production..."
					}
				}
			case "staging":
				{
					if tagged {
						description = fmt.Sprintf("Version %s is deploying to staging...", payload.Ref)
					} else {
						description = "Deploying latest commit to staging..."
					}
				}
			}
		}
	case "success":
		{
			switch payload.BuildStage {
			case "production":
				{
					if tagged {
						description = fmt.Sprintf("Version %s has deployed to production...", payload.Ref)
					} else {
						description = "The latest commit is deployed to production."
					}
				}
			case "staging":
				{
					if tagged {
						description = fmt.Sprintf("Deployed %s to staging!", payload.Ref)
					} else {
						description = "Deployed latest commit to staging!"
					}
				}
			}
		}
	default:
		return nil
	}
	embed.SetDescription(description)

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
