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
			switch payload.BuildStage {
			case "deploy":
				{
					if tagged {
						description = "Version %s is deploying..."
						description = fmt.Sprintf(description, payload.Ref)
					} else {
						description = "The latest commit is deploying..."
					}
				}
			case "production":
				{
					if tagged {
						description = "Version %s is deploying to production..."
						description = fmt.Sprintf(description, payload.Ref)
					} else {
						description = "The latest commit is deploying to production..."
					}
				}
			case "staging":
				{
					if tagged {
						description = "Version %s is deploying to staging..."
						description = fmt.Sprintf(description, payload.Ref)
					} else {
						description = "The latest commit is deploying to staging..."
					}
				}
			default:
				description = "A job is running..."
			}
		}
	case "success":
		{
			switch payload.BuildStage {
			case "deploy":
				{
					if tagged {
						description = "Version %s has deployed..."
						description = fmt.Sprintf(description, payload.Ref)
					} else {
						description = "The latest commit has deployed..."
					}
				}
			case "production":
				{
					if tagged {
						description = "Version %s has deployed to production..."
						description = fmt.Sprintf(description, payload.Ref)
					} else {
						description = "The latest commit has deployed to production..."
					}
				}
			case "staging":
				{
					if tagged {
						description = "Version %s has deployed to staging..."
						description = fmt.Sprintf(description, payload.Ref)
					} else {
						description = "The latest commit has deployed to staging..."
					}
				}
			default:
				description = "A job is running..."
			}
		}
	default:
		return nil
	}
	embed.SetDescription(description)

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
