package datadog

import (
	"encoding/json"
	"lab.venix.dev/widgetbot/gohook/structs"
)

func Handler(c structs.ProviderContext) error {
	var payload structs.DataDogMain
	_ = json.Unmarshal([]byte(c.Payload), &payload)

	println("-----PRINTING BODY-----")
	println(payload.Body)
	println("----------------")
	println(payload.LastUpdated)
	println("----------------")
	println(payload.Date)
	println("----------------")
	println(payload.Org.Id)
	println("----------------")
	println(payload.Org.Name)
	println("----------------")
	println(payload.Id)
	println("----------------")
	println(payload.EventType)
	println("----------------")
	println(payload.Msg)
	println("----------------")
	println(payload.Status)
	println("----------------")
	println(payload.Title)
	println("----------------")

	return c.Event.Handler(structs.EventContext{
		ID:       c.ID,
		Secret:   c.Secret,
		Options:  c.Options,
		Event:    c.Event,
		Provider: c.Provider,
		Payload:  payload,
	})
}
