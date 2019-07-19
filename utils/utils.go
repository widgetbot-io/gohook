package utils

import (
	"git.deploys.io/disweb/gohook/structs"
	"strings"
)

func IndexOfProvider(element string, data []structs.Provider) int {
	for k, v := range data {
		if strings.ToLower(element) == strings.ToLower(v.Name) {
			return k
		}
	}
	return -1
}

func IndexOfEvent(element string, data []structs.Event) int {
	for k, v := range data {
		if strings.ToLower(element) == strings.ToLower(v.Name) {
			return k
		}
	}
	return -1
}
