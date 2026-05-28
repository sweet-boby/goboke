package dto

import "goboke/internal/model"

// RegisterRequest represents registration data
type RegisterRequest struct {
	Username        *string `json:"username" binding:"required,min=3,max=30"`
	Avatar          *string `json:"avatar"`
	Phone           *string `json:"phone" binding:"required"`
	Password        *string `json:"password" binding:"required,min=8"`
	ConfirmPassword *string `json:"confirm_password" binding:"required"`
	IP              string  `json:"-"`
}

// LoginRequest represents login credentials
type LoginRequest struct {
	Phone    *string `json:"phone" binding:"required"`
	Password *string `json:"password" binding:"required,min=8"`
}

type UpdateUserprofileRequest struct {
	ID       *int    `json:"-"`
	Username *string `json:"username"`
	Phone    *string `json:"phone"`
	Avatar   *string `json:"avatar"`
}

type UpdateUserRoleRequest struct {
	ID   int            `json:"-"`
	Role model.UserRole `json:"role" binding:"required"`
}

type RefreshTokensRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
