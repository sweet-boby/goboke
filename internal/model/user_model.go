package model

import "time"

type UserRole string

// User roles
const (
	RoleUser      UserRole = "user"
	RoleAdmin     UserRole = "admin"
	RoleModerator UserRole = "moderator"
)

func (s UserRole) IsUserRole() bool {
	return s == RoleAdmin || s == RoleUser || s == RoleModerator
}

// Configuration
var (
	MaxFailedAttempts = 5
	LockoutDuration   = 30 * time.Minute
)

// User represents a user in the system
type User struct {
	ID             int        `json:"id"`
	Username       string     `json:"username" binding:"required,min=3,max=30"`
	Phone          string     `json:"phone"`
	Avatar         string     `json:"avatar"`
	IP             string     `json:"ip"`
	Password       string     `json:"-"` // Never return in JSON
	PasswordHash   string     `json:"-"`
	Role           UserRole   `json:"role"`
	LastLogin      *time.Time `json:"last_login"`
	FailedAttempts int        `json:"-"`
	LockedUntil    *time.Time `json:"-"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
