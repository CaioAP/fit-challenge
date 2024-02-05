package model

import "time"

type ActivityType string

const (
	Basketball   ActivityType = "basketball"
	Calisthenics ActivityType = "calisthenics"
	Crossfit     ActivityType = "crossfit"
	Cycling      ActivityType = "cycling"
	Football     ActivityType = "football"
	Gym          ActivityType = "gym"
	Hike         ActivityType = "hike"
	Handball     ActivityType = "handball"
	MartialArts  ActivityType = "martialArts"
	Run          ActivityType = "run"
	Swimming     ActivityType = "swimming"
	Volleyball   ActivityType = "volleyball"
	Other        ActivityType = "other"
)

type Activity struct {
	ID        int
	Type      ActivityType
	Date      time.Time
	Person    Person
	Challenge Challenge
}
