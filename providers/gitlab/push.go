package gitlab

import (
	"fmt"
	"lab.venix.dev/widgetbot/gohook/structs"
	"lab.venix.dev/widgetbot/gohook/utils"
	webhook "lab.venix.dev/widgetbot/gohook/webhook/gitlab"
	"strings"
)

func BranchHandler(c structs.EventContext, new bool) error {
	payload := c.Payload.(webhook.PushEventPayload)
	branch := utils.GetBranch(payload.Ref)

	embed := utils.NewEmbed().
		SetFooter(c.Provider.Logo).
		SetTimestamp().
		SetURL(payload.Project.WebURL).
		SetAuthor(payload.UserName, payload.UserAvatar)

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
	payload := c.Payload.(webhook.PushEventPayload)
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
		SetTitle(fmt.Sprintf("[%s:%s] %s", payload.Project.Name, branch, commit)).
		SetColour(0x0089ee).
		SetFooter(c.Provider.Logo).
		SetTimestamp().
		SetURL(payload.Project.WebURL).
		SetAuthor(payload.UserName, payload.UserAvatar)

	groups := utils.GitlabGroupBy(payload.Commits)

	limit := 25
	for k := range groups {
		limit = limit - 1
		if limit == 0 {
			break
		}

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
				commitString += fmt.Sprintf("[%s](%s) - %s \n", b.ID[:7], b.URL, commitMessage[:256])
			} else {
				commitString += fmt.Sprintf("`%s` - %s \n", b.ID[:7], commitMessage[:256])
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
