package service

import (
	"errors"
	"goboke/internal/dto"
	"goboke/internal/model"
	"goboke/internal/repository"
	"goboke/internal/util/jwt"
	"goboke/internal/util/password"
	"strings"
	"time"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Register(req dto.RegisterRequest) (*model.User, error) {

	if req.Password == nil {
		return nil, errors.New("Passwords is nil")
	}

	if req.ConfirmPassword == nil {
		return nil, errors.New("ConfirmPassword is nil")
	}

	if req.Username == nil {
		return nil, errors.New("Username is nil")
	}

	if req.Phone == nil {
		return nil, errors.New("Phone is nil")
	}

	// TODO: Validate password confirmation
	if *req.Password != *req.ConfirmPassword {

		return nil, errors.New("Passwords do not match")
	}

	// TODO: Validate password strength
	if !password.IsStrongPassword(*req.Password) {

		return nil, errors.New("Password does not meet strength requirements")
	}

	// TODO: Check if username already exists
	// TODO: Check if email already exists

	user, err := s.userRepo.Create(model.User{
		Username: *req.Username,
		Phone:    *req.Phone,
		Avatar:   *req.Avatar,
		IP:       req.IP,
		Password: *req.Password,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(req dto.LoginRequest) (*jwt.TokenResponse, error) {

	// TODO: Find user by username
	user := s.userRepo.FindUserByPhone(*req.Phone)
	if user == nil {

		return nil, errors.New("Invalid credentials")
	}

	// TODO: Check if account is locked
	if isAccountLocked(user) {
		return nil, errors.New("Account is temporarily locked")
	}

	// TODO: Verify password
	if !password.VerifyPassword(*req.Password, user.PasswordHash) {
		recordFailedAttempt(user)
		return nil, errors.New("Invalid credentials")
	}

	// TODO: Reset failed attempts on successful login
	resetFailedAttempts(user)

	// TODO: Update last login time
	now := time.Now()
	user.LastLogin = &now

	// TODO: Generate tokens
	tokens, err := jwt.GenerateTokens(user.ID, user.Username, user.Role)
	if err != nil {
		return nil, errors.New("Failed to generate tokens")
	}

	return tokens, nil
}

func (s *UserService) Logout(authHeader string) {

	// TODO: Extract token from "Bearer <token>" format
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	// TODO: Add token to blacklist
	jwt.BlacklistedTokens[tokenString] = true
	// TODO: Remove refresh token from store
	delete(jwt.RefreshTokens, tokenString)
}

func (s *UserService) FindUserByID(id int) *model.User {

	return s.userRepo.FindUserByID(id)

}

func (s *UserService) UpdateUserProfile(req dto.UpdateUserprofileRequest) error {
	if req.ID == nil {
		return errors.New("user id is nil")
	}

	user := s.userRepo.FindUserByID(*req.ID)

	if user == nil {
		return errors.New("user not found")
	}

	if req.Avatar != nil {
		user.Avatar = *req.Avatar
	}

	if req.Phone != nil {
		user.Phone = *req.Phone
	}

	if req.Username != nil {
		user.Username = *req.Username
	}

	return nil
}

func (s *UserService) RefreshToken(req dto.RefreshTokensRequest) (*jwt.TokenResponse, error) {
	// TODO: Validate refresh token
	userID, ok := jwt.RefreshTokens[req.RefreshToken]
	if !ok {
		return nil, errors.New("unable refresh token ")
	}
	// TODO: Get user ID from refresh token store

	// TODO: Find user by ID
	user := s.userRepo.FindUserByID(userID)
	if user == nil {
		return nil, errors.New("user not found ")
	}
	// TODO: Generate new access token
	tokens, err := jwt.GenerateTokens(user.ID, user.Username, user.Role)
	// TODO: Optionally rotate refresh token

	if err != nil {
		return nil, err
	}

	return tokens, nil
}

// TODO: Implement account lockout check
func isAccountLocked(user *model.User) bool {
	// TODO: Check if account is locked based on LockedUntil field
	if user.LockedUntil == nil {
		return false
	}

	return time.Now().Before(*user.LockedUntil)
}

// TODO: Implement failed attempt tracking
func recordFailedAttempt(user *model.User) {
	// TODO: Increment failed attempts counter
	// TODO: Lock account if max attempts reached
	user.FailedAttempts += 1
	if user.FailedAttempts > model.MaxFailedAttempts {
		t := time.Now().Add(model.LockoutDuration)
		user.LockedUntil = &t
	}
}

func resetFailedAttempts(user *model.User) {
	// TODO: Reset failed attempts counter and unlock account
	user.FailedAttempts = 0
	user.LockedUntil = nil
}
