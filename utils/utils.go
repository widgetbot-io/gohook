package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/webhooks.v5/gitlab"
	"lab.venix.dev/disweb/gohook/structs"
	"net/http"
	"strings"
)

func IndexOfAuthor(element string, data []string) int {
	for k, v := range data {
		if strings.ToLower(element) == strings.ToLower(v) {
			return k
		}
	}
	return -1
}

func EventDetection(data structs.BaseDetection) string {
	if data.ObjectKind != "" {
		return data.ObjectKind
	} else if data.EventType != "" {
		return data.EventType
	}

	return ""
}

func HasOptions(options string, option string) bool {
	// P = Private Git
	return strings.Contains(options, option)
}

func GetBranch(ref string) string {
	return strings.Join(strings.Split(ref, "/")[2:], "/")
}

func GitlabGroupBy(arrayToGroups []gitlab.Commit) map[string][]gitlab.Commit {
	output := make(map[string][]gitlab.Commit)

	for _, v := range arrayToGroups {
		output[v.Author.Name] = append(output[v.Author.Name], v)
	}

	return output
}

func GithubGroupBy(arrayToGroups []struct {
	Sha       string `json:"sha"`
	ID        string `json:"id"`
	NodeID    string `json:"node_id"`
	TreeID    string `json:"tree_id"`
	Distinct  bool   `json:"distinct"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	URL       string `json:"url"`
	Author    struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"author"`
	Committer struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"committer"`
	Added    []string `json:"added"`
	Removed  []string `json:"removed"`
	Modified []string `json:"modified"`
}) map[string][]structs.GithubCommit { // Yes i know, blame whoever made the github thing non modular.
	output := make(map[string][]structs.GithubCommit)

	for _, v := range arrayToGroups {
		output[v.Author.Name] = append(output[v.Author.Name], v)
	}

	return output
}

type Embeds struct {
	Embeds []*Embed `json:"embeds"`
}

func SendToDiscord(ID string, secret string, embed *Embed, options string) error {
	if HasOptions(options, "P") {
		embed.URL = ""
	}

	embeds := Embeds{}
	embeds.Embeds = []*Embed{embed}

	jsonBytes, _ := json.Marshal(embeds)

	_, err := http.Post(fmt.Sprintf("https://canary.discordapp.com/api/webhooks/%s/%s", ID, secret), "application/json", bytes.NewBuffer(jsonBytes))

	return err
}

func FormatSonarrTitle(series structs.SonarrSeries, episode structs.SonarrEpisode) string {
	var e string
	var s string

	if episode.EpisodeNumber >= 10 {
		e = fmt.Sprintf("%d", episode.EpisodeNumber)
	} else {
		e = fmt.Sprintf("0%d", episode.EpisodeNumber)
	}

	if episode.SeasonNumber >= 10 {
		s = fmt.Sprintf("%d", episode.SeasonNumber)
	} else {
		s = fmt.Sprintf("0%d", episode.SeasonNumber)
	}

	return fmt.Sprintf("[%s] S%sE%s - %s", series.Title, e, s, episode.Title)
}
