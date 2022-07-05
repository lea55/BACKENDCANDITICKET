package errlog

import "time"

type Entity struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Module      string    `json:"module"`
	Date        time.Time `json:"date"`
	UserID      string    `json:"user_id"`
}
