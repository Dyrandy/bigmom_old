package replays

type QueryProjectSessionsResponse struct {
	ReplaySessionCollections struct {
		Nodes []struct {
			Name     string `json:"name"`
			Sessions []struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"sessions"`
		} `json:"nodes"`
	} `json:"replaySessionCollections"`
}

type QueryProjectReplayDataResponse struct {
	ReplaySession struct {
		Name        string `json:"name"`
		ActiveEntry struct {
			Request struct {
				Host          string `json:"host"`
				Path          string `json:"path"`
				Method        string `json:"method"`
				Query         string `json:"query"`
				Raw           string `json:"raw"`
				IsTLS         bool   `json:"isTls"`
				FileExtension any    `json:"fileExtension"`
				Source        string `json:"source"`
				Port          int    `json:"port"`
			} `json:"request"`
		} `json:"activeEntry"`
	} `json:"replaySession"`
}
