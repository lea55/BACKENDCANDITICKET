package otp

import "time"

type Entity struct {
	ID        string
	Code      string
	Type      Type
	CreatedAt time.Time
	ExpireAt  time.Time
	Enabled   bool
	UserID    string
}

type Type string
