package structs

import "github.com/gin-gonic/gin"

type Event struct {
	Handler func(c EventContext) error
}

type Provider struct {
	Name    string
	Header  string
	Events  map[string]Event
	Handler func(c ProviderContext) error
}

type ProviderContext struct {
	ID       string
	Secret   string
	Provider Provider
	Event    Event
	Context  *gin.Context
}

type EventContext struct {
	ID       string
	Secret   string
	Provider Provider
	Event    Event
	Context  *gin.Context
	Payload  interface{}
}
