package handler

import (
	"goboke/internal/dto"
	"goboke/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, dto.APIResponse{
			Success: false,
			Error:   "Invalid input data",
		})
		return
	}

	req.IP = c.ClientIP()
	user, err := h.userService.Register(req)

	if err != nil {
		c.JSON(500, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
			Message: "register fail",
		})
		return
	}

	c.JSON(201, dto.APIResponse{
		Success: true,
		Message: "User registered successfully",
		Data:    user,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, dto.APIResponse{
			Success: false,
			Error:   "Invalid credentials format",
		})
		return
	}

	tokens, err := h.userService.Login(req)

	if err != nil {
		c.JSON(400, dto.APIResponse{
			Success: false,
			Message: "Invalid credentials format",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(200, dto.APIResponse{
		Success: true,
		Data:    tokens,
		Message: "Login successful",
	})

}

// POST /auth/logout - User logout
func (h *UserHandler) Logout(c *gin.Context) {
	// TODO: Extract token from Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, dto.APIResponse{
			Success: false,
			Error:   "Authorization header required",
		})
		return
	}

	h.userService.Logout(authHeader)

	c.JSON(200, dto.APIResponse{
		Success: true,
		Message: "Logout successful",
	})
}

// GET /user/profile - Get current user profile
func (h *UserHandler) GetUserProfile(c *gin.Context) {
	// TODO: Get user ID from context (set by authMiddleware)
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(400, dto.APIResponse{
			Success: false,
		})
	}
	// TODO: Find user by ID
	user := h.userService.FindUserByID(userID.(int))
	// TODO: Return user profile (without sensitive data)

	c.JSON(200, dto.APIResponse{
		Success: true,
		Data:    user, // TODO: Return user data
		Message: "Profile retrieved successfully",
	})
}

// PUT /user/profile - Update user profile
func (h *UserHandler) UpdateUserProfile(c *gin.Context) {
	var req dto.UpdateUserprofileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, dto.APIResponse{
			Success: false,
			Error:   "Invalid input data",
		})
		return
	}

	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(400, dto.APIResponse{
			Success: false,
		})
	}
	userIDInt := userID.(int)
	req.ID = &userIDInt

	// TODO: Get user ID from context
	// TODO: Find user by ID
	// TODO: Check if new email is already taken
	// TODO: Update user profile

	h.userService.UpdateUserProfile(req)

	c.JSON(200, dto.APIResponse{
		Success: true,
		Message: "Profile updated successfully",
	})
}

// POST /auth/refresh - Refresh access token
func (h *UserHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokensRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(401, dto.APIResponse{
			Success: false,
			Error:   "Refresh token required",
		})
		return
	}
	tokens, err := h.userService.RefreshToken(req)

	if err != nil {
		c.JSON(500, dto.APIResponse{
			Success: false,
			Message: "regresh token fail",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(200, dto.APIResponse{
		Success: true,
		Message: "Token refreshed successfully",
		Data:    tokens,
	})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.FindAllUser()

	if err != nil {
		c.JSON(404, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
			Message: "not found all user list",
		})
		return
	}

	c.JSON(200, dto.APIResponse{
		Success: true,
		Data:    users,
	})

}

func (h *UserHandler) UpdateUserRole(c *gin.Context) {
	id := c.Param("id")

	userID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	var req dto.UpdateUserRoleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, dto.APIResponse{
			Success: false,
			Message: "json parse fail",
		})
		return
	}

	req.ID = userID
	err = h.userService.UpdateUserRole(req)

	if err != nil {
		c.JSON(500, dto.APIResponse{
			Success: false,
			Error:   err.Error(),
			Message: "update fail",
		})
		return
	}

	c.JSON(200, dto.APIResponse{
		Success: true,
		Message: "set user role success",
	})

}
