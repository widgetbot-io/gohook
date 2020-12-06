package gitlab

import (
	"fmt"
	"github.com/widgetbot-io/gohook/structs"
	"github.com/widgetbot-io/gohook/utils"
	webhook "github.com/widgetbot-io/gohook/webhook/gitlab"
)

func NoteHandler(c structs.EventContext) error {
	var title string = ""

	payload := c.Payload.(webhook.CommentEventPayload)

	embed := utils.NewEmbed().
		SetFooter(c.Provider.Logo).
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

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
