package gitlab

import (
	"fmt"
	"git.deploys.io/disweb/gohook/structs"
	"git.deploys.io/disweb/gohook/utils"
	"github.com/sirupsen/logrus"
	webhook "gopkg.in/go-playground/webhooks.v5/gitlab"
	"strings"
)

func PushHandler(c structs.EventContext) error {
	payload := c.Payload.(webhook.PushEventPayload)
	branch := utils.GetBranch(payload.Ref)

	commit := ""
	if len(payload.Commits) == 1 {
		commit = "1 commit"
	} else {
		commit = fmt.Sprintf("%d commits", len(payload.Commits))
	}

	embed := utils.NewEmbed().
		SetTitle(fmt.Sprintf("[%s:%s] %s", payload.Project.Name, branch, commit)).
		SetColour(0x0089ee).
		SetURL(payload.Project.WebURL).
		SetAuthor(payload.UserName, payload.UserAvatar)

	// var reduced []webhook.Commit
	var authors []string

	for _, commit := range payload.Commits {
		if utils.IndexOfAuthor(commit.Author.Name, authors) == -1 {
			authors = append(authors, commit.Author.Name)
		}
		for _, author := range authors {
			if commit.Author.Name == author {

			}
		}
	}

	for _, commit := range payload.Commits {
		commitString := ""
		if strings.HasPrefix(commit.Message, "!") || strings.HasPrefix(commit.Message, "$") {
			commitString = "This commit message has been marked as private."
		} else {
			commitString = commit.Message
		}

		embed.AddField(fmt.Sprintf("Commit from %s", commit.Author.Name), fmt.Sprintf("[`%s`](%s) %s", commit.ID[:7], commit.URL, commitString), false)
	}

	logrus.Info(len(payload.Commits))
	logrus.Info(len(embed.Fields))

	return utils.SendToDiscord(c.ID, c.Secret, embed)
}
