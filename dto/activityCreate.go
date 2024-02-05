package dto

import "github.com/caioap/desafio_bonde/model"

type ActivityCreate struct {
	Type        model.ActivityType `json:"type"`
	Date        string             `json:"date"`
	PersonId    int                `json:"personId"`
	ChallengeId int                `json:"challengeId"`
}
