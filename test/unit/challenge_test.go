package unit

import (
	"testing"
	"time"

	"github.com/caioap/desafio_bonde/model"
)

var initDate, _ = time.Parse(time.RFC3339, "2024-01-01T00:00:00Z")
var finishDate, _ = time.Parse(time.RFC3339, "2024-01-10T23:59:59Z")
var challenge = model.Challenge{
	ID:          1,
	Name:        "Test",
	Description: "Testing test",
	Goal:        10,
	MaxPerDay:   1,
	StartDate:   initDate,
	FinishDate:  finishDate,
}

type calculateRemainingDays struct {
	date     time.Time
	expected int
}

var date20231231, _ = time.Parse("2006-01-02", "2023-12-31")
var date20240101, _ = time.Parse("2006-01-02", "2024-01-01")
var date20240105, _ = time.Parse("2006-01-02", "2024-01-05")
var date20240110, _ = time.Parse("2006-01-02", "2024-01-10")
var date20240111, _ = time.Parse("2006-01-02", "2024-01-11")

var calculateRemainingDaysTests = []calculateRemainingDays{
	{date20240101, 9},
	{date20240105, 5},
	{date20240110, 0},
	{date20240111, -1},
}

func TestCalculateRemainingDays(t *testing.T) {
	for _, test := range calculateRemainingDaysTests {
		result := challenge.CalculateRemainingDays(test.date)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}

type maxPerDay struct {
	quantity int
	expected bool
}

var maxPerDayTests = []maxPerDay{
	{0, true},
	{1, false},
}

func TestMaxPerDay(t *testing.T) {
	for _, test := range maxPerDayTests {
		result := challenge.ValidateMaxPerDay(test.quantity)
		if result != test.expected {
			t.Errorf("Expected %t, got %t", test.expected, result)
		}
	}
}

type validateDate struct {
	date     time.Time
	expected bool
}

var validateDateTests = []validateDate{
	{date20231231, false},
	{date20240101, true},
	{date20240105, true},
	{date20240110, true},
	{date20240111, false},
}

func TestValidateDate(t *testing.T) {
	for _, test := range validateDateTests {
		result := challenge.ValidateDate(test.date)
		if result != test.expected {
			t.Errorf("Expected %t, got %t", test.expected, result)
		}
	}
}
