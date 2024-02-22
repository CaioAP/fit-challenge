package model

import (
	"time"
)

type Challenge struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Goal         int       `json:"goal"`
	MaxPerDay    int       `json:"maxPerDay"`
	StartDate    time.Time `json:"startDate"`
	FinishDate   time.Time `json:"finishDate"`
	CreatedAt    time.Time `json:"createdAt"`
	Participants int       `json:"participants"`
	Ranking      int       `json:"ranking"`
	Activities   int       `json:"activities"`
	Owner        Person    `json:"owner"`
	Person       []Person  `json:"people"`
}

func (c *Challenge) CalculateRemainingDays(date time.Time) int {
	if date.After(c.FinishDate) {
		return -1
	}
	difference := c.FinishDate.Sub(date)
	differenceInDays := int(difference.Hours() / 24)
	return int(differenceInDays)
}

func (c *Challenge) ValidateMaxPerDay(quantity int) bool {
	return quantity < c.MaxPerDay
}

func (c *Challenge) ValidateDate(date time.Time) bool {
	return (date.Equal(c.StartDate) || date.Equal(c.FinishDate)) ||
		(date.After(c.StartDate) && date.Before(c.FinishDate))
}

func (c *Challenge) AddPerson(person Person) {
	c.Person = append(c.Person, person)
}
