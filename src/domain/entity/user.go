package entity

import "time"

type User struct {
	ID              string      `json:"id"`
	Names           string      `json:"names"`
	Surnames        string      `json:"surnames"`
	Email           string      `json:"email"`
	Password        string      `json:"password"`
	LastSessionIp   string      `json:"last_session_ip"`
	LastSessionDate time.Time   `json:"last_session_date"`
	Contact         UserContact `json:"contact"`
	Enabled         bool        `json:"enabled"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	Rol             Rol         `json:"rol"`
	DeviceID        string      `json:"device_id"`
	//new thigns
	CoverImage   string `json:"cover_image"`
	ProfilePhoto string `json:"profile_photo"`
}

type BasicUser struct {
	Names        string `json:"names"`
	Surnames     string `json:"surnames"`
	ID           string `json:"id"`
	Email        string `json:"email"`
	RolCode      string `json:"rol_code"`
	ProfileImage string `json:"image"`
}

type UserContact struct {
	Country     string `json:"country"`
	City        string `json:"city"`
	CountryCode string `json:"country_code"`
	Currency    string `json:"currency"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
}
