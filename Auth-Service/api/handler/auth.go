package handler

import (
	"fmt"
	"log/slog"
	"math/rand"
	"time"

	"auth/api/token"

	pb "auth/genprotos/auth"

	"github.com/gin-gonic/gin"
)

// Register handles the creation of a new user
// @Summary Register a new user
// @Description Register a new user with username, email, password, and full name
// @Tags Auth
// @Accept json
// @Produce json
// @Param Register body pb.RegisterRequest true "Register user"
// @Success 200 {object} pb.RegisterResponse "Registration successful"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/register [post]
func (h *Handler) Register(ctx *gin.Context) {
	req := &pb.RegisterRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	req.Role = "user"
	tokens := token.GenereteJWTToken(req)
	req.Token = tokens.RefreshToken
	res, err := h.Auth.Register(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error", "details": err.Error()})
		slog.Info(err.Error())
		return
	}
	ctx.JSON(200, res)
}

// Login handles user login
// @Summary Login a user
// @Description Login a user with username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param Login body pb.LoginRequest true "Login user"
// @Success 200 {object} pb.LoginResponse "Login successful"
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/login [post]
func (h *Handler) Login(ctx *gin.Context) {
	req := &pb.LoginRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	res, err := h.Auth.Login(ctx, req)
	if err != nil {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	if !res.Success {
		ctx.JSON(400, res)
		return
	}
	_, err = token.ExtractClaim(res.Token)
	if err != nil {
		ctx.JSON(400, err)
	}
	ctx.JSON(200, res)
}

// Logout handles user logout
// @Summary Logout a user
// @Description Logout a user with token
// @Tags Auth
// @Accept json
// @Produce json
// @Param Logout body pb.LogoutRequest true "Logout user"
// @Success 200 {object} pb.LogoutResponse "Logout successful"
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/logout [post]
func (h *Handler) Logout(ctx *gin.Context) {
	req := &pb.LogoutRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	res, err := h.Auth.Logout(ctx, req)
	if err != nil {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	ctx.JSON(200, res)
}

// ForgotPassword handles forgot password requests
// @Summary Forgot password
// @Description Request password reset
// @Tags Auth
// @Accept json
// @Produce json
// @Param ForgotPassword body pb.ForgotPasswordRequest true "Forgot password"
// @Success 200 {object} pb.ForgotPasswordResponse "Password reset instructions sent"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/forgot-password [post]
func (h *Handler) ForgotPassword(ctx *gin.Context) {
	req := &pb.ForgotPasswordRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	err = h.Redis.SaveEmailCode(req.Email, code, 10*time.Minute)
	if err != nil {
		slog.Info(err.Error())
	}
	ctx.JSON(200, "Code: "+code)
}

// ResetPassword handles password reset
// @Summary Reset password
// @Description Reset password with token and new password
// @Tags Auth
// @Accept json
// @Produce json
// @Param ResetPassword body pb.ResetPasswordRequest true "Reset password"
// @Success 200 {object} pb.ResetPasswordResponse "Password successfully reset"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/reset-password [post]
func (h *Handler) ResetPassword(ctx *gin.Context) {
	req := &pb.ResetPasswordRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	code, err := h.Redis.Get("email_code:" + req.Email)
	if err != nil {
		slog.Info(err.Error())
	}
	
	if code != req.EmailPassword {
		ctx.JSON(400, "error while resetting password. Please check your email and try again")
		return
	}
	res, err := h.Auth.ResetPassword(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(200, res)
}
