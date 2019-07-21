package main

import (
	"encoding/json"
	"git.deploys.io/disweb/gohook/providers/gitlab"
	"git.deploys.io/disweb/gohook/providers/sonarr"
	"git.deploys.io/disweb/gohook/structs"
	"git.deploys.io/disweb/gohook/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var EventCount int
var ProviderList string

var Providers = map[string]structs.Provider{}

func main() {
	version := "v0.0.1"
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: false,
	})
	log.WithFields(log.Fields{
		"version": version,
	}).Info("Loading Application...")

	setupRoutes(router)
	loadProviders()

	log.WithFields(log.Fields{
		"version":   version,
		"providers": len(Providers),
		"events":    EventCount,
	}).Info("Loaded providers.")

	log.WithFields(log.Fields{
		"version": version,
	}).Info("Loaded Application!")
	_ = router.Run(":8443")
}

func setupRoutes(router *gin.Engine) {
	router.POST("/api/hook/:ID/:Secret/:Provider", func(c *gin.Context) {
		var event structs.Event
		var eventName string
		var provider structs.Provider
		var BaseDetection structs.BaseDetection

		payload, _ := utils.Parse(c.Request)
		_ = json.Unmarshal([]byte(payload), &BaseDetection)
		idParam := c.Param("ID")
		secretParam := c.Param("Secret")
		providerParam := c.Param("Provider")
		provider = Providers[providerParam]

		if provider.Name == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Provider not found", "provider": providerParam})
			return
		}

		if provider.Header != "" {
			event = provider.Events[c.GetHeader(provider.Header)]
			eventName = c.GetHeader(provider.Header)
		} else {
			event = provider.Events[utils.EventDetection(BaseDetection)]
			eventName = utils.EventDetection(BaseDetection)
		}

		if event.Handler == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found", "event": eventName, "provider": provider.Name})
			return
		}

		err := provider.Handler(structs.ProviderContext{
			ID:       idParam,
			Secret:   secretParam,
			Event:    event,
			Provider: provider,
			Payload:  payload,
			Context:  c,
		})

		if err != nil {
			log.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal server error occurred when handling the event.", "provider": providerParam, "event": c.GetHeader(provider.Header)})
			return
		}

		c.JSON(200, gin.H{"message": "Event successfully handled"})
		return
	})
}

func loadProviders() {
	// Other things to send to/from:
	// Slack, RocketChat, HipChat, Telegram, riot.im, IRC, XMPP
	// Xenforo 2, IPSuite, MyBB, phpBB, Flarem, Discourse
	addProvider(structs.Provider{
		Name:    "gitlab",
		Header:  "X-Gitlab-Event",
		Handler: gitlab.Handler,
		Events: map[string]structs.Event{
			"Push Hook": {
				Handler: gitlab.PushHandler,
			},
		},
	})
	addProvider(structs.Provider{
		Name:      "sonarr",
		EventName: "eventType",
		Handler:   sonarr.Handler,
		Events: map[string]structs.Event{
			"Test": {
				Handler: sonarr.TestHandler,
			},
			"Grab": {
				Handler: sonarr.GrabHandler,
			},
			"Download": {
				Handler: sonarr.DownloadHandler,
			},
		},
	})
	addProvider(structs.Provider{
		Name: "github",
	})
	/*  addProvider(structs.Provider{
		Name: "CircleCI",
	})
	addProvider(structs.Provider{
		Name: "Trello",
	})
	addProvider(structs.Provider{
		Name: "Asana",
	})
	addProvider(structs.Provider{
		Name: "Datadog",
	})
	addProvider(structs.Provider{
		Name: "BitBucket",
	})
	addProvider(structs.Provider{
		Name: "Gitlab",
	})
	addProvider(structs.Provider{
		Name: "Gitlab",
	})

	addProvider(structs.Provider{
		Name: "Radarr",
	})
	addProvider(structs.Provider{
		Name: "Ombi",
	})
	addProvider(structs.Provider{
		Name: "Plex",
	}) */
}

func addProvider(info structs.Provider) {
	ProviderList += info.Name + ", "
	EventCount += len(info.Events)
	Providers[info.Name] = info
}
