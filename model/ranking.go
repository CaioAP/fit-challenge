package model

import "time"

type Ranking struct {
	ID        int       `json:"id"`
	Completed int       `json:"completed"`
	Remaining int       `json:"remaining"`
	Position  int       `json:"position"`
	UpdatedAt time.Time `json:"updatedAt"`
	Person    Person    `json:"person"`
	Challenge Challenge `json:"challenge"`
}
