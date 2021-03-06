package github

import (
	"fmt"
	webhook "gopkg.in/go-playground/webhooks.v5/github"
	"github.com/widgetbot-io/gohook/structs"
	"github.com/widgetbot-io/gohook/utils"
	"strings"
)

func BranchHandler(c structs.EventContext, new bool) error {
	payload := c.Payload.(webhook.PushPayload)
	branch := utils.GetBranch(payload.Ref)

	embed := utils.NewEmbed().
		SetFooter(c.Provider.Logo).
		SetTimestamp().
		SetURL(payload.Repository.HTMLURL).
		SetAuthor(payload.Pusher.Name, payload.Sender.AvatarURL)

	if new {
		embed.SetTitle(fmt.Sprintf("Branch created: %s", branch)).
			SetColour(0x00ff00)
	} else {
		embed.SetTitle(fmt.Sprintf("Branch deleted: %s", branch)).
			SetColour(0xff0000)
	}

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}

func PushHandler(c structs.EventContext) error {
	payload := c.Payload.(webhook.PushPayload)
	branch := utils.GetBranch(payload.Ref)

	if strings.HasPrefix(branch, "!") || strings.HasPrefix(branch, "$") {
		return nil
	}

	if payload.Before == "0000000000000000000000000000000000000000" {
		return BranchHandler(c, true)
	} else if payload.After == "0000000000000000000000000000000000000000" {
		return BranchHandler(c, false)
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
		SetFooter(c.Provider.Logo).
		SetTimestamp().
		SetURL(payload.Repository.HTMLURL).
		SetAuthor(payload.Pusher.Name, payload.Sender.AvatarURL)

	groups := utils.GithubGroupBy(payload.Commits)

	for k := range groups {
		group := groups[k]

		commitString := ""
		for _, b := range group {
			commitMessage := ""

			if strings.HasPrefix(b.Message, "!") || strings.HasPrefix(b.Message, "$") {
				commitMessage = "This commit message has been marked as private."
			} else {
				commitMessage = b.Message
			}

			if !utils.HasOptions(c.Options, "P") {
				commitString += fmt.Sprintf("[%s](%s) - %s \n", b.ID[:7], b.URL, commitMessage)
			} else {
				commitString += fmt.Sprintf("`%s` - %s \n", b.ID[:7], commitMessage)
			}
		}

		commit := "Commit"
		if len(group) > 1 {
			commit = "Commits"
		}

		embed.AddField(fmt.Sprintf("%s from %s", commit, k), commitString, false)
	}

	return utils.SendToDiscord(c.ID, c.Secret, embed, c.Options)
}
