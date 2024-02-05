package model

import "time"

type Ranking struct {
	ID        int
	Completed int
	Remaining int
	Position  int
	UpdatedAt time.Time
	Person    Person
	Challenge Challenge
}
