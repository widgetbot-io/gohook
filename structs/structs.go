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

type GithubCommit struct {
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
}
