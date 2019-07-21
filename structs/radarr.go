type RadarrGrab struct {
	RadarrBase
	Series   RadarrMovie    `json:"movie"`
	Episodes []RadarrRemoteMovie `json:"remoteMovie"`
	Release  RadarrRelease   `json:"release"`
}

type RadarrDownload struct {
	RadarrBase
	Series      RadarrMovie      `json:"movie"`
	Episodes    []RadarrRemoteMovie   `json:"remoteMovie"`
	EpisodeFile RadarrMovieFile `json:"movieFile"`
	IsUpgrade   bool              `json:"isUpgrade"`
}

type RadarrRename struct {
	RadarrBase
	Series RadarrMovie `json:"movie"`
}

type RadarrTest struct {
	RadarrBase
	Series   RadarrMovie    `json:"movie"`
	Episodes []RadarrRemoteMovie`json:"remoteMovie"`
	Episodes []RadarrRelease `json:"release"`
}