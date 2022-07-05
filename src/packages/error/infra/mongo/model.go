package mdberrlog

import (
	"time"
)

type model struct {
	ID          string    `bson:"_id,omitempty"`
	DomainID    string    `json:"domain_id"`
	Description string    `json:"description"`
	Module      string    `json:"module"`
	Date        time.Time `json:"date"`
	UserID      string    `json:"user_id"`
}

func (md *model) fromEntity(data errlog.Entity) {
	md.DomainID = data.ID
	md.Description = data.Description
	md.Module = data.Module
	md.Date = data.Date
	md.UserID = data.UserID
}
