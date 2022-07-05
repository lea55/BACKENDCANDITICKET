package appotp

import (
	otp "github.com/lea55/BACKENDCANDITICKET/src/packages/otp/domain"
)

type RqGenerate struct {
	Type   otp.Type
	UserID string
}
