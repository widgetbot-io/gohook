package utils

import (
	"bytes"
	"encoding/json"
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

	_, err := http.Post("https://canary.discordapp.com/api/webhooks/586650758607536165/r21LYTWAlVYTQ_EHMBjPlK5f8S66p-IFFMXmv-lOqH-gMD-jD4n2kYO5dAhFtKUGiv98", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}

	return nil
}
