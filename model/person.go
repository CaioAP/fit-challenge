package model

import "time"

type Person struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
}
