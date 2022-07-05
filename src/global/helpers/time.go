package helpers

import (
	"errors"
	"os"
	"time"
)

type Time struct {
	Format       string
	SimpleFormat string
}

func NewTime() *Time {
	return &Time{
		Format:       os.Getenv("TIME_FORMAT"),
		SimpleFormat: os.Getenv("SIMPLE_TIME_FORMAT"),
	}
}

func (t Time) ValidateDate(dateStr string) (time.Time, error) {
	date, err := time.Parse(t.Format, dateStr)
	if err != nil {
		return time.Time{}, errors.New("El formato de fecha no es válido")
	}
	return date.UTC(), nil
}
