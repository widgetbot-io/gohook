package structs

import "github.com/gin-gonic/gin"

type Event struct {
	Name    string
	Handler func(c Context)
}

type Provider struct {
	Name    string
	Header  string
	Events  []Event
	Handler func(c Context) error
}

type Context struct {
	ID       string
	Secret   string
	Provider Provider
	Event    Event
	Context  *gin.Context
}
