package handler

import (
	pb "auth/genprotos/user"

	"github.com/gin-gonic/gin"
)

// GetProfile handles fetching the user profile
// @Summary Get user profile
// @Description Retrieve user profile information
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param token query string true "Bearer Token"
// @Success 200 {object} pb.GetProfileResponse "Profile retrieved"
// @Failure 401 {string} string "Unauthorized"
// @Router /user/profile [get]
func (h *Handler) GetProfile(ctx *gin.Context) {
	token := ctx.Query("token")
	req := &pb.GetProfileRequest{Token: token}
	res, err := h.User.GetProfile(ctx, req)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// UpdateProfile handles updating the user profile
// @Summary Update user profile
// @Description Update user profile information
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param UpdateProfile body pb.UpdateProfileRequest true "Update profile"
// @Success 200 {object} pb.UpdateProfileResponse "Profile updated"
// @Failure 401 {string} string "Unauthorized"
// @Router /user/profile [put]
func (h *Handler) UpdateProfile(ctx *gin.Context) {
	req := &pb.UpdateProfileRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		panic(err)
	}
	res, err := h.User.UpdateProfile(ctx, req)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// ChangePassword handles changing the user password
// @Summary Change user password
// @Description Change user password
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param ChangePassword body pb.ChangePasswordRequest true "Change password"
// @Success 200 {object} pb.ChangePasswordResponse "Password changed"
// @Failure 401 {string} string "Unauthorized"
// @Router /user/password [put]
func (h *Handler) ChangePassword(ctx *gin.Context) {
	req := &pb.ChangePasswordRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		panic(err)
	}
	res, err := h.User.ChangePassword(ctx, req)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}
