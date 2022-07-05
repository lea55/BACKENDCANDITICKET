package useruc

import (
	"github.com/lea55/BACKENDCANDITICKET/src/global/platform"
	rol "github.com/lea55/BACKENDCANDITICKET/src/packages/rol/domain"
	user "github.com/lea55/BACKENDCANDITICKET/src/packages/user/domain"
)

type RqUpdateEmailRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	NewEmail string `json:"new_email" binding:"required,email"`
}

type RqUpdatePasswordRequest struct {
	UserEmail string `json:"user_email" binding:"required"`
}
type RqUpdatePassword struct {
	NewPassword string `json:"new_password" binding:"required"`
	OtpCode     string `json:"otp_code" binding:"required"`
	UserEmail   string `json:"user_email" binding:"required"`
}
type RqUpdateEmail struct {
	UserID   string `json:"user_id" binding:"required"`
	NewEmail string `json:"new_email" binding:"required,email"`
	OtpCode  string `json:"otp_code" binding:"required"`
}

type RqRegisterUser struct {
	Name     string `json:"name"  binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UseCaseModel struct {
	repo    user.Repository
	rolRepo rol.Repository
	otp     *appotp.UseCase
	email   platform.Email
}

func NewUseCaseModel(
	repo user.Repository,
	rolRepo rol.Repository,
	otp *appotp.UseCase,
	email platform.Email,
) *UseCaseModel {
	return &UseCaseModel{
		email:   email,
		repo:    repo,
		rolRepo: rolRepo,
		otp:     otp,
	}
}
