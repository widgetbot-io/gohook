package github

import (
	"fmt"
	"git.deploys.io/disweb/gohook/structs"
	"git.deploys.io/disweb/gohook/utils"
	webhook "gopkg.in/go-playground/webhooks.v5/github"
	"strings"
)

func PushHandler(c structs.EventContext) error {
	payload := c.Payload.(webhook.PushPayload)
	branch := utils.GetBranch(payload.Ref)

	if strings.HasPrefix(branch, "!") || strings.HasPrefix(branch, "$") {
		return nil
	}

	commit := ""
	if len(payload.Commits) == 1 {
		commit = "1 commit"
	} else {
		commit = fmt.Sprintf("%d commits", len(payload.Commits))
	}

	embed := utils.NewEmbed().
		SetTitle(fmt.Sprintf("[%s:%s] %s", payload.Repository.Name, branch, commit)).
		SetColour(0x0089ee).
		SetFooter("Gohook", "https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png").
		SetTimestamp().
		SetURL(payload.Repository.HTMLURL).
		SetAuthor(payload.Pusher.Name, payload.Sender.AvatarURL)

	for _, commit := range payload.Commits {
		commitString := ""
		if strings.HasPrefix(commit.Message, "!") || strings.HasPrefix(commit.Message, "$") {
			commitString = "This commit message has been marked as private."
		} else {
			commitString = commit.Message
		}

		embed.AddField(fmt.Sprintf("Commit from %s", commit.Author.Name), fmt.Sprintf("[`%s`](%s) %s", commit.ID[:7], commit.URL, commitString), false)
	}

	return utils.SendToDiscord(c.ID, c.Secret, embed)
}
