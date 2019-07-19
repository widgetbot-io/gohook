package structs

type Event struct {
	Name    string
	Handler func(c Context)
}

type Provider struct {
	Name   string
	Header string
	Events []Event
}

type Context struct {
}
