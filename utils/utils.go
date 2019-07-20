package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	_, err := http.Post(fmt.Sprintf("https://canary.discordapp.com/api/webhooks/%s/%s", ID, secret), "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}

	return nil
}
