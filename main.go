package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"lab.venix.dev/widgetbot/gohook/providers/datadog"
	"lab.venix.dev/widgetbot/gohook/providers/github"
	"lab.venix.dev/widgetbot/gohook/providers/gitlab"
	"lab.venix.dev/widgetbot/gohook/providers/radarr"
	"lab.venix.dev/widgetbot/gohook/providers/sonarr"
	"lab.venix.dev/widgetbot/gohook/structs"
	"lab.venix.dev/widgetbot/gohook/utils"
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
		"port":    ":8443",
	}).Info("Loaded Application!")
	_ = router.Run(":8443")
}

func setupRoutes(router *gin.Engine) {
	router.POST("/api/hook/:ID/:Secret/:Provider", func(c *gin.Context) {
		var event structs.Event
		var eventName string
		var provider structs.Provider
		var BaseDetection structs.BaseDetection

		payload, _ := utils.Parse(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(payload))
		_ = json.Unmarshal([]byte(payload), &BaseDetection)
		idParam := c.Param("ID")
		secretParam := c.Param("Secret")
		providerParam := c.Param("Provider")
		provider = Providers[providerParam]

		if provider.Name == "" {
			c.JSON(http.StatusNotImplemented, gin.H{"error": "Provider not found", "provider": providerParam})
			return
		}

		if provider.Header != "" {
			event = provider.Events[c.GetHeader(provider.Header)]
			eventName = c.GetHeader(provider.Header)
		} else {
			event = provider.Events[utils.EventDetection(BaseDetection)]
			eventName = utils.EventDetection(BaseDetection)
		}
		log.Info(event)
		log.Info(eventName)

		if event.Handler == nil {
			c.JSON(http.StatusNotImplemented, gin.H{"error": "Event not found", "event": eventName, "provider": provider.Name})
			return
		}

		err := provider.Handler(structs.ProviderContext{
			ID:        idParam,
			Secret:    secretParam,
			Event:     event,
			EventName: eventName,
			Provider:  provider,
			Options:   "",
			Payload:   payload,
			Context:   c,
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
		Name:      "datadog",
		Logo:      "https://imgix.datadoghq.com/img/dd_logo_n_70x75.png",
		EventName: "event_type",
		Handler:   datadog.Handler,
		Events: map[string]structs.Event{
			"metric_alert_monitor": {
				Handler: datadog.MetricHandler,
			},
			"synthetics_alert": {
				Handler: datadog.SyntheticsHandler,
			},
		},
	})
	addProvider(structs.Provider{
		Name:    "gitlab",
		Logo:    "https://upload.wikimedia.org/wikipedia/commons/thumb/1/18/GitLab_Logo.svg/1108px-GitLab_Logo.svg.png",
		Header:  "X-Gitlab-Event",
		Handler: gitlab.Handler,
		Events: map[string]structs.Event{
			"Push Hook": {
				Handler: gitlab.PushHandler,
			},
			"Tag Push Hook": {
				Handler: gitlab.TagHandler,
			},
			"Issue Hook": {
				Handler: gitlab.IssueHandler,
			},
			"Note Hook": {
				Handler: gitlab.NoteHandler,
			},
			"System Hook": {
				Handler: gitlab.SystemHandler,
			},
			"Pipeline Hook": {
				Handler: gitlab.PipelineHandler,
			},
			"Job Hook": {
				Handler: gitlab.JobHandler,
			},
			"Build Hook": {
				Handler: gitlab.JobHandler,
			},
			"Merge Request Hook": {},
			"Wiki Page Hook":     {},
		},
	})
	addProvider(structs.Provider{
		Name:      "sonarr",
		Logo:      "https://avatars1.githubusercontent.com/u/1082903?s=400&v=4",
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
		Name:      "radarr",
		Logo:      "https://opencollective-production.s3-us-west-1.amazonaws.com/a8160b50-2b5d-11e8-b4a5-2f63677431ab.png",
		EventName: "eventType",
		Handler:   radarr.Handler,
		Events: map[string]structs.Event{
			"Test": {
				Handler: radarr.TestHandler,
			},
			"Grab": {
				Handler: radarr.GrabHandler,
			},
			"Download": {
				Handler: radarr.DownloadHandler,
			},
		},
	})
	addProvider(structs.Provider{
		Name:    "github",
		Logo:    "https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png",
		Header:  "X-GitHub-Event",
		Handler: github.Handler,
		Events: map[string]structs.Event{
			"ping": {
				Handler: github.PingHandler,
			},
			"push": {
				Handler: github.PushHandler,
			},
			"repository": {
				Handler: github.RepositoryHandler,
			},
			"create":                         {},
			"delete":                         {},
			"deployment":                     {},
			"deployment_status":              {},
			"fork":                           {},
			"gollum":                         {},
			"installation":                   {},
			"installation_repositories":      {},
			"issue_comment":                  {},
			"issues":                         {},
			"label":                          {},
			"marketplace_purchase":           {},
			"member":                         {},
			"membership":                     {},
			"meta":                           {},
			"milestone":                      {},
			"organization":                   {},
			"org_block":                      {},
			"page_build":                     {},
			"project_card":                   {},
			"project_column":                 {},
			"project":                        {},
			"public":                         {},
			"pull_request":                   {},
			"pull_request_review":            {},
			"pull_request_review_comment":    {},
			"registry_package":               {},
			"release":                        {},
			"repository_dispatch":            {},
			"repository_import":              {},
			"repository_vulnerability_alert": {},
			"security_advisory":              {},
			"star":                           {},
			"status":                         {},
			"team":                           {},
			"team_add":                       {},
			"watch":                          {},
		},
	})
	addProvider(structs.Provider{
		Name:      "Plex",
		EventName: "event",
		Events: map[string]structs.Event{
			"library.on.deck": {}, // TODO: Find documentation for these events
			"library.new":     {},
			"media.pause":     {},
			"media.play":      {},
			"media.rate":      {},
			"media.resume":    {},
			"media.scrobble":  {},
			"media.stop":      {},
		},
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
		Name: "Ombi",
	})
	*/
}

func addProvider(info structs.Provider) {
	ProviderList += info.Name + ", "
	EventCount += len(info.Events)
	Providers[info.Name] = info
}
