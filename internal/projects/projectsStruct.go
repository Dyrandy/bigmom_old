package projects

import "time"

type QueryProjectsResponse struct {
	Projects []struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		Version   string    `json:"version"`
		Size      int       `json:"size"`
		UpdatedAt time.Time `json:"updatedAt"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"projects"`
}
