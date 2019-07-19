package main

import (
	"git.deploys.io/disweb/gohook/providers/gitlab"
	"git.deploys.io/disweb/gohook/structs"
	"git.deploys.io/disweb/gohook/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var Providers []structs.Provider
var EventCount int
var ProviderList string

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: false,
	})
	log.WithFields(log.Fields{
		"version": "v0.0.1",
	}).Info("Loading Application...")

	setupRoutes(router)
	loadProviders()

	log.WithFields(log.Fields{
		"version":   "v0.0.1",
		"providers": len(Providers),
		"events":    EventCount,
	}).Infof("Loaded providers: %s", ProviderList)

	_ = router.Run(":8443")
}

func setupRoutes(router *gin.Engine) {
	router.POST("/api/hook/:ID/:Secret/:Provider", func(c *gin.Context) {
		idParam := c.Param("ID")
		secretParam := c.Param("Secret")
		providerParam := c.Param("Provider")
		providerIndex := utils.IndexOfProvider(providerParam, Providers)

		if providerIndex == -1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Provider not found", "provider": providerParam})
			return
		}
		provider := Providers[providerIndex]

		eventIndex := utils.IndexOfEvent(c.GetHeader(provider.Header), provider.Events)

		if eventIndex == -1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found", "event": c.GetHeader(provider.Header), "provider": provider.Name})
			return
		}
		event := provider.Events[eventIndex]

		err := provider.Handler(structs.ProviderContext{
			ID:       idParam,
			Secret:   secretParam,
			Event:    event,
			Provider: provider,
			Context:  c,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal server error occurred when handling the event.", "provider": providerParam, "event": c.GetHeader(provider.Header)})
			return
		}

		c.JSON(200, gin.H{"message": "Event successfully handled"})
		return
	})
}

func loadProviders() {
	// Other things to send to:
	// Slack, RocketChat, HipChat, Telegram, riot.im, IRC, XMPP
	// Xenforo 2, IPSuite, MyBB, phpBB, Flarem, Discourse
	addProvider(structs.Provider{
		Name:    "Gitlab",
		Header:  "X-Gitlab-Event",
		Handler: gitlab.Handler,
		Events: []structs.Event{
			{
				Name:    "Push Hook",
				Handler: gitlab.PushHandler,
			},
		},
	})
	addProvider(structs.Provider{
		Name: "Github",
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
		Name: "Sonarr",
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
	Providers = append(Providers, info)
}
