package structs

type SonarrBase struct {
	EventType string `json:"eventType"`
}

// Event structs
type SonarrGrab struct {
	SonarrBase
	Series   SonarrSeries    `json:"series"`
	Episodes []SonarrEpisode `json:"episodes"`
	Release  SonarrRelease   `json:"release"`
}

type SonarrDownload struct {
	SonarrBase
	Series      SonarrSeries      `json:"series"`
	Episodes    []SonarrEpisode   `json:"episodes"`
	EpisodeFile SonarrEpisodeFile `json:"episodeFile"`
	IsUpgrade   bool              `json:"isUpgrade"`
}

type SonarrRename struct {
	SonarrBase
	Series SonarrSeries `json:"series"`
}

type SonarrTest struct {
	SonarrBase
	Series   SonarrSeries    `json:"series"`
	Episodes []SonarrEpisode `json:"episodes"`
}

// Modular Structs
type SonarrSeries struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Path   string `json:"path"`
	TvDBId int    `json:"tvdbId"`
}

type SonarrEpisode struct {
	Id             int    `json:"id"`
	EpisodeNumber  int    `json:"episodeNumber"`
	SeasonNumber   int    `json:"seasonNumber"`
	Title          string `json:"title"`
	AirDate        string `json:"airDate"`
	AirDateUTC     string `json:"airDateUtc"`
	Quality        string `json:"quality"`
	QualityVersion int    `json:"qualityVersion"`
}

type SonarrEpisodeFile struct {
	Id             int    `json:"id"`
	RelativePath   string `json:"relativePath"`
	Path           string `json:"path"`
	Quality        string `json:"quality"`
	QualityVersion int    `json:"qualityVersion"`
}

type SonarrRelease struct {
	Quality        string `json:"quality"`
	QualityVersion int    `json:"qualityVersion"`
	Size           int    `json:"size"`
}
