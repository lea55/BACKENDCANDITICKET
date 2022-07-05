package mdbotp

import (
	"time"

	otp "github.com/lea55/BACKENDCANDITICKET/src/packages/otp/domain"
)

type Model struct {
	ID        string    `bson:"_id,omitempty"`
	DomainID  string    `bson:"domain_id,omitempty"`
	Code      string    `bson:"code,omitempty"`
	Type      string    `bson:"type,omitempty"`
	CreatedAt time.Time `bson:"created_at"`
	ExpireAt  time.Time `bson:"expire_at"`
	Enabled   bool      `bson:"enabled,omitempty"`
	UserID    string    `bson:"user_id,omitempty"`
}

func (m *Model) fromEntity(data otp.Entity) {
	m.DomainID = data.ID
	m.Code = data.Code
	m.Type = string(data.Type)
	m.CreatedAt = data.CreatedAt
	m.ExpireAt = data.ExpireAt
	m.Enabled = data.Enabled
	m.UserID = data.UserID
}

func (m *Model) toEntity() otp.Entity {
	return otp.Entity{
		ID:        m.DomainID,
		Code:      m.Code,
		Type:      otp.Type(m.Type),
		CreatedAt: m.CreatedAt,
		ExpireAt:  m.ExpireAt,
		Enabled:   m.Enabled,
		UserID:    m.UserID,
	}
}
