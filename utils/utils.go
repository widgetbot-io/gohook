package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git.deploys.io/disweb/gohook/structs"
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

func GetBranch(ref string) string {
	return strings.Join(strings.Split(ref, "/")[2:], "/")
}

func GroupBy(arrayToGroups []string, key string, key2 string) []string {
	return []string{""}
}

type Embeds struct {
	Embeds []*Embed `json:"embeds"`
}

func SendToDiscord(ID string, secret string, embed *Embed) error {
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
