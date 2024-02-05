package dto

type ChallengeCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Goal        int    `json:"goal"`
	MaxPerDay   int    `json:"maxPerDay"`
	StartDate   string `json:"startDate"`
	FinishDate  string `json:"finishDate"`
	PersonId    int    `json:"personId"`
}
