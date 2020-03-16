package gitlab

import (
	"github.com/sirupsen/logrus"
	//"fmt"
	//"github.com/sirupsen/logrus"
	webhook "gopkg.in/go-playground/webhooks.v5/gitlab"
	"lab.venix.dev/widgetbot/gohook/structs"
	"lab.venix.dev/widgetbot/gohook/utils"
)

func JobHandler(c structs.EventContext) error {
	payload := c.Payload.(webhook.JobEventPayload)
	// status := ""

	embed := utils.NewEmbed()
	//	SetAuthor(payload.User.Name, payload.User.AvatarURL).
	//	SetTitle(fmt.Sprintf("[%s:%d]", payload.ProjectName, payload.Commit.ID)).
	//	SetFooter(c.Provider.Logo).
	//	SetColour(0x0089ee).
	//	SetURL(fmt.Sprintf("%s/-/jobs/%d", payload.Repository.URL, payload.JobID)).
	//	AddField("a", "b", true).
	//	SetTimestamp()
	//
	//logrus.Info(embed)

	logrus.Info(payload.JobStage)
	if payload.JobStage == "deploy" {
		logrus.Info("depoy")
		if payload.Tag == true {
			logrus.Info("true")
			embed.SetTitle("test").
				SetAuthor(payload.User.Name, c.Provider.Logo).
				SetFooter(c.Provider.Logo).
				SetDescription("Version v0.0.0 is being deployed to xxx...").
				SetColour(0x0000ff).
				SetTimestamp()
		}
	}

	//embed := utils.NewEmbed().
	//	SetTitle("a").
	//	SetAuthor("Test", c.Provider.Logo).
	//	SetFooter(c.Provider.Logo).
	//	SetDescription("Beep beep").
	//	SetColour(utils.RandomColor()).
	//	SetTimestamp()

	//switch payload.ObjectAttributes.Status {
	//case "cancelled":
	//	{
	//		status = "Cancelled"
	//		embed.SetColour(0xFF0000)
	//		embed.AddField(fmt.Sprintf("Pipeline #%s", payload.ObjectAttributes.Ref), fmt.Sprintf("Pipeline has been %s", status), false)
	//	}
	//case "running":
	//	{
	//		status = "Running"
	//		embed.SetColour(0xffff00)
	//		embed.AddField(fmt.Sprintf("Pipeline #%s", payload.ObjectAttributes.Ref), fmt.Sprintf("Pipeline has started %s", status), false)
	//	}
	//case "failed":
	//	{
	//		status = "Failed"
	//		embed.SetColour(0xFF0000)
	//		embed.AddField(fmt.Sprintf("Pipeline #%s", payload.ObjectAttributes.Ref), fmt.Sprintf("Pipeline %s", status), false)
	//	}
	//case "success":
	//	{
	//		status = "Succeeded"
	//		embed.SetColour(0x00ff27)
	//		embed.AddField(fmt.Sprintf("Pipeline #%s", payload.ObjectAttributes.Ref), fmt.Sprintf("Pipeline %s in %d seconds.", status, payload.ObjectAttributes.Duration), false)
	//	}
	//}

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
