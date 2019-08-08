package gitlab

import (
	"fmt"
	"git.deploys.io/disweb/gohook/structs"
	"git.deploys.io/disweb/gohook/utils"
	webhook "gopkg.in/go-playground/webhooks.v5/gitlab"
)

func NoteHandler(c structs.EventContext) error {
	var title string = ""

	payload := c.Payload.(webhook.CommentEventPayload)

	embed := utils.NewEmbed().
		SetFooter("https://upload.wikimedia.org/wikipedia/commons/thumb/1/18/GitLab_Logo.svg/1108px-GitLab_Logo.svg.png").
		SetColour(0xff9700).
		SetTimestamp().
		SetURL(payload.Project.WebURL).
		SetAuthor(payload.User.Name, payload.User.AvatarURL)

	switch payload.ObjectAttributes.NotebookType {
	case "Commit":
		{
			title = fmt.Sprintf("Commit (%s)", payload.Commit.ID[:7])
			embed.SetDescription(payload.ObjectAttributes.Note)
			embed.SetURL(fmt.Sprintf("%s/commits/%s", payload.Project.WebURL, payload.Commit.ID[:7]))
		}
	case "MergeRequest":
		{
			title = fmt.Sprintf("Merge Request #%d", payload.MergeRequest.IID)
			embed.SetDescription(payload.ObjectAttributes.Note)
			embed.SetURL(fmt.Sprintf("%s/merge_requests/%d", payload.Project.WebURL, payload.MergeRequest.IID))
		}
	case "Issue":
		{
			title = fmt.Sprintf("Issue #%d", payload.Issue.IID)
			embed.SetDescription(payload.ObjectAttributes.Note)
			embed.SetURL(fmt.Sprintf("%s/issues/%d", payload.Project.WebURL, payload.Issue.IID))
		}
	case "Snippet":
		{
			title = fmt.Sprintf("Snippet #%d", payload.Snippet.ID)
			embed.SetDescription(payload.ObjectAttributes.Note)
			embed.SetURL(fmt.Sprintf("%s/snippets/%d", payload.Project.WebURL, payload.Snippet.ID))
		}
	}
	embed.SetTitle(fmt.Sprintf("[%s] - Commented on %s", payload.Project.Name, title))

	return utils.SendToDiscord(c.ID, c.Secret, embed)
}
