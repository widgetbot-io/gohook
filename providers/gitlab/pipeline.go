package gitlab

import (
	"fmt"
	"lab.venix.dev/widgetbot/gohook/structs"
	"lab.venix.dev/widgetbot/gohook/utils"
	webhook "lab.venix.dev/widgetbot/gohook/webhook/gitlab"
)

func PipelineHandler(c structs.EventContext) error {
	payload := c.Payload.(webhook.PipelineEventPayload)
	status := ""

	embed := utils.NewEmbed().
		SetAuthor(payload.User.Name, payload.User.AvatarURL).
		SetTitle(fmt.Sprintf("[%s:%s]", payload.Project.Name, payload.ObjectAttributes.Ref)).
		SetFooter(c.Provider.Logo).
		SetColour(0x0089ee).
		SetURL(fmt.Sprintf("%s/pipelines/%s", payload.Project.WebURL, payload.ObjectAttributes.ID)).
		SetTimestamp()

	switch payload.ObjectAttributes.Status {
	case "cancelled":
		{
			status = "Cancelled"
			embed.SetColour(0xFF0000)
			embed.AddField(fmt.Sprintf("Pipeline #%s", payload.ObjectAttributes.Ref), fmt.Sprintf("Pipeline has been %s", status), false)
		}
	case "running":
		{
			status = "Running"
			embed.SetColour(0xffff00)
			embed.AddField(fmt.Sprintf("Pipeline #%s", payload.ObjectAttributes.Ref), fmt.Sprintf("Pipeline has started %s", status), false)
		}
	case "failed":
		{
			status = "Failed"
			embed.SetColour(0xFF0000)
			embed.AddField(fmt.Sprintf("Pipeline #%s", payload.ObjectAttributes.Ref), fmt.Sprintf("Pipeline %s", status), false)
		}
	case "success":
		{
			status = "Succeeded"
			embed.SetColour(0x00ff27)
			embed.AddField(fmt.Sprintf("Pipeline #%s", payload.ObjectAttributes.Ref), fmt.Sprintf("Pipeline %s in %d seconds.", status, payload.ObjectAttributes.Duration), false)
		}

	default:
		{
			return nil
		}
	}

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
