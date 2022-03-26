package entity

import "time"

type Event struct {
	ID               string              `json:"id"`
	Title            string              `json:"title"`
	Description      string              `json:"description"`
	Lat              string              `json:"lat"`
	Lng              string              `json:"lng"`
	StartDate        time.Time           `json:"start_date"`
	EndDate          time.Time           `json:"end_date"`
	TimeZone         string              `json:"time_zone"`
	Image            string              `json:"image"`
	ActiveSeats      bool                `json:"active_seats"`
	Active           bool                `json:"active"`
	Period           bool                `json:"period"`
	LabelEvent       []LabelEventItem    `json:"label_event"`
	OrganizerProfile []Organizer         `json:"organizer_profile"`
	CategorySupport  CategorySupportItem `json:"category_support"`
}
