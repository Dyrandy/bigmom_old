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

type QueryCurrentProjectResponse struct {
	CurrentProject struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		Version   string    `json:"version"`
		UpdatedAt time.Time `json:"updatedAt"`
	} `json:"currentProject"`
}

type CreateTamperRuleResponse struct {
	CreateTamperRule struct {
		Typename string `json:"__typename"`
	} `json:"createTamperRule"`
}

type TamperRuleCollectionReponse struct {
	Data struct {
		Rules []struct {
			Name        string `json:"name"`
			Strategy    string `json:"strategy"`
			IsRegex     bool   `json:"isRegex"`
			MatchTerm   string `json:"matchTerm"`
			ReplaceTerm string `json:"replaceTerm"`
			IsEnabled   bool   `json:"isEnabled"`
		} `json:"rules"`
	} `json:"tamperRuleCollection"`
}
