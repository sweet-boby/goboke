package model

import "time"

// User roles
const (
	RoleUser      = "user"
	RoleAdmin     = "admin"
	RoleModerator = "moderator"
)

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
	Role           string     `json:"role"`
	LastLogin      *time.Time `json:"last_login"`
	FailedAttempts int        `json:"-"`
	LockedUntil    *time.Time `json:"-"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
