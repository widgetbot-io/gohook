package structs

type RadarrBase struct {
	EventType string `json:"eventType"`
}

// Event structs
type RadarrGrab struct {
	RadarrBase
	Series   RadarrMovie         `json:"movie"`
	Episodes []RadarrRemoteMovie `json:"remoteMovie"`
	Release  RadarrRelease       `json:"release"`
}

type RadarrDownload struct {
	RadarrBase
	Series      RadarrMovie         `json:"movie"`
	Episodes    []RadarrRemoteMovie `json:"remoteMovie"`
	EpisodeFile RadarrMovieFile     `json:"movieFile"`
	IsUpgrade   bool                `json:"isUpgrade"`
}

type RadarrRename struct {
	RadarrBase
	Series RadarrMovie `json:"movie"`
}

type RadarrTest struct {
	RadarrBase
	Series   RadarrMovie         `json:"movie"`
	Episodes []RadarrRemoteMovie `json:"remoteMovie"`
	Episodes []RadarrRelease     `json:"release"`
}

type RadarrMovie struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"releaseDate"`
}

type RadarrRemoteMovie struct {
	TmdbId int    `json:"tmdbId"`
	ImdbId string `json:"imdbId"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
}

type RadarrRelease struct {
	Quality        string `json:"quality"`
	QualityVersion int    `json:"qualityVersion"`
	ReleaseGroup   string `json:"releaseGroup"`
	ReleaseTitle   string `json:"releaseTitle"`
	Indexer        string `json:"indexer"`
	Size           int    `json:"size"`
}

type RadarrMovieFile struct {
	Id             int    `json:"id"`
	RelativePath   string `json:"relativePath"`
	Path           string `json:"path"`
	Quality        string `json:"quality"`
	QualityVersion int    `json:"qualityVersion"`
	ReleaseGroup   string `json:"releaseGroup"`
}
