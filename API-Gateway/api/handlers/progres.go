package handlers

import (
	"context"
	"net/http"

	ln "gateway/genprotos/progres"

	"github.com/gin-gonic/gin"
)

// SetGoal godoc
// @Summary Set a goal
// @Description Set a new goal for the user
// @Tags progress
// @Accept json
// @Produce json
// @Param goal body ln.Goal true "Set Goal"
// @Success 200 {object} ln.GoalResponse
// @Router /progress/set-goal [post]
func (h *ProgressHandler) SetGoal(c *gin.Context) {
	var req ln.Goal
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Progress.SetGoal(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateGoal godoc
// @Summary Update a goal
// @Description Update an existing goal for the user
// @Tags progress
// @Accept json
// @Produce json
// @Param goal_id path string true "Goal ID"
// @Param goal body ln.UpdateGoalRequest true "Update Goal"
// @Success 200 {object} ln.GoalResponse
// @Router /progress/update-goal/{goal_id} [put]
func (h *ProgressHandler) UpdateGoal(c *gin.Context) {
	goalID := c.Param("goal_id")
	var req ln.UpdateGoalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.GoalId = goalID

	res, err := h.Progress.UpdateGoal(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetLanguageProgress godoc
// @Summary Get language progress
// @Description Get the user's progress for a specific language
// @Tags progress
// @Produce json
// @Param language_code path string true "Language Code"
// @Param user_id query string true "User ID"
// @Success 200 {object} ln.ProgressResponse
// @Router /progress/{language_code} [get]
func (h *ProgressHandler) GetLanguageProgress(c *gin.Context) {
	languageCode := c.Param("language_code")
	userID := c.Query("user_id")
	req := &ln.GetLanguageProgressRequest{
		UserId:       userID,
		LanguageCode: languageCode,
	}

	res, err := h.Progress.GetLanguageProgress(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetDailyProgress godoc
// @Summary Get daily progress
// @Description Get the user's daily progress
// @Tags progress
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {object} ln.DailyProgressResponse
// @Router /progress/daily [get]
func (h *ProgressHandler) GetDailyProgress(c *gin.Context) {
	userID := c.Query("user_id")
	req := &ln.GetUserRequest{UserId: userID}

	res, err := h.Progress.GetDailyProgress(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetWeeklyProgress godoc
// @Summary Get weekly progress
// @Description Get the user's weekly progress
// @Tags progress
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {object} ln.WeeklyProgressResponse
// @Router /progress/weekly [get]
func (h *ProgressHandler) GetWeeklyProgress(c *gin.Context) {
	userID := c.Query("user_id")
	req := &ln.GetUserRequest{UserId: userID}

	res, err := h.Progress.GetWeeklyProgress(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetMonthlyProgress godoc
// @Summary Get monthly progress
// @Description Get the user's monthly progress
// @Tags progress
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {object} ln.MonthlyProgressResponse
// @Router /progress/monthly [get]
func (h *ProgressHandler) GetMonthlyProgress(c *gin.Context) {
	userID := c.Query("user_id")
	req := &ln.GetUserRequest{UserId: userID}

	res, err := h.Progress.GetMonthlyProgress(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAchievements godoc
// @Summary Get achievements
// @Description Get the user's achievements
// @Tags progress
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {object} ln.AchievementsResponse
// @Router /progress/achievements [get]
func (h *ProgressHandler) GetAchievements(c *gin.Context) {
	userID := c.Query("user_id")
	req := &ln.GetUserRequest{UserId: userID}

	res, err := h.Progress.GetAchievements(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetLeaderboard godoc
// @Summary Get leaderboard
// @Description Get the leaderboard for a specific language
// @Tags progress
// @Produce json
// @Param language_code path string true "Language Code"
// @Param user_id query string true "User ID"
// @Success 200 {object} ln.LeaderboardResponse
// @Router /progress/leaderboard/{language_code} [get]
func (h *ProgressHandler) GetLeaderboard(c *gin.Context) {
	languageCode := c.Param("language_code")
	userID := c.Query("user_id")
	req := &ln.GetLeaderboardRequest{
		UserId:       userID,
		LanguageCode: languageCode,
	}

	res, err := h.Progress.GetLeaderboard(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetSkills godoc
// @Summary Get skills
// @Description Get the user's skills for a specific language
// @Tags progress
// @Produce json
// @Param language_code path string true "Language Code"
// @Param user_id query string true "User ID"
// @Success 200 {object} ln.SkillsResponse
// @Router /progress/skills/{language_code} [get]
func (h *ProgressHandler) GetSkills(c *gin.Context) {
	languageCode := c.Param("language_code")
	userID := c.Query("user_id")
	req := &ln.GetSkillsRequest{
		UserId:       userID,
		LanguageCode: languageCode,
	}

	res, err := h.Progress.GetSkills(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetGoals godoc
// @Summary Get goals
// @Description Get the user's goals
// @Tags progress
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {object} ln.GoalsResponse
// @Router /progress/goals [get]
func (h *ProgressHandler) GetGoals(c *gin.Context) {
	userID := c.Query("user_id")
	req := &ln.GetUserRequest{UserId: userID}

	res, err := h.Progress.GetGoals(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetStatistics godoc
// @Summary Get statistics
// @Description Get the user's statistics for a specific language
// @Tags progress
// @Produce json
// @Param language_code path string true "Language Code"
// @Param user_id query string true "User ID"
// @Success 200 {object} ln.StatisticsResponse
// @Router /progress/statistics/{language_code} [get]
func (h *ProgressHandler) GetStatistics(c *gin.Context) {
	languageCode := c.Param("language_code")
	userID := c.Query("user_id")
	req := &ln.GetStatisticsRequest{
		UserId:       userID,
		LanguageCode: languageCode,
	}

	res, err := h.Progress.GetStatistics(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteGoal godoc
// @Summary Delete a goal
// @Description Delete an existing goal
// @Tags progress
// @Produce json
// @Param goal_id path string true "Goal ID"
// @Success 200 {object} ln.GoalResponse
// @Router /progress/delete-goal/{goal_id} [delete]
func (h *ProgressHandler) DeleteGoal(c *gin.Context) {
	goalID := c.Param("goal_id")
	req := &ln.DeleteGoalRequest{GoalId: goalID}

	res, err := h.Progress.DeleteGoal(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
