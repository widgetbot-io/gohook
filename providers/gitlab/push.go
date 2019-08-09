package gitlab

import (
	"fmt"
	"git.deploys.io/disweb/gohook/structs"
	"git.deploys.io/disweb/gohook/utils"
	webhook "gopkg.in/go-playground/webhooks.v5/gitlab"
	"strings"
)

func PushHandler(c structs.EventContext) error {
	payload := c.Payload.(webhook.PushEventPayload)
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
		SetTitle(fmt.Sprintf("[%s:%s] %s", payload.Project.Name, branch, commit)).
		SetColour(0x0089ee).
		SetFooter(c.Provider.Logo).
		SetTimestamp().
		SetURL(payload.Project.WebURL).
		SetAuthor(payload.UserName, payload.UserAvatar)

	groups := utils.GitlabGroupBy(payload.Commits)

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

			commitString += fmt.Sprintf("[%s](%s) - %s \n", b.ID[:7], b.URL, commitMessage)
		}

		commit := "Commit"
		if len(group) > 1 {
			commit = "Commits"
		}

		embed.AddField(fmt.Sprintf("%s from %s", commit, k), commitString, false)
	}

	return utils.SendToDiscord(c.ID, c.Secret, embed)
}
