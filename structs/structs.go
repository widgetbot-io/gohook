package structs

import "github.com/gin-gonic/gin"

type Event struct {
	Handler func(c EventContext) error
}

type Provider struct {
	Name      string
	Header    string
	EventName string
	Events    map[string]Event
	Handler   func(c ProviderContext) error
}

type ProviderContext struct {
	ID       string
	Secret   string
	Provider Provider
	Event    Event
	Context  *gin.Context
	Payload  []byte
}

type EventContext struct {
	ID       string
	Secret   string
	Provider Provider
	Event    Event
	Context  *gin.Context
	Payload  interface{}
}

type BaseDetection struct {
	EventType  string `json:"eventType,omitempty"`
	ObjectKind string `json:"object_kind,omitempty"`
}
